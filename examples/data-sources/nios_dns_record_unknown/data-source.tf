// Retrieve a specific Unknown record by filters
data "nios_dns_record_unknown" "get_unknown_record_using_filters" {
  filters = {
    name = "record-spf.example.com"
  }
}

// Retrieve specific Unknown records using Extensible Attributes
data "nios_dns_record_unknown" "get_unknown_records_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Unknown records in Default View
data "nios_dns_record_unknown" "get_all_unknown_records_in_default_view" {}
