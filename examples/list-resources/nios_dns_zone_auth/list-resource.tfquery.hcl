// List specific Authoritative Zones using filters
list "nios_dns_zone_auth" "list_auth_zones_using_filters" {
  provider = nios
  config {
    filters = {
      fqdn = "example.com"
    }
  }
}

// List specific Authoritative Zones using Extensible Attributes
list "nios_dns_zone_auth" "list_auth_zones_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Authoritative Zones with resource details included
list "nios_dns_zone_auth" "list_auth_zones_with_resource" {
  provider         = nios
  include_resource = true
}
