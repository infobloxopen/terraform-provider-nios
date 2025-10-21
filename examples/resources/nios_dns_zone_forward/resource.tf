//Create an NS Group forward stub server (Required as parent for forward zones)
resource "nios_dns_nsgroup_forwardstubserver" "nsgroup_forward_stub" {
  name = "example_ns_group_forward_stub_server"
  external_servers = [
    {
      name    = "example.com"
      address = "2.3.4.4"
    }
  ]
}
// Create a DNS zone forward with basic fields
resource "nios_dns_zone_forward" "zone_forward_basic_fields" {
  fqdn              = "example1.example.com"
  external_ns_group = nios_dns_nsgroup_forwardstubserver.nsgroup_forward_stub.name
}

// Create a DNS zone forward with additional fields
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

// Create an IPV4 DNS zone forward
resource "nios_dns_zone_forward" "zone_forward_ipv4" {
  fqdn = "192.1.0.0/24"
  forward_to = [
    {
      name    = "ns1.example.com"
      address = "1.1.1.1"
    }
  ]
  zone_format = "IPV4"
}

// Create an IPV6 DNS zone forward
resource "nios_dns_zone_forward" "zone_forward_ipv6_mapping" {
  fqdn = "3001:db8::/64"
  forward_to = [
    {
      name    = "ns1.example.com"
      address = "1.1.1.1"
    }
  ]
  zone_format = "IPV6"
}
