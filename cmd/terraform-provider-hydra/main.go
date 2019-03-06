package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/phogolabs/terraform-provider/hydra"
)

func main() {
	provider := &hydra.Provider{}

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return provider.Definition()
		},
	})
}
