// Retrieve a specific MX Shared Record by filters
data "nios_dns_sharedrecord_mx" "get_record_with_filter" {
  filters = {
    name = "sharedmx_record"
  }
}

// Retrieve specific MX Shared Records using Extensible Attributes
data "nios_dns_sharedrecord_mx" "get_record_with_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all MX Shared Records
data "nios_dns_sharedrecord_mx" "get_all_sharedmx_records" {}
