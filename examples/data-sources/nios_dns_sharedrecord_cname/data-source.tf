// Retrieve a specific Shared CNAME Record by filters
data "nios_dns_sharedrecord_cname" "get_sharedrecord_cname_using_filters" {
  filters = {
    name = "example-shared-record-cname"
  }
}

// Retrieve specific Shared CNAME Records using Extensible Attributes
data "nios_dns_sharedrecord_cname" "get_sharedrecord_cname_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Shared CNAME records
data "nios_dns_sharedrecord_cname" "get_all_sharedrecord_cname" {}
