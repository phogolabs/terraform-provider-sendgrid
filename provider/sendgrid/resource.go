package sendgrid

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// TemplateResource represent's the flient resource
type TemplateResource struct{}

// Definition returns the resource
func (r *TemplateResource) Definition() *schema.Resource {
	return &schema.Resource{
		Create: r.create,
		Read:   r.read,
		Update: r.update,
		Delete: r.delete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"generation": &schema.Schema{
				Type:     schema.TypeString,
				Default:  "dynamic",
				Optional: true,
				Required: false,
			},
		},
	}
}

func (r *TemplateResource) create(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	template := r.get(d)
	err := client.CreateTemplate(template)
	if err != nil {
		return err
	}

	r.set(d, template)
	return nil
}

func (r *TemplateResource) read(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	template, err := client.GetTemplate(d.Id())
	if err != nil {
		return err
	}

	r.set(d, template)
	return nil
}

func (r *TemplateResource) update(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	template := r.get(d)
	err := client.UpdateTemplate(d.Id(), template)
	if err != nil {
		return err
	}

	r.set(d, template)
	return nil
}

func (r *TemplateResource) delete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	return client.DeleteTemplate(d.Id())
}

func (r *TemplateResource) get(d *schema.ResourceData) *Template {
	template := &Template{
		ID:         d.Id(),
		Name:       d.Get("name").(string),
		Generation: d.Get("generation").(string),
	}

	return template
}

func (r *TemplateResource) set(d *schema.ResourceData, template *Template) {
	d.SetId(template.ID)
	d.Set("name", template.Name)
	d.Set("generation", template.Generation)
}

// TemplateVersionResource represent's the flient resource
type TemplateVersionResource struct{}

// Definition returns the resource
func (r *TemplateVersionResource) Definition() *schema.Resource {
	return &schema.Resource{
		Create: r.create,
		Read:   r.read,
		Update: r.update,
		Delete: r.delete,
		Schema: map[string]*schema.Schema{
			"template_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"active": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Required: false,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"subject": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"html_content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"plain_content": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"editor": &schema.Schema{
				Type:     schema.TypeString,
				Default:  "code",
				Optional: true,
				Required: false,
			},
		},
	}
}

func (r *TemplateVersionResource) create(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	template := r.get(d)
	err := client.CreateTemplateVersion(template)
	if err != nil {
		return err
	}

	r.set(d, template)
	return nil
}

func (r *TemplateVersionResource) read(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	version, err := client.GetTemplateVersion(d.Get("template_id").(string), d.Id())
	if err != nil {
		return err
	}

	r.set(d, version)
	return nil
}

func (r *TemplateVersionResource) update(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	version := r.get(d)
	version.Editor = ""

	err := client.UpdateTemplateVersion(d.Id(), version)
	if err != nil {
		return err
	}

	r.set(d, version)
	return nil
}

func (r *TemplateVersionResource) delete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)
	return client.DeleteTemplateVersion(d.Get("template_id").(string), d.Id())
}

func (r *TemplateVersionResource) get(d *schema.ResourceData) *TemplateVersion {
	version := &TemplateVersion{
		ID:           d.Id(),
		Name:         d.Get("name").(string),
		TemplateID:   d.Get("template_id").(string),
		Subject:      d.Get("subject").(string),
		HTMLContent:  d.Get("html_content").(string),
		PlainContent: d.Get("plain_content").(string),
		Editor:       d.Get("editor").(string),
	}

	if active := d.Get("active").(bool); active {
		version.Active = 1
	} else {
		version.Active = 0
	}

	return version
}

func (r *TemplateVersionResource) set(d *schema.ResourceData, version *TemplateVersion) {
	d.SetId(version.ID)
	d.Set("template_id", version.TemplateID)
	d.Set("name", version.Name)
	d.Set("subject", version.Subject)
	d.Set("html_content", version.HTMLContent)
	d.Set("plain_content", version.PlainContent)
	d.Set("editor", version.Editor)
	d.Set("active", version.Active == 1)
}
