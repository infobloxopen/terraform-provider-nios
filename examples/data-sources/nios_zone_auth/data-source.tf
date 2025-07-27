// Retrieve a specific Auth Zone using filters
data "nios_dns_zone_auth" "get_auth_zone_using_filters" {
  filters = {
    view = "default"
    fqdn = "example1.com"
  }
}

// Retrieve Auth Zones using Extensible Attributes
data "nios_dns_zone_auth" "get_zones_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Auth Zones
data "nios_dns_zone_auth" "get_all_auth_zones" {}
