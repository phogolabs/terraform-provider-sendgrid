# Ory Hydra

It's a terraform provider that allow creating a client in [Ory Hydra OAuth2](https://www.ory.sh) server.

## Usage

You can use the [example](./example.tf) or code snippet below:

```
provider "hydra" {
  endpoint = "http://localhost:4445"
}

resource "hydra_client" "guest" {
  description    = "A guest credentials"
  client_id      = "guest"
  client_secret  = "guest1234"
  scope          = "hydra.warden hydra.keys.* hydra.introspect"
  grant_types    = ["client_credentials", "refresh_token"]
  response_types = ["token"]
}
```
