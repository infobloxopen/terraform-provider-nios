// List specific Delegated Zones using filters
list "nios_dns_zone_delegated" "list_delegated_zones_using_filters" {
  provider = nios
  config {
    filters = {
      fqdn = "zone-delegated.example_auth.com"
    }
  }
}

// List specific Delegated Zones using Extensible Attributes
list "nios_dns_zone_delegated" "list_delegated_zones_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Delegated Zones with resource details included
list "nios_dns_zone_delegated" "list_delegated_zones_with_resource" {
  provider         = nios
  include_resource = true
}
