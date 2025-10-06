// Retrieve a specific Zone RP using filters
data "nios_dns_zone_rp" "get_zone_rp_using_filters" {
  filters = {
    view = "default"
    fqdn = "example1.com"
  }
}

// Retrieve Zone RPs using Extensible Attributes
data "nios_dns_zone_rp" "get_zones_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Auth Zone RPs
data "nios_dns_zone_rp" "get_all_zone_rp" {}
