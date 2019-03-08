package hydra

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ory/hydra/sdk/go/hydra"
	"github.com/ory/hydra/sdk/go/hydra/swagger"
)

// ClientResource represent's the flient resource
type ClientResource struct{}

// Definition returns the resource
func (r *ClientResource) Definition() *schema.Resource {
	return &schema.Resource{
		Create: r.create,
		Read:   r.read,
		Update: r.update,
		Delete: r.delete,
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Required: false,
			},
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"client_secret": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Optional:  true,
				Required:  false,
			},
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Required: false,
			},
			"grant_types": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Required: false,
			},
			"response_types": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Required: false,
			},
		},
	}
}

func (r *ClientResource) create(d *schema.ResourceData, m interface{}) error {
	client := m.(*hydra.CodeGenSDK)

	user, response, err := client.CreateOAuth2Client(r.get(d))
	if err == nil {
		err = handleErr(response)
	}

	if err == nil {
		r.set(d, user)
	}

	return err
}

func (r *ClientResource) read(d *schema.ResourceData, m interface{}) error {
	client := m.(*hydra.CodeGenSDK)

	user, response, err := client.GetOAuth2Client(d.Get("client_id").(string))
	if err == nil {
		err = handleErr(response)
	}

	if err == nil {
		r.set(d, user)
	}

	return err
}

func (r *ClientResource) update(d *schema.ResourceData, m interface{}) error {
	client := m.(*hydra.CodeGenSDK)

	user, response, err := client.UpdateOAuth2Client(d.Id(), r.get(d))
	if err == nil {
		err = handleErr(response)
	}

	if err == nil {
		r.set(d, user)
	}

	return err
}

func (r *ClientResource) delete(d *schema.ResourceData, m interface{}) error {
	client := m.(*hydra.CodeGenSDK)

	response, err := client.DeleteOAuth2Client(d.Get("client_id").(string))
	if err == nil {
		err = handleErr(response)
	}

	return err
}

func (r *ClientResource) get(d *schema.ResourceData) swagger.OAuth2Client {
	client := swagger.OAuth2Client{
		ClientId:      d.Get("client_id").(string),
		ClientSecret:  d.Get("client_secret").(string),
		GrantTypes:    slice(d.Get("grant_types")),
		ResponseTypes: slice(d.Get("response_types")),
	}

	return client
}

func (r *ClientResource) set(d *schema.ResourceData, client *swagger.OAuth2Client) {
	d.SetId(client.ClientId)
	d.Set("client_id", client.ClientId)
	d.Set("client_secret", client.ClientSecret)
	d.Set("grant_types", client.GrantTypes)
	d.Set("response_types", client.ResponseTypes)
}

func slice(slice interface{}) []string {
	result := []string{}

	for _, item := range slice.([]interface{}) {
		result = append(result, fmt.Sprintf("%v", item))
	}

	return result
}
