//go:generate packer-sdc mapstructure-to-hcl2 -type Config

package morpheus

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gomorpheus/morpheus-go-sdk"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

type Config struct {
	ctx interpolate.Context

	Url         string `mapstructure:"url"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	AccessToken string `mapstructure:"access_token"`

	// The ID of the automation task to execute
	TaskID int `mapstructure:"task_id"`

	// The parameters to pass to the automation task
	TaskParams map[interface{}]interface{} `mapstructure:"task_parameters"`

	// The ID of the automation workflow to execute
	WorkflowID int `mapstructure:"workflow_id"`

	// The phase of the provisioning workflow to execute
	WorkflowPhase string `mapstructure:"workflow_phase"`

	// The parameters to pass to the automation workflow
	WorkflowParams map[interface{}]interface{} `mapstructure:"workflow_parameters"`
}

type Provisioner struct {
	config Config
}

func (p *Provisioner) ConfigSpec() hcldec.ObjectSpec {
	return p.config.FlatMapstructure().HCL2Spec()
}

func (p *Provisioner) Prepare(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		PluginType:         "packer.provisioner.morpheus",
		Interpolate:        true,
		InterpolateContext: &p.config.ctx,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{},
		},
	}, raws...)
	if err != nil {
		return err
	}
	return nil
}

func (p *Provisioner) Provision(_ context.Context, ui packer.Ui, _ packer.Communicator, generatedData map[string]interface{}) error {
	ui.Say("Provisioning with Morpheus provisioner...")
	client := morpheus.NewClient(p.config.Url)
	if p.config.AccessToken != "" {
		client.SetAccessToken(p.config.AccessToken, "", 86400, "write")
	} else {
		client.SetUsernameAndPassword(p.config.Username, p.config.Password)
	}
	resp, err := client.Login()
	if err != nil {
		fmt.Println("LOGIN ERROR: ", err)
	}
	fmt.Println("LOGIN RESPONSE:", resp)
	instanceId := generatedData["ID"].(int64)
	var jobExecutionId int

	// Execute automation task
	if p.config.TaskID != 0 {
		ui.Sayf("Executing Morpheus: running task %d on instance %d", p.config.TaskID, instanceId)
		payloadConfig := make(map[string]interface{})
		jobPayload := make(map[string]interface{})
		jobPayload["name"] = fmt.Sprintf("Packer Provisioner Execution - %d", instanceId)
		jobPayload["targetType"] = "instance"
		jobPayload["instances"] = []int64{instanceId}
		payloadConfig["job"] = jobPayload

		resp, err = client.Execute(&morpheus.Request{
			Method: "POST",
			Body:   payloadConfig,
			Path:   fmt.Sprintf("/api/tasks/%d/execute", p.config.TaskID),
			Result: &TaskExecutionResult{},
		})

		if err != nil {
			if resp != nil && resp.StatusCode == 404 {
				log.Printf("API 404: %s - %v", resp, err)
				return nil
			} else {
				log.Printf("API FAILURE: %s - %v", resp, err)
			}
		}
		log.Printf("API RESPONSE: %s", resp)

		// store resource data
		output := resp.Result.(*TaskExecutionResult)
		jobExecutionId = output.JobExecution.ID
	}

	// Execute automation workflow
	if p.config.WorkflowID != 0 {
		ui.Sayf("Executing Morpheus: running workflow %d on instance %d", p.config.WorkflowID, instanceId)
		payloadConfig := make(map[string]interface{})
		jobPayload := make(map[string]interface{})
		jobPayload["name"] = fmt.Sprintf("Packer Provisioner Execution - %d", instanceId)
		jobPayload["targetType"] = "instance"
		jobPayload["instances"] = []int64{instanceId}
		payloadConfig["job"] = jobPayload

		resp, err = client.Execute(&morpheus.Request{
			Method: "POST",
			Body:   payloadConfig,
			Path:   fmt.Sprintf("/api/task-sets/%d/execute", p.config.TaskID),
			Result: &TaskExecutionResult{},
		})

		if err != nil {
			if resp != nil && resp.StatusCode == 404 {
				log.Printf("API 404: %s - %v", resp, err)
				return nil
			} else {
				log.Printf("API FAILURE: %s - %v", resp, err)
			}
		}
		log.Printf("API RESPONSE: %s", resp)

		// store resource data
		output := resp.Result.(*TaskExecutionResult)
		jobExecutionId = output.JobExecution.ID
	}

	// Fetch Execution Response
	resp, err = client.Execute(&morpheus.Request{
		Method: "GET",
		Path:   fmt.Sprintf("/api/job-executions/%d", jobExecutionId),
		Result: &JobExecutionResult{},
	})

	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Printf("API 404: %s - %v", resp, err)
			return nil
		} else {
			log.Printf("API FAILURE: %s - %v", resp, err)
		}
	}
	log.Printf("API RESPONSE: %s", resp)

	// store resource data
	executionOutput := resp.Result.(*JobExecutionResult)
	log.Println(executionOutput)

	// Status List: provisioning, pending, cancelled, removing
	// Poll Instance for Status
	currentStatus := "queued"
	completedStatuses := []string{"error", "success"}

	for !stringInSlice(completedStatuses, currentStatus) {
		resp, err = client.Execute(&morpheus.Request{
			Method: "GET",
			Path:   fmt.Sprintf("/api/job-executions/%d", jobExecutionId),
			Result: &JobExecutionResult{},
		})

		if err != nil {
			if resp != nil && resp.StatusCode == 404 {
				log.Printf("API 404: %s - %v", resp, err)
				return nil
			} else {
				log.Printf("API FAILURE: %s - %v", resp, err)
			}
		}
		log.Printf("API RESPONSE: %s", resp)

		// store resource data
		executionOutput := resp.Result.(*JobExecutionResult)
		currentStatus = executionOutput.JobExecution.Status
		// sleep 30 seconds between polls
		time.Sleep(30 * time.Second)
	}

	// Fetch final payload data
	resp, err = client.Execute(&morpheus.Request{
		Method: "GET",
		Path:   fmt.Sprintf("/api/job-executions/%d", jobExecutionId),
		Result: &JobExecutionResult{},
	})

	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			log.Printf("API 404: %s - %v", resp, err)
			return nil
		} else {
			log.Printf("API FAILURE: %s - %v", resp, err)
		}
	}
	log.Printf("API RESPONSE: %s", resp)

	executionOutput = resp.Result.(*JobExecutionResult)
	ui.Message(executionOutput.JobExecution.Process.Events[0].Output)
	return nil
}

type TaskExecutionResult struct {
	Success bool `json:"success"`
	Job     struct {
		ID int `json:"id"`
	} `json:"job"`
	JobExecution struct {
		ID        int         `json:"id"`
		ProcessId interface{} `json:"processId"`
	} `json:"jobExecution"`
	Message string            `json:"msg"`
	Errors  map[string]string `json:"errors"`
}

type JobExecutionResult struct {
	Success      bool              `json:"success"`
	JobExecution JobExecution      `json:"jobExecution"`
	Message      string            `json:"msg"`
	Errors       map[string]string `json:"errors"`
}

type ProcessType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
type CreatedBy struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
}

type UpdatedBy struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
}

type Events struct {
	ID            int         `json:"id"`
	ProcessId     int         `json:"processId"`
	AccountId     int         `json:"accountId"`
	UniqueId      string      `json:"uniqueId"`
	ProcessType   ProcessType `json:"processType"`
	Description   string      `json:"description"`
	RefType       string      `json:"refType"`
	RefId         int         `json:"refId"`
	SubType       interface{} `json:"subType"`
	SubId         interface{} `json:"subId"`
	ZoneId        int         `json:"zoneId"`
	IntegrationId interface{} `json:"integrationId"`
	InstanceId    int         `json:"instanceId"`
	ContainerId   int         `json:"containerId"`
	ServerId      int         `json:"serverId"`
	ContainerName string      `json:"containerName"`
	DisplayName   string      `json:"displayName"`
	Status        string      `json:"status"`
	Reason        interface{} `json:"reason"`
	Percent       float64     `json:"percent"`
	StatusEta     int         `json:"statusEta"`
	Message       string      `json:"message"`
	Output        string      `json:"output"`
	Error         string      `json:"error"`
	StartDate     time.Time   `json:"startDate"`
	EndDate       time.Time   `json:"endDate"`
	Duration      int         `json:"duration"`
	DateCreated   time.Time   `json:"dateCreated"`
	LastUpdated   time.Time   `json:"lastUpdated"`
	CreatedBy     CreatedBy   `json:"createdBy"`
	UpdatedBy     UpdatedBy   `json:"updatedBy"`
}

type Process struct {
	ID            int         `json:"id"`
	AccountId     int         `json:"accountId"`
	UniqueId      string      `json:"uniqueId"`
	ProcessType   ProcessType `json:"processType"`
	DisplayName   string      `json:"displayName"`
	Description   string      `json:"description"`
	SubType       interface{} `json:"subType"`
	SubId         interface{} `json:"subId"`
	ZoneId        int         `json:"zoneId"`
	IntegrationId interface{} `json:"integrationId"`
	AppId         interface{} `json:"appId"`
	InstanceId    int         `json:"instanceId"`
	ContainerId   int         `json:"containerId"`
	ServerId      int         `json:"serverId"`
	ContainerName string      `json:"containerName"`
	Status        string      `json:"status"`
	Reason        interface{} `json:"reason"`
	Percent       float64     `json:"percent"`
	StatusEta     int         `json:"statusEta"`
	Message       interface{} `json:"message"`
	Output        interface{} `json:"output"`
	Error         interface{} `json:"error"`
	StartDate     time.Time   `json:"startDate"`
	EndDate       time.Time   `json:"endDate"`
	Duration      int         `json:"duration"`
	DateCreated   time.Time   `json:"dateCreated"`
	LastUpdated   time.Time   `json:"lastUpdated"`
	CreatedBy     CreatedBy   `json:"createdBy"`
	UpdatedBy     UpdatedBy   `json:"updatedBy"`
	Events        []Events    `json:"events"`
}

type Type struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type Job struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description interface{} `json:"description"`
	Type        Type        `json:"type"`
}

type Createdby struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Displayname string `json:"displayName"`
}

type JobExecution struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	Process       Process     `json:"process"`
	Job           Job         `json:"job"`
	Description   interface{} `json:"description"`
	Datecreated   time.Time   `json:"dateCreated"`
	Startdate     time.Time   `json:"startDate"`
	Enddate       time.Time   `json:"endDate"`
	Duration      int         `json:"duration"`
	Resultdata    interface{} `json:"resultData"`
	Status        string      `json:"status"`
	Statusmessage interface{} `json:"statusMessage"`
	Automation    bool        `json:"automation"`
	Createdby     Createdby   `json:"createdBy"`
}

func stringInSlice(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
