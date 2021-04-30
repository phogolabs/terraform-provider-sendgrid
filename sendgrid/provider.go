package sendgrid

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Provider for Hydra OAuth2 Server
type Provider struct{}

// Definition returns the provider's schema
func (p *Provider) Definition() *schema.Provider {
	var (
		template = &TemplateResource{}
		version  = &TemplateVersionResource{}
	)

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
			"sendgrid_template":         template.Definition(),
			"sendgrid_template_version": version.Definition(),
		},
	}
}

func (p *Provider) configure(d *schema.ResourceData) (interface{}, error) {
	return &Client{
		Key: d.Get("api_key").(string),
	}, nil
}
