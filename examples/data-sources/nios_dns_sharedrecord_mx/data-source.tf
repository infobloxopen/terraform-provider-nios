// Retrieve a specific Shared MX record by filters
data "nios_dns_sharedrecord_mx" "get_record_with_filter" {
  filters = {
    name = "sharedmx_record"
  }
}

// Retrieve specific Shared MX records using Extensible Attributes
data "nios_dns_sharedrecord_mx" "get_record_with_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve Shared MX records
data "nios_dns_sharedrecord_mx" "get_all_sharedmx_records" {}
