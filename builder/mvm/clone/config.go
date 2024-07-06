//go:generate packer-sdc mapstructure-to-hcl2 -type Config

package clone

import (
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/communicator"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	Comm                communicator.Config `mapstructure:",squash"`
	Url                 string              `mapstructure:"url"`
	Username            string              `mapstructure:"username"`
	Password            string              `mapstructure:"password"`
	AccessToken         string              `mapstructure:"access_token"`
	VirtualMachineName  string              `mapstructure:"vm_name"`
	VirtualImageID      int64               `mapstructure:"virtual_image_id"`
	TemplateName        string              `mapstructure:"template_name"`
	ServicePlanID       int64               `mapstructure:"plan_id"`
	CloudID             int64               `mapstructure:"cloud_id"`
	GroupID             int64               `mapstructure:"group_id"`
	NetworkID           int64               `mapstructure:"network_id"`
}

func (b *Builder) Prepare(raws ...interface{}) (generatedVars []string, warnings []string, err error) {
	err = config.Decode(&b.config, &config.DecodeOpts{
		PluginType:  "mvm",
		Interpolate: true,
	}, raws...)
	if err != nil {
		return nil, nil, err
	}

	// Validate that all required fields are present
	var errs *packersdk.MultiError
	required := map[string]string{
		"username": b.config.Username,
		"password": b.config.Password,
	}
	for k, v := range required {
		if v == "" {
			errs = packersdk.MultiErrorAppend(errs, fmt.Errorf("you must specify a %s", k))
		}
	}

	// Return the placeholder for the generated data that will become available to provisioners and post-processors.
	// If the builder doesn't generate any data, just return an empty slice of string: []string{}
	buildGeneratedData := []string{"GeneratedMockData"}
	return buildGeneratedData, nil, nil
}
