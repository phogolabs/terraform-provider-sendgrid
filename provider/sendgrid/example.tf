provider "sendgrid" {
  api_key = "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"
}

resource "sendgrid_template" "my_template" {
  name        = "my template"
  subject     = "Welcome to my Website!"
  htm_content = ""
}
