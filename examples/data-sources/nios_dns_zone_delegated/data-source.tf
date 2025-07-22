// Retrieve a specific DNS zone delegated record by filters
data "nios_dns_zone_delegated" "get_zones_using_filters" {
  filters = {
    fqdn = "zone-delegated.example.com"
  }
}

// Retrieve specific DNS zone delegated records using Extensible Attributes
data "nios_dns_zone_delegated" "get_zones_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DNS zone delegated zones
data "nios_dns_zone_delegated" "get_all_zones" {}
