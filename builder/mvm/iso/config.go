//go:generate packer-sdc struct-markdown
//go:generate packer-sdc mapstructure-to-hcl2 -type Config,NetworkInterface,StorageVolume

package iso

import (
	"time"

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
	// Amount of time to wait for VM's IP, similar to 'ssh_timeout'.
	// Defaults to 30m (30 minutes). See the Golang
	// [ParseDuration](https://golang.org/pkg/time/#ParseDuration) documentation
	// for full details.
	IPWaitTimeout         time.Duration `mapstructure:"ip_wait_timeout"`
	HTTPTemplateDirectory string        `mapstructure:"http_template_directory"`
	// Whether to convert the instance to a virtual image
	ConvertToTemplate bool `mapstructure:"convert_to_template"`
	//SkipAgentInstall            bool   `mapstructure:"skip_agent_install"`
	// The name of the MVM cluster to provision the instance on.
	ClusterName string `mapstructure:"cluster_name"`
	// The name of the instance to provision.
	VirtualMachineName string `mapstructure:"vm_name" required:"true"`
	// The id of the ISO virtual image to use as the instance instance source image.
	VirtualImageID int64 `mapstructure:"virtual_image_id" required:"true"`
	// The name of the virtual image to create.
	TemplateName string `mapstructure:"template_name"`
	// The ID of the service plan that will be associated with the instance.
	ServicePlanID int64 `mapstructure:"plan_id" required:"true"`
	// The name of the cloud that contains the MVM cluster.
	Cloud string `mapstructure:"cloud" required:"true"`
	// The name of the Morpheus group to deploy the instance into.
	Group string `mapstructure:"group" required:"true"`
	// The name of the instance to provision.
	Description         string              `mapstructure:"description"`
	Environment         string              `mapstructure:"environment"`
	Labels              []string            `mapstructure:"labels"`
	NetworkInterfaces   []NetworkInterface  `mapstructure:"network_interface" required:"true"`
	StorageVolumes      []StorageVolume     `mapstructure:"storage_volume" required:"true"`
	HostID              int64               `mapstructure:"host"`
	AttachVirtioDrivers bool                `mapstructure:"attach_virtio_drivers"`
	Ctx                 interpolate.Context `mapstructure-to-hcl2:",skip"`
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
