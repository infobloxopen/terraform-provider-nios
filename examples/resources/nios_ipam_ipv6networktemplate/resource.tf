// Manage an IPAM IPv6 Network Template with Basic Fields
resource "nios_ipam_ipv6networktemplate" "example_network" {
  name      = "example-network-template"
  cidr = 24

}

// Manage an IPAM IPv6 Network Template with Additional Fields
resource "nios_ipam_ipv6network" "complete_example" {
  // Required attributes
  name      = "example-network-template2"
  cidr = 24

  // Basic configuration
  network_view = "default"
  comment      = "Example IPv6 network template"

  ddns_enable_option_fqdn    = true
  ddns_generate_hostname     = true
  ddns_server_always_updates = true
  ddns_ttl                   = 0
  disable                    = true

  enable_ddns             = true
  enable_ifmap_publishing = true
  extattrs = {
    Site = "location-1"
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
}
