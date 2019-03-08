provider "hydra" {
  endpoint = "http://localhost:4445"
}

resource "hydra_client" "guest" {
  description    = "A guest credentials"
  client_id      = "guest"
  client_secret  = "swordfish"
  scope          = "hydra.warden hydra.keys.* hydra.introspect"
  grant_types    = ["client_credentials", "refresh_token"]
  response_types = ["token"]
}
