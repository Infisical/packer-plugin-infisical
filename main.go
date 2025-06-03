package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/infisical/packer-plugin-infisical/datasource/secrets"
	"github.com/infisical/packer-plugin-infisical/version"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterDatasource("secrets", new(secrets.Datasource))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
