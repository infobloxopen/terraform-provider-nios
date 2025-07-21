// Retrieve a specific alias record by filters
data "nios_dns_record_alias" "get_record_with_filter" {
  filters = {
    name = "alias-record.example.com"
  }
}

// Retrieve specific alias records using Extensible Attributes
data "nios_dns_record_alias" "get_record_with_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all alias records
data "nios_dns_record_alias" "get_all_records_in_default_view" {}
