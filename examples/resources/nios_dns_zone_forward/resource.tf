// Create a DNS zone forward record with basic fields
resource "nios_dns_zone_forward" "zone_forward_basic_fields" {
  fqdn              = "example1.example.com"
  external_ns_group = "nsg1"
}

// Create a DNS zone forward record with additional fields
resource "nios_dns_zone_forward" "zone_forward_additional_fields" {
  fqdn = "example2.example.com"
  forward_to = [
    {
      name    = "ns1.example.com"
      address = "1.1.1.1"
    }
  ]
  forwarding_servers = [
    {
      name                    = "infoblox.172_28_82_248"
      forwarders_only         = true
      use_override_forwarders = true
      forward_to = [
        {
          name    = "kk.fwd.com"
          address = "10.2.1.31"
        }
      ]
    }
  ]
  view = "default"
  extattrs = {
    Site = "location-1"
  }
}
