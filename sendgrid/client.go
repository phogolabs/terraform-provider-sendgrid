package sendgrid

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
)

// Template represents a template
type Template struct {
	ID         string             `json:"id,omitempty"`
	Name       string             `json:"name,omitempty"`
	Generation string             `json:"generation,omitempty"`
	Versions   []*TemplateVersion `json:"versions,omitempty"`
}

// TemplateVersion is a template content's version
type TemplateVersion struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	TemplateID   string `json:"template_id,omitempty"`
	Subject      string `json:"subject,omitempty"`
	HTMLContent  string `json:"html_content,omitempty"`
	PlainContent string `json:"plain_content,omitempty"`
	Editor       string `json:"editor,omitempty"`
	Active       int    `json:"active,omitempty"`
}

// Client represents a SendGrid Client
type Client struct {
	Key string
}

// CreateTemplate creates a template
func (c *Client) CreateTemplate(template *Template) error {
	response, err := c.do(rest.Post, "", template)

	if err == nil {
		err = c.unmarshal(response.Body, template)
	}

	return err
}

// GetTemplate gets a template
func (c *Client) GetTemplate(id string) (*Template, error) {
	response, err := c.do(rest.Get, id, nil)

	if err == nil {
		template := &Template{}

		if err = c.unmarshal(response.Body, template); err == nil {
			return template, nil
		}
	}

	return nil, err
}

// UpdateTemplate updates a template
func (c *Client) UpdateTemplate(id string, template *Template) error {
	response, err := c.do(rest.Patch, id, template)

	if err == nil {
		err = c.unmarshal(response.Body, template)
	}

	return err
}

// DeleteTemplate deletes a template
func (c *Client) DeleteTemplate(id string) error {
	_, err := c.do(rest.Delete, id, nil)
	return err
}

// CreateTemplateVersion creates a template version
func (c *Client) CreateTemplateVersion(version *TemplateVersion) error {
	response, err := c.do(rest.Post, c.path(version.TemplateID, "versions"), version)

	if err == nil {
		err = c.unmarshal(response.Body, version)
	}

	return err
}

// GetTemplateVersion gets a template version
func (c *Client) GetTemplateVersion(template, id string) (*TemplateVersion, error) {
	response, err := c.do(rest.Get, c.path(template, "versions", id), nil)

	if err == nil {
		version := &TemplateVersion{}

		if err = c.unmarshal(response.Body, version); err == nil {
			return version, nil
		}
	}

	return nil, err
}

// UpdateTemplateVersion updates a template
func (c *Client) UpdateTemplateVersion(id string, version *TemplateVersion) error {
	response, err := c.do(rest.Patch, c.path(version.TemplateID, "versions", id), version)

	if err == nil {
		err = c.unmarshal(response.Body, version)
	}

	return err
}

func (c *Client) path(path ...string) string {
	return filepath.Join(path...)
}

// DeleteTemplateVersion deletes a template version
func (c *Client) DeleteTemplateVersion(template, id string) error {
	_, err := c.do(rest.Delete, c.path(template, "versions", id), nil)
	return err
}

func (c *Client) do(method rest.Method, path string, body interface{}) (*rest.Response, error) {
	path = fmt.Sprintf("/v3/templates/%s", path)
	path = strings.TrimSuffix(path, "/")

	request := sendgrid.GetRequest(c.Key, path, "")
	request.Method = method

	if body != nil {
		content, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		request.Body = content
	}

	response, err := sendgrid.API(request)
	if err == nil {
		err = c.error(response)
	}

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) error(response *rest.Response) error {
	if code := response.StatusCode; code >= 400 && code <= 500 {
		err := &Error{}

		if uerr := json.Unmarshal([]byte(response.Body), err); uerr == nil {
			return err
		}

		message := response.Body
		if message == "" {
			message = http.StatusText(code)
		}

		return fmt.Errorf(message)
	}

	return nil
}

func (c *Client) unmarshal(body string, entity interface{}) error {
	return json.Unmarshal([]byte(body), entity)
}

// Error represents an error
type Error struct {
	Message string `json:"error"`
}

// Error returns the message
func (e *Error) Error() string {
	return e.Message
}
