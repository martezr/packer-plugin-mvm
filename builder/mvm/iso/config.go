//go:generate packer-sdc struct-markdown
//go:generate packer-sdc mapstructure-to-hcl2 -type Config,NetworkInterface,StorageVolume

package iso

import (
	packerCommon "github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/communicator"
	"github.com/hashicorp/packer-plugin-sdk/multistep/commonsteps"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
	"github.com/martezr/packer-plugin-mvm/builder/mvm/common"
)

type Config struct {
	packerCommon.PackerConfig   `mapstructure:",squash"`
	commonsteps.HTTPConfig      `mapstructure:",squash"`
	BootConfig                  `mapstructure:",squash"`
	Comm                        communicator.Config `mapstructure:",squash"`
	common.ConnectConfiguration `mapstructure:",squash"`
	HTTPTemplateDirectory       string `mapstructure:"http_template_directory"`
	ConvertToTemplate           bool   `mapstructure:"convert_to_template"`
	SkipAgentInstall            bool   `mapstructure:"skip_agent_install"`
	ClusterName                 string `mapstructure:"cluster_name"`
	VirtualMachineName          string `mapstructure:"vm_name" required:"true"`
	VirtualImageID              int64  `mapstructure:"virtual_image_id"`
	TemplateName                string `mapstructure:"template_name"`
	// The ID of the service plan that will be associated with the instance.
	ServicePlanID int64 `mapstructure:"plan_id" required:"true"`
	// The ID of the cloud that contains the MVM cluster.
	CloudID int64 `mapstructure:"cloud_id" required:"true"`
	// The ID of the Morpheus group to deploy the instance into.
	GroupID           int64               `mapstructure:"group_id"`
	NetworkInterfaces []NetworkInterface  `mapstructure:"network_interface" required:"true"`
	StorageVolumes    []StorageVolume     `mapstructure:"storage_volume" required:"true"`
	Ctx               interpolate.Context `mapstructure-to-hcl2:",skip"`
}

type NetworkInterface struct {
	// The ID of the network to connect the interface to.
	NetworkId int64 `mapstructure:"network_id" required:"true"`
	// The ID of the network interface type used by the network interface.
	NetworkInterfaceTypeId int64 `mapstructure:"network_interface_type_id" required:"true"`
}

type StorageVolume struct {
	// The name of the storage volume.
	Name string `mapstructure:"name"`
	// Whether the storage volume is the root volume.
	RootVolume bool `mapstructure:"root_volume"`
	// The size in GB of the storage volume.
	Size          int64 `mapstructure:"size"`
	StorageTypeID int64 `mapstructure:"storage_type_id"`
	// The ID of the datastore for the storage volume.
	DatastoreID string `mapstructure:"datastore_id"`
}

func (c *Config) Prepare(raws ...interface{}) (generatedVars []string, warnings []string, err error) {
	err = config.Decode(c, &config.DecodeOpts{
		PluginType:         "mvm",
		Interpolate:        true,
		InterpolateContext: &c.Ctx,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{
				"boot_command",
				//"http_content",
			},
		},
	}, raws...)
	if err != nil {
		return nil, nil, err
	}

	var errs *packersdk.MultiError

	errs = packersdk.MultiErrorAppend(errs, c.Comm.Prepare(&c.Ctx)...)
	errs = packersdk.MultiErrorAppend(errs, c.BootConfig.Prepare(&c.Ctx)...)
	errs = packersdk.MultiErrorAppend(errs, c.HTTPConfig.Prepare(&c.Ctx)...)

	if errs != nil && len(errs.Errors) > 0 {
		return nil, warnings, errs
	}
	return nil, warnings, nil
}
