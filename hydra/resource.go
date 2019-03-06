package hydra

import (
	"fmt"
	"net/http"

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
				Required:  true,
			},
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"grant_types": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"response_types": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
		},
	}
}

func (r *ClientResource) create(d *schema.ResourceData, m interface{}) error {
	client := m.(*hydra.CodeGenSDK)

	credentials := swagger.OAuth2Client{
		ClientId:      d.Get("client_id").(string),
		ClientSecret:  d.Get("client_secret").(string),
		GrantTypes:    slice(d.Get("grant_types")),
		ResponseTypes: slice(d.Get("response_types")),
	}

	_, response, err := client.CreateOAuth2Client(credentials)
	if err == nil {
		err = errorf(response)
	}

	if err == nil {
		d.SetId(credentials.ClientId)
	}

	return err
}

func (r *ClientResource) read(d *schema.ResourceData, m interface{}) error {
	client := m.(*hydra.CodeGenSDK)

	credentials, response, err := client.GetOAuth2Client(d.Get("client_id").(string))
	if err == nil {
		err = errorf(response)
	}

	if err == nil {
		d.SetId(credentials.ClientId)
		d.Set("client_id", credentials.ClientId)
		d.Set("client_secret", credentials.ClientSecret)
		d.Set("grant_types", credentials.GrantTypes)
		d.Set("response_types", credentials.ResponseTypes)
	}

	return err
}

func (r *ClientResource) update(d *schema.ResourceData, m interface{}) error {
	client := m.(*hydra.CodeGenSDK)

	credentials := swagger.OAuth2Client{
		ClientId:      d.Get("client_id").(string),
		ClientSecret:  d.Get("client_secret").(string),
		GrantTypes:    slice(d.Get("grant_types")),
		ResponseTypes: slice(d.Get("response_types")),
	}

	_, response, err := client.UpdateOAuth2Client(d.Id(), credentials)
	if err == nil {
		err = errorf(response)
	}

	if err == nil {
		d.SetId(credentials.ClientId)
	}

	return err
}

func (r *ClientResource) delete(d *schema.ResourceData, m interface{}) error {
	client := m.(*hydra.CodeGenSDK)

	response, err := client.DeleteOAuth2Client(d.Get("client_id").(string))
	if err == nil {
		err = errorf(response)
	}

	return err
}

func errorf(response *swagger.APIResponse) error {
	defer response.Body.Close()

	if code := response.StatusCode; code >= 400 && code <= 500 {
		return fmt.Errorf("%v %v", http.StatusText(code), response.Message)
	}

	return nil
}

func slice(slice interface{}) []string {
	result := []string{}

	for _, item := range slice.([]interface{}) {
		result = append(result, fmt.Sprintf("%v", item))
	}

	return result
}
