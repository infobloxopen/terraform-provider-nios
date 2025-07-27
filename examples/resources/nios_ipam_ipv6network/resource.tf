// Create an IPAM IPv6 Network with Basic Fields
resource "nios_ipam_ipv6network" "example_network" {
  network      = "10::/64"
  network_view = "default"
  comment      = "Created by Terraform"

  // Optional: Configure extensible attributes
  extattrs = {
    "Site" = "location-1"
  }
}

// Create an IPAM IPv6 Network with Additional Fields
resource "nios_ipam_ipv6network" "complete_example" {
  // Required attributes
  network = "11::/64"

  // Basic configuration
  network_view = "default"
  comment      = "Complete network container example with all possible writable attributes"

  ddns_enable_option_fqdn    = true
  ddns_generate_hostname     = true
  ddns_server_always_updates = true
  ddns_ttl                   = 0
  disable                    = true

  enable_ddns             = true
  enable_ifmap_publishing = true
  extattrs = {
    "Site" = "location-1"
  }

  options = [
    {
      name         = "dhcp6.fqdn",
      num          = 39,
      value        = "test_options.com",
      vendor_class = "DHCPv6"
    }
  ]
  port_control_blackout_setting = {
    enable_blackout = false
  }
  preferred_lifetime          = 27000
  recycle_leases              = true
  update_dns_on_lease_renewal = true
  valid_lifetime              = 43200

  use_ddns_enable_option_fqdn     = true
  use_ddns_generate_hostname      = true
  use_blackout_setting            = true
  use_ddns_ttl                    = true
  use_enable_ddns                 = true
  use_enable_ifmap_publishing     = true
  use_options                     = true
  use_preferred_lifetime          = true
  use_recycle_leases              = true
  use_update_dns_on_lease_renewal = true
  use_valid_lifetime              = true
}

// Create an IPAM IPv6 Network with Function Call
resource "nios_ipam_ipv6network" "example_func_call" {
  func_call = {
    attribute_name  = "network"
    object_function = "next_available_network"
    result_field    = "networks"
    object          = "ipv6network"
    object_parameters = {
      network      = "10::/64"
      network_view = "default"
    }
    parameters = {
      cidr = 72
    }
  }
  comment = "Network container created with function call"
  depends_on = [
    nios_ipam_ipv6network.example_network,
  ]
}
