package sendgrid

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider for Hydra OAuth2 Server
type Provider struct{}

// Definition returns the provider's schema
func (p *Provider) Definition() *schema.Provider {
	client := &ClientResource{}

	return &schema.Provider{
		ConfigureFunc: p.configure,
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SENDGRID_API_KEY", nil),
				Description: "The SendGrid API Key",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"sendgrid_template": client.Definition(),
		},
	}
}

func (p *Provider) configure(d *schema.ResourceData) (interface{}, error) {
	return &Client{
		Key: d.Get("api_key").(string),
	}, nil
}
