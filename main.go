package main

import (
	"fmt"
	"os"

	"github.com/martezr/packer-plugin-mvm/builder/mvm/clone"
	"github.com/martezr/packer-plugin-mvm/builder/mvm/iso"

	"github.com/hashicorp/packer-plugin-mvm/version"
	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder("iso", new(iso.Builder))
	pps.RegisterBuilder("clone", new(clone.Builder))

	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
