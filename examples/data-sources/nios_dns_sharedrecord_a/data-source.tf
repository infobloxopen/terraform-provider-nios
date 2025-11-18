// Retrieve a specific Shared A Record by filters
data "nios_dns_sharedrecord_a" "get_sharedrecord_a_with_filter" {
  filters = {
    name = "sharedrecord_a_basic"
  }
}

// Retrieve specific Shared A Records using Extensible Attributes
data "nios_dns_sharedrecord_a" "get_sharedrecord_a_with_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}
