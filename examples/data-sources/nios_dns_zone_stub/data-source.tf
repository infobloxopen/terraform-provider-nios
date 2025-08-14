// Retrieve a specific DNS zone stub by filters
data "nios_dns_zone_stub" "get_zone_stub_using_filters" {
  filters = {
    fqdn = "example_stub_zone.example.com"
  }
}

// Retrieve specific DNS zone stub using Extensible Attributes
data "nios_dns_zone_stub" "get_zone_stub_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DNS zone stub
data "nios_dns_zone_stub" "get_all_zone_stub" {}
