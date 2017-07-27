package main

import (
	"github.com/celkins/terraform-provider-example/example"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: example.Provider})
}
