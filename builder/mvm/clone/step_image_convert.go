package clone

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gomorpheus/morpheus-go-sdk"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

// This is a definition of a builder step and should implement multistep.Step
type StepConvertInstance struct {
	builder *Builder
}

// Run should execute the purpose of this step
func (s *StepConvertInstance) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	var (
		instance = state.Get("server").(*morpheus.Instance)
		ui       = state.Get("ui").(packersdk.Ui)
	)

	ui.Say("Converting compute instance to image")

	//	clonePayload := make(map[string]interface{})
	//	clonePayload["templateName"] = "demo123"

	data, err := s.builder.moclient.Execute(&morpheus.Request{
		Method: "PUT",
		//		Body:   clonePayload,
		Path: fmt.Sprintf("/api/instances/%d/import-snapshot", instance.ID),
	})
	if err != nil {
		log.Println(err)
	}

	log.Println(data.Status)
	time.Sleep(180 * time.Second)

	imageResponse, imageErr := s.builder.moclient.ListVirtualImages(&morpheus.Request{
		QueryParams: map[string]string{
			"name": "pack-%",
		},
	})

	if imageErr != nil {
		log.Println(imageErr)
	}
	result := imageResponse.Result.(*morpheus.ListVirtualImagesResult)
	firstRecord := (*result.VirtualImages)[0]
	log.Println("IMAGE STATUS: ", firstRecord)
	virtualImageId := firstRecord.ID

	// Status List: provisioning, pending, cancelled, removing
	// Poll Virtual Images for Status
	currentStatus := "Saving"
	completedStatuses := []string{"Active"}
	log.Println("Polling order status...")

	for !stringInSlice(completedStatuses, currentStatus) {
		resp, err := s.builder.moclient.GetVirtualImage(virtualImageId, &morpheus.Request{})
		if err != nil {
			log.Println("API ERROR: ", err)
		}
		result := resp.Result.(*morpheus.GetVirtualImageResult)
		currentStatus = result.VirtualImage.Status
		log.Println("Current status:", currentStatus)
		time.Sleep(30 * time.Second)
	}

	resp, err := s.builder.moclient.UpdateVirtualImage(virtualImageId, &morpheus.Request{
		Body: map[string]interface{}{
			"name": s.builder.config.TemplateName,
		},
	})
	if err != nil {
		log.Println("API ERROR: ", err)
	}
	log.Println(resp.Status)
	// Determines that should continue to the next step
	return multistep.ActionContinue
}

// Cleanup can be used to clean up any artifact created by the step.
// A step's clean up always run at the end of a build, regardless of whether provisioning succeeds or fails.
func (s *StepConvertInstance) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
}

func stringInSlice(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
