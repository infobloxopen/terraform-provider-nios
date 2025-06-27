// Retrieve a specific A record by name
data "nios_dns_record_a" "get_record_using_filters" {
  filters = {
    "name" = "example_record.example.com"
  }
}

// Retrieve specific A records using Extensible Attributes
data "nios_dns_record_a" "get_record_using_extensible_attributes" {
  extattrfilters = {
    "Site" = "location-1"
  }
}

// Retrieve all A records
data "nios_dns_record_a" "get_all_records_in_default_view" {}
