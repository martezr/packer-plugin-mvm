package clone

import (
	"context"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/communicator"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/multistep/commonsteps"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/martezr/packer-plugin-mvm/builder/mvm/common"
)

type Builder struct {
	config Config
	runner multistep.Runner
}

func (b *Builder) ConfigSpec() hcldec.ObjectSpec { return b.config.FlatMapstructure().HCL2Spec() }

func (b *Builder) Run(ctx context.Context, ui packer.Ui, hook packer.Hook) (packer.Artifact, error) {
	// Setup the state bag and initial state for the steps
	state := new(multistep.BasicStateBag)
	state.Put("debug", b.config.PackerDebug)
	state.Put("hook", hook)
	state.Put("ui", ui)

	steps := []multistep.Step{}

	steps = append(steps,
		&common.StepConnect{
			Config: &b.config.ConnectConfiguration,
		},
		&StepProvisionVM{builder: b},
		&communicator.StepConnect{
			Config:    &b.config.Comm,
			Host:      communicator.CommHost(b.config.Comm.Host(), "instance_ip"),
			SSHConfig: b.config.Comm.SSHConfigFunc(),
		},
	)

	if b.config.Comm.Type != "none" {
		steps = append(steps,
			&communicator.StepConnect{
				Config:    &b.config.Comm,
				Host:      communicator.CommHost(b.config.Comm.Host(), "instance_ip"),
				SSHConfig: b.config.Comm.SSHConfigFunc(),
			},
			&commonsteps.StepProvision{},
		)
	} else {
		steps = append(steps,
			&commonsteps.StepProvision{},
		)
	}

	steps = append(steps,
		&common.StepStopInstance{},
		&StepConvertInstance{builder: b},
		//&common.StepRemoveInstance{},
	)

	// Set the value of the generated data that will become available to provisioners.
	// To share the data with post-processors, use the StateData in the artifact.
	state.Put("generated_data", map[string]interface{}{
		"GeneratedMockData": "mock-build-data",
		//	"InstanceId":        state.Get("instance_id"),
	})

	// Run!
	b.runner = commonsteps.NewRunner(steps, b.config.PackerConfig, ui)
	b.runner.Run(ctx, state)

	// If there was an error, return that
	if err, ok := state.GetOk("error"); ok {
		return nil, err.(error)
	}

	artifact := &common.Artifact{
		// Add the builder generated data to the artifact StateData so that post-processors
		// can access them.
		//InstanceId: state.Get("instance_id").(int64),
		StateData: map[string]interface{}{
			"generated_data": state.Get("generated_data"),
		},
	}
	return artifact, nil
}
