// Create an Entrust nShield HSM Group with Basic Fields
resource "nios_security_hsm_entrustnshieldgroup" "basic" {
  name          = "example-hsm-entrustnshieldgroup"
  key_server_ip = "192.168.1.100"
  entrustnshield_hsm = [
    {
      keyhash   = "abc123def456abc123def456abc123def456abc123de"
      remote_ip = "192.168.1.101"
    }
  ]
}

//Create an Entrust nShield HSM Group with Additional Fields
resource "nios_security_hsm_entrustnshieldgroup" "additional" {
  name            = "example-hsm-entrustnshieldgroup1"
  protection      = "SOFTCARD"
  card_name       = "example-softcard"
  pass_phrase     = "examplepassphrase@123"
  key_server_ip   = "192.168.1.100"
  key_server_port = 9004
  comment         = "Example EntrustnShield group"
  entrustnshield_hsm = [
    {
      keyhash     = "abc123def456abc123def456abc123def456abc123de"
      remote_ip   = "192.168.1.101"
      remote_port = 9004
      disable     = false
    }
  ]
}
