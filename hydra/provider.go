package hydra

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ory/hydra/sdk/go/hydra"
)

// Provider for Hydra OAuth2 Server
type Provider struct{}

// Definition returns the provider's schema
func (p *Provider) Definition() *schema.Provider {
	client := &ClientResource{}

	return &schema.Provider{
		ConfigureFunc: p.configure,
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HYDRA_ADMIN_URL", nil),
				Description: "The administrative URL of ORY Hydra",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"hydra_client": client.Definition(),
		},
	}
}

func (p *Provider) configure(d *schema.ResourceData) (interface{}, error) {
	config := &hydra.Configuration{
		AdminURL: d.Get("endpoint").(string),
	}

	return hydra.NewSDK(config)
}
