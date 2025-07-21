resource "nios_ipam_network" "example_network" {
  network      = "10.0.0.0/24"
  network_view = "default"
  comment      = "Created by Terraform"
  extattrs = {
    "Site" = "location-1"
  }
}

resource "nios_ipam_network" "complete_example" {
  // Required attributes
  network = "11.0.0.0/24"

  // Basic configuration
  network_view = "default"
  comment      = "Complete network container example with all possible writable attributes"

  # auto_create_reversezone = false
  bootfile       = "pxelinux.0"
  bootserver     = "192.168.1.10"
  use_authority  = true
  use_bootfile   = true
  use_bootserver = true

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
    "Site" = "DataCenter1"
  }
}


resource "nios_ipam_network" "example_func_call" {
  func_call = {
    attribute_name  = "network"
    object_function = "next_available_network"
    result_field    = "networks"
    object          = "network"
    object_parameters = {
      network      = "10.0.0.0/24"
      network_view = "default"
    }
    parameters = {
      cidr = 28
    }
  }
  comment = "Network container created with function call"
  depends_on = [
    nios_ipam_network.example_network
  ]
}
