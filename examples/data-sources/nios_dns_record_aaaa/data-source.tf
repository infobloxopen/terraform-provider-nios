// Retrieve a specific AAAA record by name
data "nios_dns_record_aaaa" "get_record_with_filter" {
  filters = {
    "name" = "example_record.example.com"
  }
}

// Retrieve specific AAAA records using Extensible Attributes
data "nios_dns_record_aaaa" "get_record_with_extattr_filter" {
  extattrfilters = {
    "Site" = "location-1"
  }
}

// Retrieve all AAAA records
data "nios_dns_record_aaaa" "get_all_records_in_default_view" {}