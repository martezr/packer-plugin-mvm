// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package clone

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string                `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType         *string                `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion         *string                `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug               *bool                  `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce               *bool                  `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError             *string                `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars            map[string]string      `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars       []string               `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	Type                      *string                `mapstructure:"communicator" cty:"communicator" hcl:"communicator"`
	PauseBeforeConnect        *string                `mapstructure:"pause_before_connecting" cty:"pause_before_connecting" hcl:"pause_before_connecting"`
	SSHHost                   *string                `mapstructure:"ssh_host" cty:"ssh_host" hcl:"ssh_host"`
	SSHPort                   *int                   `mapstructure:"ssh_port" cty:"ssh_port" hcl:"ssh_port"`
	SSHUsername               *string                `mapstructure:"ssh_username" cty:"ssh_username" hcl:"ssh_username"`
	SSHPassword               *string                `mapstructure:"ssh_password" cty:"ssh_password" hcl:"ssh_password"`
	SSHKeyPairName            *string                `mapstructure:"ssh_keypair_name" undocumented:"true" cty:"ssh_keypair_name" hcl:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   *string                `mapstructure:"temporary_key_pair_name" undocumented:"true" cty:"temporary_key_pair_name" hcl:"temporary_key_pair_name"`
	SSHTemporaryKeyPairType   *string                `mapstructure:"temporary_key_pair_type" cty:"temporary_key_pair_type" hcl:"temporary_key_pair_type"`
	SSHTemporaryKeyPairBits   *int                   `mapstructure:"temporary_key_pair_bits" cty:"temporary_key_pair_bits" hcl:"temporary_key_pair_bits"`
	SSHCiphers                []string               `mapstructure:"ssh_ciphers" cty:"ssh_ciphers" hcl:"ssh_ciphers"`
	SSHClearAuthorizedKeys    *bool                  `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys" hcl:"ssh_clear_authorized_keys"`
	SSHKEXAlgos               []string               `mapstructure:"ssh_key_exchange_algorithms" cty:"ssh_key_exchange_algorithms" hcl:"ssh_key_exchange_algorithms"`
	SSHPrivateKeyFile         *string                `mapstructure:"ssh_private_key_file" undocumented:"true" cty:"ssh_private_key_file" hcl:"ssh_private_key_file"`
	SSHCertificateFile        *string                `mapstructure:"ssh_certificate_file" cty:"ssh_certificate_file" hcl:"ssh_certificate_file"`
	SSHPty                    *bool                  `mapstructure:"ssh_pty" cty:"ssh_pty" hcl:"ssh_pty"`
	SSHTimeout                *string                `mapstructure:"ssh_timeout" cty:"ssh_timeout" hcl:"ssh_timeout"`
	SSHWaitTimeout            *string                `mapstructure:"ssh_wait_timeout" undocumented:"true" cty:"ssh_wait_timeout" hcl:"ssh_wait_timeout"`
	SSHAgentAuth              *bool                  `mapstructure:"ssh_agent_auth" undocumented:"true" cty:"ssh_agent_auth" hcl:"ssh_agent_auth"`
	SSHDisableAgentForwarding *bool                  `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding" hcl:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      *int                   `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts" hcl:"ssh_handshake_attempts"`
	SSHBastionHost            *string                `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host" hcl:"ssh_bastion_host"`
	SSHBastionPort            *int                   `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port" hcl:"ssh_bastion_port"`
	SSHBastionAgentAuth       *bool                  `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth" hcl:"ssh_bastion_agent_auth"`
	SSHBastionUsername        *string                `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username" hcl:"ssh_bastion_username"`
	SSHBastionPassword        *string                `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password" hcl:"ssh_bastion_password"`
	SSHBastionInteractive     *bool                  `mapstructure:"ssh_bastion_interactive" cty:"ssh_bastion_interactive" hcl:"ssh_bastion_interactive"`
	SSHBastionPrivateKeyFile  *string                `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file" hcl:"ssh_bastion_private_key_file"`
	SSHBastionCertificateFile *string                `mapstructure:"ssh_bastion_certificate_file" cty:"ssh_bastion_certificate_file" hcl:"ssh_bastion_certificate_file"`
	SSHFileTransferMethod     *string                `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method" hcl:"ssh_file_transfer_method"`
	SSHProxyHost              *string                `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host" hcl:"ssh_proxy_host"`
	SSHProxyPort              *int                   `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port" hcl:"ssh_proxy_port"`
	SSHProxyUsername          *string                `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username" hcl:"ssh_proxy_username"`
	SSHProxyPassword          *string                `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password" hcl:"ssh_proxy_password"`
	SSHKeepAliveInterval      *string                `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval" hcl:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       *string                `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout" hcl:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string               `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels" hcl:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string               `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels" hcl:"ssh_local_tunnels"`
	SSHPublicKey              []byte                 `mapstructure:"ssh_public_key" undocumented:"true" cty:"ssh_public_key" hcl:"ssh_public_key"`
	SSHPrivateKey             []byte                 `mapstructure:"ssh_private_key" undocumented:"true" cty:"ssh_private_key" hcl:"ssh_private_key"`
	WinRMUser                 *string                `mapstructure:"winrm_username" cty:"winrm_username" hcl:"winrm_username"`
	WinRMPassword             *string                `mapstructure:"winrm_password" cty:"winrm_password" hcl:"winrm_password"`
	WinRMHost                 *string                `mapstructure:"winrm_host" cty:"winrm_host" hcl:"winrm_host"`
	WinRMNoProxy              *bool                  `mapstructure:"winrm_no_proxy" cty:"winrm_no_proxy" hcl:"winrm_no_proxy"`
	WinRMPort                 *int                   `mapstructure:"winrm_port" cty:"winrm_port" hcl:"winrm_port"`
	WinRMTimeout              *string                `mapstructure:"winrm_timeout" cty:"winrm_timeout" hcl:"winrm_timeout"`
	WinRMUseSSL               *bool                  `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl" hcl:"winrm_use_ssl"`
	WinRMInsecure             *bool                  `mapstructure:"winrm_insecure" cty:"winrm_insecure" hcl:"winrm_insecure"`
	WinRMUseNTLM              *bool                  `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm" hcl:"winrm_use_ntlm"`
	Url                       *string                `mapstructure:"url" cty:"url" hcl:"url"`
	Username                  *string                `mapstructure:"username" cty:"username" hcl:"username"`
	Password                  *string                `mapstructure:"password" cty:"password" hcl:"password"`
	AccessToken               *string                `mapstructure:"access_token" cty:"access_token" hcl:"access_token"`
	ConvertToTemplate         *bool                  `mapstructure:"convert_to_template" json:"convert_to_template" cty:"convert_to_template" hcl:"convert_to_template"`
	SkipAgentInstall          *bool                  `mapstructure:"skip_agent_install" cty:"skip_agent_install" hcl:"skip_agent_install"`
	ClusterName               *string                `mapstructure:"cluster_name" json:"cluster_name" required:"true" cty:"cluster_name" hcl:"cluster_name"`
	VirtualMachineName        *string                `mapstructure:"vm_name" required:"true" cty:"vm_name" hcl:"vm_name"`
	VirtualImageID            *int64                 `mapstructure:"virtual_image_id" cty:"virtual_image_id" hcl:"virtual_image_id"`
	TemplateName              *string                `mapstructure:"template_name" cty:"template_name" hcl:"template_name"`
	ServicePlanID             *int64                 `mapstructure:"plan_id" json:"plan_id" required:"true" cty:"plan_id" hcl:"plan_id"`
	Cloud                     *string                `mapstructure:"cloud" json:"cloud" required:"true" cty:"cloud" hcl:"cloud"`
	GroupID                   *int64                 `mapstructure:"group_id" json:"group_id" required:"true" cty:"group_id" hcl:"group_id"`
	NetworkInterfaces         []FlatNetworkInterface `mapstructure:"network_interface" required:"true" cty:"network_interface" hcl:"network_interface"`
	StorageVolumes            []FlatStorageVolume    `mapstructure:"storage_volume" required:"true" cty:"storage_volume" hcl:"storage_volume"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":            &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":          &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":          &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":                 &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                 &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":              &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":        &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":   &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"communicator":                 &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":      &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                     &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                     &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                 &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                 &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":             &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":      &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"temporary_key_pair_type":      &hcldec.AttrSpec{Name: "temporary_key_pair_type", Type: cty.String, Required: false},
		"temporary_key_pair_bits":      &hcldec.AttrSpec{Name: "temporary_key_pair_bits", Type: cty.Number, Required: false},
		"ssh_ciphers":                  &hcldec.AttrSpec{Name: "ssh_ciphers", Type: cty.List(cty.String), Required: false},
		"ssh_clear_authorized_keys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_key_exchange_algorithms":  &hcldec.AttrSpec{Name: "ssh_key_exchange_algorithms", Type: cty.List(cty.String), Required: false},
		"ssh_private_key_file":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_certificate_file":         &hcldec.AttrSpec{Name: "ssh_certificate_file", Type: cty.String, Required: false},
		"ssh_pty":                      &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                  &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_wait_timeout":             &hcldec.AttrSpec{Name: "ssh_wait_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":               &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding": &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":       &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":             &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":             &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":       &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":         &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":         &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_interactive":      &hcldec.AttrSpec{Name: "ssh_bastion_interactive", Type: cty.Bool, Required: false},
		"ssh_bastion_private_key_file": &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_bastion_certificate_file": &hcldec.AttrSpec{Name: "ssh_bastion_certificate_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":     &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":               &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":               &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":           &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":           &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":      &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":       &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":           &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":            &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":               &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":              &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":               &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":               &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                   &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_no_proxy":               &hcldec.AttrSpec{Name: "winrm_no_proxy", Type: cty.Bool, Required: false},
		"winrm_port":                   &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":               &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":               &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"url":                          &hcldec.AttrSpec{Name: "url", Type: cty.String, Required: false},
		"username":                     &hcldec.AttrSpec{Name: "username", Type: cty.String, Required: false},
		"password":                     &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"access_token":                 &hcldec.AttrSpec{Name: "access_token", Type: cty.String, Required: false},
		"convert_to_template":          &hcldec.AttrSpec{Name: "convert_to_template", Type: cty.Bool, Required: false},
		"skip_agent_install":           &hcldec.AttrSpec{Name: "skip_agent_install", Type: cty.Bool, Required: false},
		"cluster_name":                 &hcldec.AttrSpec{Name: "cluster_name", Type: cty.String, Required: false},
		"vm_name":                      &hcldec.AttrSpec{Name: "vm_name", Type: cty.String, Required: false},
		"virtual_image_id":             &hcldec.AttrSpec{Name: "virtual_image_id", Type: cty.Number, Required: false},
		"template_name":                &hcldec.AttrSpec{Name: "template_name", Type: cty.String, Required: false},
		"plan_id":                      &hcldec.AttrSpec{Name: "plan_id", Type: cty.Number, Required: false},
		"cloud":                        &hcldec.AttrSpec{Name: "cloud", Type: cty.String, Required: false},
		"group_id":                     &hcldec.AttrSpec{Name: "group_id", Type: cty.Number, Required: false},
		"network_interface":            &hcldec.BlockListSpec{TypeName: "network_interface", Nested: hcldec.ObjectSpec((*FlatNetworkInterface)(nil).HCL2Spec())},
		"storage_volume":               &hcldec.BlockListSpec{TypeName: "storage_volume", Nested: hcldec.ObjectSpec((*FlatStorageVolume)(nil).HCL2Spec())},
	}
	return s
}

// FlatNetworkInterface is an auto-generated flat version of NetworkInterface.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatNetworkInterface struct {
	NetworkId              *int64 `mapstructure:"network_id" required:"true" cty:"network_id" hcl:"network_id"`
	NetworkInterfaceTypeId *int64 `mapstructure:"network_interface_type_id" required:"true" cty:"network_interface_type_id" hcl:"network_interface_type_id"`
}

// FlatMapstructure returns a new FlatNetworkInterface.
// FlatNetworkInterface is an auto-generated flat version of NetworkInterface.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*NetworkInterface) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatNetworkInterface)
}

// HCL2Spec returns the hcl spec of a NetworkInterface.
// This spec is used by HCL to read the fields of NetworkInterface.
// The decoded values from this spec will then be applied to a FlatNetworkInterface.
func (*FlatNetworkInterface) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"network_id":                &hcldec.AttrSpec{Name: "network_id", Type: cty.Number, Required: false},
		"network_interface_type_id": &hcldec.AttrSpec{Name: "network_interface_type_id", Type: cty.Number, Required: false},
	}
	return s
}

// FlatStorageVolume is an auto-generated flat version of StorageVolume.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatStorageVolume struct {
	Name          *string `mapstructure:"name" cty:"name" hcl:"name"`
	RootVolume    *bool   `mapstructure:"root_volume" cty:"root_volume" hcl:"root_volume"`
	Size          *int64  `mapstructure:"size" cty:"size" hcl:"size"`
	StorageTypeID *int64  `mapstructure:"storage_type_id" cty:"storage_type_id" hcl:"storage_type_id"`
	DatastoreID   *int64  `mapstructure:"datastore_id" cty:"datastore_id" hcl:"datastore_id"`
}

// FlatMapstructure returns a new FlatStorageVolume.
// FlatStorageVolume is an auto-generated flat version of StorageVolume.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*StorageVolume) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatStorageVolume)
}

// HCL2Spec returns the hcl spec of a StorageVolume.
// This spec is used by HCL to read the fields of StorageVolume.
// The decoded values from this spec will then be applied to a FlatStorageVolume.
func (*FlatStorageVolume) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"name":            &hcldec.AttrSpec{Name: "name", Type: cty.String, Required: false},
		"root_volume":     &hcldec.AttrSpec{Name: "root_volume", Type: cty.Bool, Required: false},
		"size":            &hcldec.AttrSpec{Name: "size", Type: cty.Number, Required: false},
		"storage_type_id": &hcldec.AttrSpec{Name: "storage_type_id", Type: cty.Number, Required: false},
		"datastore_id":    &hcldec.AttrSpec{Name: "datastore_id", Type: cty.Number, Required: false},
	}
	return s
}
