// Retrieve a specific Shared SRV Record by filters
data "nios_dns_sharedrecord_srv" "get_sharedrecord_srv_with_filter" {
  filters = {
    name = "sharedrecord_srv.example.com"
  }
}

// Retrieve specific Shared SRV Records using Extensible Attributes
data "nios_dns_sharedrecord_srv" "get_sharedrecord_srv_with_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Shared SRV Records
data "nios_dns_sharedrecord_srv" "get_all_srv_sharedrecords" {}
