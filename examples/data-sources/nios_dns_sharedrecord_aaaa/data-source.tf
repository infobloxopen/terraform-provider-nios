// Retrieve a specific Shared AAAA Record by filters
data "nios_dns_sharedrecord_aaaa" "get_sharedrecord_aaaa_with_filter" {
  filters = {
    name = "sharedrecord_aaaa_basic"
  }
}

// Retrieve specific Shared AAAA Records using Extensible Attributes
data "nios_dns_sharedrecord_aaaa" "get_sharedrecord_aaaa_with_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Shared AAAA Records
data "nios_dns_sharedrecord_aaaa" "get_all_shared_aaaa_records" {}
