package hydra

import "github.com/hashicorp/terraform/helper/schema"

// Provider returns the provider's schema
func Provider() *schema.Provider {
	client := &ClientResource{}

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HYDRA_ADMIN_URL", nil),
				Description: "The administrative URL of ORY Hydra",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"hydra_client": client.Resource(),
		},
	}
}

// ClientResource represent's the flient resource
type ClientResource struct {
}

// Resource returns the resource
func (r *ClientResource) Resource() *schema.Resource {
	return &schema.Resource{
		Create: r.create,
		Read:   r.read,
		Update: r.update,
		Delete: r.delete,
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
			},
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"client_secret": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"grant_types": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
			},
			"response_types": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
			},
		},
	}
}

func (r *ClientResource) create(d *schema.ResourceData, m interface{}) error {
	return nil
}

func (r *ClientResource) read(d *schema.ResourceData, m interface{}) error {
	return nil
}

func (r *ClientResource) update(d *schema.ResourceData, m interface{}) error {
	return nil
}

func (r *ClientResource) delete(d *schema.ResourceData, m interface{}) error {
	return nil
}
