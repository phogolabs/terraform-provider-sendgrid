package hydra

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ory/hydra/sdk/go/hydra"
	"github.com/ory/hydra/sdk/go/hydra/swagger"
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

// Error represents a hydra error
type Error struct {
	Code        int    `json:"status_code"`
	Message     string `json:"error"`
	Description string `json:"error_description"`
}

func (e *Error) Error() string {
	message := fmt.Sprintf("code: %v message: %s %s", e.Code, e.Message, e.Description)
	message = strings.TrimSpace(message)
	return message
}

func handleErr(response *swagger.APIResponse) error {
	if code := response.StatusCode; code >= 400 && code <= 500 {
		err := &Error{}

		if uerr := json.Unmarshal(response.Payload, err); uerr == nil {
			return err
		}

		message := string(response.Payload)
		if message == "" {
			message = http.StatusText(code)
		}

		return fmt.Errorf(message)
	}

	return nil
}
