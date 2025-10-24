// Create an Auth Zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_auth_zone" {
  fqdn = "example_auth.com"
  view = "default"
}

// Create a Reverse Mapping IPV4 DNS Zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_auth_reverse_zone" {
  fqdn        = "111.0.0.0/24"
  view        = "default"
  zone_format = "IPV4"
  extattrs = {
    Site = "location-1"
  }
}

// Create a DNS zone delegated with basic fields
resource "nios_dns_zone_delegated" "zone_delegated_basic_fields" {
  fqdn = "zone-delegated.example_auth.com"
  delegate_to = [
    {
      name    = "ns1.example.com"
      address = "10.10.10.10"
    }
  ]
  depends_on = [nios_dns_zone_auth.parent_auth_zone]
}

// Create a DNS zone delegated with IPv4 mapping
resource "nios_dns_zone_delegated" "zone_delegated_ip4_mapping" {
  fqdn = "111.0.0.10/32"
  delegate_to = [
    {
      name    = "ns2.example.com"
      address = "2.2.2.2"
    }
  ]
  zone_format = "IPV4"
  depends_on  = [nios_dns_zone_auth.parent_auth_reverse_zone]

}

// Create a DNS zone delegated with additional fields
resource "nios_dns_zone_delegated" "zone_delegated_additional_fields" {
  fqdn = "zone-delegated-2.example_auth.com"
  delegate_to = [
    {
      name    = "ns2.example.com"
      address = "20.20.20.20"
    }
  ]

  // Additional fields
  comment           = "This is a delegated zone for example.com"
  delegated_ttl     = 3600
  use_delegated_ttl = true
  ms_ad_integrated  = false
  ms_ddns_mode      = "ANY"
  prefix            = "example-prefix"
  zone_format       = "FORWARD"
  view              = "default"

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
  depends_on = [nios_dns_zone_auth.parent_auth_zone]
}
