//go:generate packer-sdc struct-markdown
//go:generate packer-sdc mapstructure-to-hcl2 -type Config,NetworkInterface,StorageVolume

package iso

import (
	packerCommon "github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/communicator"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/martezr/packer-plugin-mvm/builder/mvm/common"
)

type Config struct {
	packerCommon.PackerConfig   `mapstructure:",squash"`
	Comm                        communicator.Config `mapstructure:",squash"`
	common.ConnectConfiguration `mapstructure:",squash"`
	ConvertToTemplate           bool               `mapstructure:"convert_to_template"`
	SkipAgentInstall            bool               `mapstructure:"skip_agent_install"`
	ClusterName                 string             `mapstructure:"cluster_name"`
	VirtualMachineName          string             `mapstructure:"vm_name" required:"true"`
	VirtualImageID              int64              `mapstructure:"virtual_image_id"`
	TemplateName                string             `mapstructure:"template_name"`
	ServicePlanID               int64              `mapstructure:"plan_id" required:"true"`
	CloudID                     int64              `mapstructure:"cloud_id" required:"true"`
	GroupID                     int64              `mapstructure:"group_id"`
	NetworkInterfaces           []NetworkInterface `mapstructure:"network_interface" required:"true"`
	StorageVolumes              []StorageVolume    `mapstructure:"storage_volume" required:"true"`
}

type NetworkInterface struct {
	NetworkId              int64 `mapstructure:"network_id" required:"true"`
	NetworkInterfaceTypeId int64 `mapstructure:"network_interface_type_id" required:"true"`
}

type StorageVolume struct {
	Name          string `mapstructure:"name"`
	RootVolume    bool   `mapstructure:"root_volume"`
	Size          int64  `mapstructure:"size"`
	StorageTypeID int64  `mapstructure:"storage_type_id"`
	DatastoreID   int64  `mapstructure:"datastore_id"`
}

func (b *Builder) Prepare(raws ...interface{}) (generatedVars []string, warnings []string, err error) {
	err = config.Decode(&b.config, &config.DecodeOpts{
		PluginType:  "mvm",
		Interpolate: true,
	}, raws...)
	if err != nil {
		return nil, nil, err
	}

	// Return the placeholder for the generated data that will become available to provisioners and post-processors.
	// If the builder doesn't generate any data, just return an empty slice of string: []string{}
	buildGeneratedData := []string{"GeneratedMockData"}
	return buildGeneratedData, nil, nil
}
