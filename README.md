# SendGrid Provider

[![Documentation][godoc-img]][godoc-url]
![License][license-img]
[![Build Status][action-img]][action-url]
[![Go Report Card][report-img]][report-url]

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

[godoc-url]: https://godoc.org/github.com/phogolabs/terraform-provider
[godoc-img]: https://godoc.org/github.com/phogolabs/terraform-provider?status.svg
[license-img]: https://img.shields.io/badge/license-MIT-blue.svg
[report-img]: https://goreportcard.com/badge/github.com/phogolabs/terraform-provider
[report-url]: https://goreportcard.com/report/github.com/phogolabs/terraform-provider
[codecov-url]: https://codecov.io/gh/phogolabs/terraform-provider
[codecov-img]: https://codecov.io/gh/phogolabs/terraform-provider/branch/master/graph/badge.svg
[action-img]: https://github.com/phogolabs/terraform-provider/workflows/main/badge.svg
[action-url]: https://github.com/phogolabs/terraform-provider/actions
