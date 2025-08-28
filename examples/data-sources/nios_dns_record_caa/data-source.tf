// Retrieve a specific CAA record by filters
data "nios_dns_record_caa" "get_record_caa__with_filter" {
  filters = {
    name = "caa_record.example.com"
  }
}

// Retrieve specific CAA records using Extensible Attributes
data "nios_dns_record_caa" "get_record_caa_with_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all CAA records in the default view
data "nios_dns_record_caa" "get_all_caa_records_in_default_view" {}
