// Create Radius Authservice with Basic Fields
resource "nios_security_radius_authservice" "radius_authservice_with_basic_fields" {
  name = "radius_authservice1"
  servers = [
    {
      acct_port      = 1813
      address        = "2.2.3.1"
      auth_port      = 1812
      auth_type      = "PAP"
      disable        = false
      use_accounting = false
      use_mgmt_port  = false
      shared_secret  = "test"
    }
  ]
}

// Create Radius Authservice with Additional Fields
resource "nios_security_radius_authservice" "radius_authservice_with_additional_field" {
  name = "radius_authservice2"
  servers = [
    {
      acct_port      = 1813
      address        = "2.2.3.1"
      auth_port      = 1812
      auth_type      = "PAP"
      disable        = false
      use_accounting = false
      use_mgmt_port  = false
      shared_secret  = "test"
    }
  ]
  acct_retries      = 3000
  acct_timeout      = 2300
  auth_retries      = 8
  auth_timeout      = 1200
  cache_ttl         = 3000
  comment           = "This is a commment"
  disable           = false
  enable_cache      = true
  mode              = "ROUND_ROBIN"
  recovery_interval = 240
}
