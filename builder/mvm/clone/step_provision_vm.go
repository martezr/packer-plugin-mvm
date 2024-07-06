package clone

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gomorpheus/morpheus-go-sdk"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

// This is a definition of a builder step and should implement multistep.Step
type StepProvisionVM struct {
	builder *Builder
}

type NetworkInterface struct {
	Network struct {
		ID string `json:"id"`
	} `json:"network"`
	NetworkInterfaceTypeID int64 `json:"networkInterfaceTypeID"`
}

type StorageVolume struct {
	ID          int64  `json:"id"`
	RootVolume  bool   `json:"rootVolume"`
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	StorageType int64  `json:"storageType"`
	DatastoreId int64  `json:"datastoreId"`
}

// Run should execute the purpose of this step
func (s *StepProvisionVM) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packersdk.Ui)

	// Config
	config := make(map[string]interface{})

	// TODO: Update the instance to MVM
	instanceTypeResponse, err := s.builder.moclient.FindInstanceTypeByName("ubuntu")
	if err != nil {
		log.Printf("API FAILURE: %s - %s", instanceTypeResponse, err)
	}
	log.Printf("API RESPONSE: %s", instanceTypeResponse)
	instanceTypeResult := instanceTypeResponse.Result.(*morpheus.GetInstanceTypeResult)
	instanceType := instanceTypeResult.InstanceType

	// Resource Pool
	resourcePoolResp, err := s.builder.moclient.Execute(&morpheus.Request{
		Method:      "GET",
		Path:        fmt.Sprintf("/api/options/zonePools?layoutId=%d", instanceType.InstanceTypeLayouts[0].ID),
		QueryParams: map[string]string{},
	})
	if err != nil {
		log.Println(err)
	}

	var itemResponsePayload ResourcePoolOptions
	json.Unmarshal(resourcePoolResp.Body, &itemResponsePayload)
	var resourcePoolId int
	for _, v := range itemResponsePayload.Data {
		if v.ProviderType == "mvm" {
			resourcePoolId = v.Id
		}
	}

	config["resourcePoolId"] = resourcePoolId
	config["poolProviderType"] = "mvm"
	// Create User
	config["createUser"] = true

	// Image ID
	//config["imageId"] = s.builder.config.VirtualImageID

	// Skip Agent Install
	config["noAgent"] = true

	instancePayload := map[string]interface{}{
		"name": s.builder.config.VirtualMachineName,
		"type": instanceType.Code,
		"instanceType": map[string]interface{}{
			"code": instanceType.Code,
		},
		"site": map[string]interface{}{
			"id": s.builder.config.GroupID,
		},
		"plan": map[string]interface{}{
			"id": s.builder.config.ServicePlanID,
		},
		"layout": map[string]interface{}{
			"id": 307,
			//	"id":   instanceType.InstanceTypeLayouts[2].ID,
			//	"code": instanceType.InstanceTypeLayouts[2].Code,
			//	"name": instanceType.InstanceTypeLayouts[2].Name,
		},
	}

	payload := map[string]interface{}{
		"zoneId":   s.builder.config.CloudID,
		"instance": instancePayload,
		"config":   config,
	}

	// Network Interfaces
	var Nics []NetworkInterface
	var NetworkDemo NetworkInterface
	NetworkDemo.Network.ID = fmt.Sprintf("network-%d", s.builder.config.NetworkID)
	// TODO: Expose Network Interface Type ID
	NetworkDemo.NetworkInterfaceTypeID = 4
	Nics = append(Nics, NetworkDemo)
	payload["networkInterfaces"] = Nics

	// TODO: Expose storage volume
	// Storage Volumes
	var Volumes []StorageVolume
	var StorageDemo StorageVolume
	StorageDemo.ID = -1
	StorageDemo.Name = "root"
	StorageDemo.RootVolume = true
	StorageDemo.Size = 20
	StorageDemo.StorageType = 1
	StorageDemo.DatastoreId = 10
	Volumes = append(Volumes, StorageDemo)

	payload["volumes"] = Volumes

	payload["layoutSize"] = 1

	// TODO: Remove additional logging
	out, _ := json.Marshal(payload)
	fmt.Println("JSON PAYLOAD")
	ui.Say(fmt.Sprintf("JSON PAYLOAD %s", string(out)))

	req := &morpheus.Request{Body: payload}
	resp, err := s.builder.moclient.CreateInstance(req)
	if err != nil {
		log.Printf("API FAILURE: %s - %s", resp, err)
	}
	log.Printf("API RESPONSE: %s", resp)
	result := resp.Result.(*morpheus.CreateInstanceResult)
	instance := result.Instance

	// TODO: Add polling logic for checking instance status
	time.Sleep(300 * time.Second)

	respGet, err := s.builder.moclient.GetInstance(instance.ID, req)
	if err != nil {
		log.Printf("API FAILURE: %s - %s", resp, err)
	}
	log.Printf("API RESPONSE: %s", resp)
	resultGet := respGet.Result.(*morpheus.GetInstanceResult)
	instanceGet := resultGet.Instance

	state.Put("server", instanceGet)
	state.Put("server_ip", instanceGet.ConnectionInfo[0].Ip)
	state.Put("server_id", instance.ID)

	ui.Say(fmt.Sprintf("Instance IP Address %s", instanceGet.ConnectionInfo[0].Ip))

	// Determines that should continue to the next step
	return multistep.ActionContinue
}

// Cleanup can be used to clean up any artifact created by the step.
// A step's clean up always run at the end of a build, regardless of whether provisioning succeeds or fails.
func (s *StepProvisionVM) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
	/*	ui := state.Get("ui").(packer.Ui)

		if v, ok := state.GetOk("instance"); ok {
			ui.Say("Cleanup: destroying compute instance")
			respGet, err := s.builder.moclient.DeleteInstance(instance.ID, req)
			if err != nil {
				log.Printf("API FAILURE: %s - %s", resp, err)
			}
			log.Printf("API RESPONSE: %s", resp)
			resultGet := respGet.Result.(*morpheus.GetInstanceResult)
			instanceGet := resultGet.Instance

			if err := s.builder.exo.DeleteInstance(ctx, s.builder.config.InstanceZone, instance); err != nil {
				ui.Error(fmt.Sprintf("Unable to delete compute instance: %v", err))
			}
		}
	*/
}

type ResourcePoolOptions struct {
	Success bool `json:"success"`
	Data    []struct {
		Id           int    `json:"id"`
		Name         string `json:"name"`
		IsGroup      bool   `json:"isGroup"`
		Group        string `json:"group"`
		IsDefault    bool   `json:"isDefault"`
		Type         string `json:"type"`
		ProviderType string `json:"providerType"`
		Value        string `json:"value"`
	} `json:"data"`
}
