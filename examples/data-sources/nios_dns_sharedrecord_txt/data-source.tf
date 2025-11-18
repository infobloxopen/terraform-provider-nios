// Retrieve a specific Shared TXT Record by filters
data "nios_dns_sharedrecord_txt" "get_sharedrecord_txt_using_filters" {
  filters = {
    name = "example-shared-record-txt"
  }
}

// Retrieve specific Shared TXT Records using Extensible Attributes
data "nios_dns_sharedrecord_txt" "get_sharedrecord_txt_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Shared TXT records
data "nios_dns_sharedrecord_txt" "get_all_sharedrecord_txt" {}
