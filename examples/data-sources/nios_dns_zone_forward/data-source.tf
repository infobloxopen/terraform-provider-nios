// Retrieve a specific DNS zone forward record by filters
data "nios_dns_zone_forward" "get_record_using_filters" {
  filters = {
    fqdn = "zone-forward1.example.com"
  }
}

// Retrieve specific DNS zone forward records using Extensible Attributes
data "nios_dns_zone_forward" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DNS zone forward records in default view
data "nios_dns_zone_forward" "get_all_records_in_default_view" {}
