// Manage an IPAM IPv6 Network Template with Basic Fields
resource "nios_ipam_ipv6networktemplate" "ipam_ipv6_network_template_basic" {
  name = "example-network-template"
  cidr = 24

  // add cloud_api_compatible = true if Terraform Internal ID extensible attribute has cloud access
  cloud_api_compatible = false
}

// Manage an IPAM IPv6 Network Template with Additional Fields
resource "nios_ipam_ipv6networktemplate" "ipam_ipv6_network_template_with_additional_fields" {
  // Required attributes
  name = "example-network-template2"
  cidr = 24

  // Basic configuration
  comment = "Example IPv6 network template"

  ddns_enable_option_fqdn    = true
  ddns_generate_hostname     = true
  ddns_server_always_updates = true
  ddns_ttl                   = 100

  enable_ddns = true
  extattrs = {
    "Tenant ID" = "location-1"
  }

  options = [
    {
      name         = "dhcp6.fqdn",
      num          = 39,
      value        = "test_options.com",
      vendor_class = "DHCPv6"
    }
  ]
  preferred_lifetime          = 27000
  recycle_leases              = true
  update_dns_on_lease_renewal = true

  // add cloud_api_compatible = true if Terraform Internal ID extensible attribute has cloud access
  cloud_api_compatible = true

  use_ddns_enable_option_fqdn     = true
  use_ddns_generate_hostname      = true
  use_ddns_ttl                    = true
  use_enable_ddns                 = true
  use_options                     = true
  use_preferred_lifetime          = true
  use_recycle_leases              = true
  use_update_dns_on_lease_renewal = true
}
