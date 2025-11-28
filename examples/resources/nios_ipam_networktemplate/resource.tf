// Create an IPAM Network Template with Basic Fields
resource "nios_ipam_network" "ipam_network_template_basic" {
  name      = "example_network_template"
  netmask = 24
}

// Create an IPAM Network Template with Additional Fields
resource "nios_ipam_networktemplate" "ipam_network_template_with_additional_fields" {
  // Required attributes
  name      = "example_network_template2"

  comment      = "Example IPAM Network Template"

  auto_create_reversezone = false
  allow_any_netmask    = true
  bootfile                = "pxelinux.0"
  bootserver              = "192.168.1.10"
  use_authority           = true
  use_bootfile            = true
  use_bootserver          = true

  // DDNS settings
  enable_ddns                     = true
  use_enable_ddns                 = true
  ddns_domainname                 = "example.com"
  ddns_generate_hostname          = true
  ddns_ttl                        = 3600
  ddns_update_fixed_addresses     = true
  ddns_use_option81               = true
  use_ddns_domainname             = true
  use_ddns_generate_hostname      = true
  use_ddns_ttl                    = true
  use_ddns_update_fixed_addresses = true
  use_ddns_use_option81           = true

  // Email and notification settings
  email_list     = ["admin@example.com", "network@example.com"]
  use_email_list = true

  // Water mark settings
  high_water_mark       = 95
  high_water_mark_reset = 85
  low_water_mark        = 10
  low_water_mark_reset  = 20

  // Extensible attributes
  extattrs = {
    Site = "location-1"
  }
}

