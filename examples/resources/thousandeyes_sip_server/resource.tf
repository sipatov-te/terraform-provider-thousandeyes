resource "thousandeyes_sip_server" "example_sip_server_test" {
  test_name      = "Example SIP server test set from Terraform provider"
  interval       = 120
  alerts_enabled = false
  target_sip_credentials {
    auth_user     = ""
    protocol      = "tcp"
    port          = 5060
    sip_registrar = "example.org"
  }
  agents = ["3"] # Singapore
}

