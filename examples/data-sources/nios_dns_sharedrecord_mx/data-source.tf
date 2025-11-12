// Retrieve a specific MX Shared Record by filters
data "nios_dns_sharedrecord_mx" "get_mx_sharedrecord_with_filter" {
  filters = {
    name = "sharedrecord_mx_basic"
  }
}

// Retrieve specific MX Shared Records using Extensible Attributes
data "nios_dns_sharedrecord_mx" "get_mx_sharedrecord_with_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all MX Shared Records
data "nios_dns_sharedrecord_mx" "get_all_mx_sharedrecords" {}
