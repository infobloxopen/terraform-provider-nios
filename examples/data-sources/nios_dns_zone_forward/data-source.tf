// Retrieve a specific DNS zone forward by filters
data "nios_dns_zone_forward" "get_zone_forward_using_filters" {
  filters = {
    fqdn = "zone-forward1.example.com"
  }
}

// Retrieve specific DNS zone forward using Extensible Attributes
data "nios_dns_zone_forward" "get_zone_forward_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DNS zone forward
data "nios_dns_zone_forward" "get_all_zone_forward" {}
