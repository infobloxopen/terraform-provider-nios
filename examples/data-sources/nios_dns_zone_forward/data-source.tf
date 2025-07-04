// Retrieve a specific zone forward record by name
data "nios_dns_zone_forward" "get_record_using_filters" {
  filters = {
    "fqdn" = "example_test1.example.com"
  }
}

// Retrieve specific zone forward records using Extensible Attributes
data "nios_dns_zone_forward" "get_record_using_extensible_attributes" {
  extattrfilters = {
    "Site" = "Hokkaido"
  }
}

// Retrieve all zone forward records
data "nios_dns_zone_forward" "get_all_records_in_default_view" {}
