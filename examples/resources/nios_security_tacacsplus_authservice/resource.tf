// Create TACACS+ Authservice with Basic Fields
resource "nios_security_tacacsplus_authservice" "tacacsplus_authservice_with_basic_fields" {
  name = "tacacsplus_authservice1"
  servers = [
    {
      address        = "2.2.3.3"
      auth_type      = "CHAP"
      disable        = false
      port           = 49
      use_accounting = false
      use_mgmt_port  = false
      shared_secret  = "test"
    }
  ]
}

// Create TACACS+ Authservice with Additional Fields
resource "nios_security_tacacsplus_authservice" "tacacsplus_authservice_with_additional_field" {
  name = "tacacsplus_authservice2"
  servers = [
    {
      address        = "2.2.3.3"
      auth_type      = "CHAP"
      disable        = false
      port           = 49
      use_accounting = false
      use_mgmt_port  = false
      shared_secret  = "test"
    }
  ]
  comment      = "Example TACACS Plus Auth Service."
  disable      = false
  acct_retries = 2
  acct_timeout = 2300
  auth_retries = 2
  auth_timeout = 7000
}
