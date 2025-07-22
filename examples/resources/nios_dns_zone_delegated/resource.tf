// Create a DNS zone delegated with basic fields
resource "nios_dns_zone_delegated" "zone_forward_basic_fields" {
  fqdn              = "zone-delegated.example.com"
  delegate_to = [
      {
        name    = "ns1.example.com"
        address = "1.1.1.1"
      }
    ]
}

// Create a DNS zone delegated with additional fields
resource "nios_dns_zone_delegated" "zone_forward_additional_fields" {
  fqdn = "zone-delegated-example2.example.com"
  delegate_to = [
    {
      name    = "ns2.example.com"
      address = "2.2.2.2"
    }
  ]

  // Additional fields
  comment = "This is a delegated zone for example.com"
  delegated_ttl    = 3600
  use_delegated_ttl = true
  ms_ad_integrated = false
  ms_ddns_mode = "ANY"
  prefix = "example-prefix"
  zone_format = "FORWARD"
  view = "default"

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}
