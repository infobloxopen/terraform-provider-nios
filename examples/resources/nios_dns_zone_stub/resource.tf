// Create a DNS zone stub with basic fields
resource "nios_dns_zone_stub" "zone_stub_basic_fields" {
  fqdn              = "example_stub_zone.example.com"
  stub_from = [
      {
        name    = "stub.example.com"
        address = "1.1.1.1"
      }
    ]
}

// Create a DNS zone stub with additional fields
resource "nios_dns_zone_stub" "zone_stub_additional_fields" {
  fqdn = "example_stub_zone2.example.com"
  stub_from = [
        {
          name    = "stub.example.com"
          address = "1.1.1.1"
        }
      ]

  // Additional Fields
  ms_ddns_mode = "ANY"
  prefix = "stub-prefix"
  comment = "This is a stub zone with additional fields"
  view = "default"
  extattrs = {
    Site = "location-1"
  }
}

// Create an IPV4 DNS zone stub
resource "nios_dns_zone_stub" "zone_stub_ipv4" {
  fqdn = "10.1.0.0/25"
  forward_to = [
    {
      name    = "stub.example.com"
      address = "1.1.1.1"
    }
  ]
  zone_format = "IPV4"
}

// Create an IPV6 DNS zone stub
resource "nios_dns_zone_stub" "zone_stub_ipv6_mapping" {
  fqdn = "3001:db8::/64"
  forward_to = [
    {
      name    = "stub.example.com"
      address = "1.1.1.1"
    }
  ]
  zone_format = "IPV6"
}
