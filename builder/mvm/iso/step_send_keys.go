package iso

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/gomorpheus/morpheus-go-sdk"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

// This is a definition of a builder step and should implement multistep.Step
type StepSendKeys struct{}

// Run should execute the purpose of this step
func (s *StepSendKeys) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	var (
		instance = state.Get("instance").(*morpheus.Instance)
		ui       = state.Get("ui").(packersdk.Ui)
	)

	ui.Say("Sending console keys...")
	c := state.Get("client").(*morpheus.Client)
	commands := []string{
		"root<enter><wait>",
		"setup-alpine<enter>",
		"us<enter>",
		"us<enter>",
		"alpinetemp<enter>",
		"eth0<enter>",
		"dhcp<enter>",
		"n<enter>",
		"Password123#<enter>", // Enter Password
		"Password123#<enter>", // Confirm Password
		"<enter>",             // Timezone - Lowercase doesn't register properly
		"none<enter>",         // Proxy
		"r<enter>",            // APK Mirror
		"no<enter>",           // Create user
		"openssh<enter>",      // SSH Server
		"yes<enter>",          // Allow root login
		"none<enter>",         // SSH Key
		"sda<enter>",          // Install DISK
		"sys<enter>",          // Install DISK
		"y<enter>",            // Wipe Disk
		"reboot",              // Reboot System
	}
	for _, command := range commands {
		runCommand(*c, instance, command)
	}

	// Determines that should continue to the next step
	return multistep.ActionContinue
}

// Cleanup can be used to clean up any artifact created by the step.
// A step's clean up always run at the end of a build, regardless of whether provisioning succeeds or fails.
func (s *StepSendKeys) Cleanup(_ multistep.StateBag) {
	// Nothing to clean
}

func runCommand(c morpheus.Client, instance *morpheus.Instance, command string) {
	payload := make(map[string]interface{})
	payload["sendKeys"] = true
	payload["script"] = command
	resp, err := c.CreateExecutionRequest(&morpheus.Request{
		Body: payload,
		QueryParams: map[string]string{
			"instanceId": strconv.Itoa(int(instance.ID)),
		},
	})
	if err != nil {
		log.Printf("API FAILURE: %s - %s", resp, err)
	}
	log.Printf("API RESPONSE: %s", resp)
	resultGet := resp.Result.(*morpheus.ExecutionRequestResult)
	time.Sleep(5 * time.Second)

	c.GetExecutionRequest(resultGet.ExecutionRequest.UniqueID, &morpheus.Request{})
	log.Println(resp.JsonData)
}
