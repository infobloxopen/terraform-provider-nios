// Retrieve a specific Shared MX Record by filters
data "nios_dns_sharedrecord_mx" "get_sharedrecord_mx_with_filter" {
  filters = {
    name = "sharedrecord_mx_basic"
  }
}

// Retrieve specific Shared MX Records using Extensible Attributes
data "nios_dns_sharedrecord_mx" "get_sharedrecord_mx_with_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Shared MX Records
data "nios_dns_sharedrecord_mx" "get_all_mx_sharedrecords" {}
