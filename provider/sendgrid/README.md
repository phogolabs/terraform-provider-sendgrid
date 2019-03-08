# SendGrid Provider

It's a terraform provider that allow creating a template in [SendGrid](https://www.sendgrid.com).

## Usage

You can use the [example](./example.tf) or code snippet below:

```
provider "sendgrid" {
  api_key = "your-api-key"
}

resource "sendgrid_template" "my_template" {
  name  = "my-template-1"
}

resource "sendgrid_template_version" "content" {
  name          = "content"
  template_id   = "${sendgrid_template.my_template.id}"
  subject       = "Welcome to the House !!!!"
  html_content  = "<h1>Welcome to My House</h1>"
  plain_content = "Welcome"
  active        = true
}
```
