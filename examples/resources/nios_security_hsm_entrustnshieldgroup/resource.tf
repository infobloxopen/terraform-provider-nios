// Create an EntrustnShield HSM Group with Basic Fields
resource "nios_security_hsm_entrustnshieldgroup" "basic" {
  name          = "example-hsm-entrustnshieldgroup"
  key_server_ip = "< RFS IP>"
  entrustnshield_hsm = [
    {
      keyhash   = "< Keyhash value >"
      remote_ip = "< Remote IP >"
    }
  ]
}

//Create an EntrustnShield HSM Group with Additional Fields
resource "nios_security_hsm_entrustnshieldgroup" "additional" {
  name            = "example-hsm-entrustnshieldgroup1"
  protection      = "SOFTCARD"
  card_name       = "example-softcard"
  pass_phrase     = "examplepassphrase@123"
  key_server_ip   = "< RFS IP >"
  key_server_port = 9004
  comment         = "Example EntrustnShield group"
  entrustnshield_hsm = [
    {
      keyhash     = "< Keyhash value >"
      remote_ip   = "< Remote IP >"
      remote_port = 9004
      disable     = false
    }
  ]
}
