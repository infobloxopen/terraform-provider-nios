// List specific Stub Zones using filters
list "nios_dns_zone_stub" "list_stub_zones_using_filters" {
  provider = nios
  config {
    filters = {
      fqdn = "example_stub_zone.example.com"
    }
  }
}

// List specific Stub Zones using Extensible Attributes
list "nios_dns_zone_stub" "list_stub_zones_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Stub Zones with resource details included
list "nios_dns_zone_stub" "list_stub_zones_with_resource" {
  provider         = nios
  include_resource = true
}
