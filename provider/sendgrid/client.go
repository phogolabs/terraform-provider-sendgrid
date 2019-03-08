package sendgrid

// Template represents a template
type Template struct {
	ID         string             `json:"id"`
	Name       string             `json:"name"`
	Generation string             `json:"generation"`
	Versions   []*TemplateVersion `json:"versions"`
}

// TemplateVersion is a template content's version
type TemplateVersion struct {
	ID           string `json:"id"`
	TemplateID   string `json:"template_id"`
	Subject      string `json:"subject"`
	HTMLContent  string `json:"html_content"`
	PlainContent string `json:"plain_content"`
	Editor       string `json:"editor"`
}

// Client represents a SendGrid Client
type Client struct {
	Key string
}

// CreateTemplate creates a template
func (c *Client) CreateTemplate(template *Template) error {
	return nil
}

// GetTemplate gets a template
func (c *Client) GetTemplate(id string) (*Template, error) {
	return nil, nil
}

// UpdateTemplate updates a template
func (c *Client) UpdateTemplate() error {
	return nil
}

// DeleteTemplate deletes a template
func (c *Client) DeleteTemplate(id string) error {
	return nil
}

func (*Client) do(method, path string, body interface{}) error {
	request := sendgrid.GetRequest(c.Key, path, "")
	request.Method = method
	return nil
}
