// List specific Shared SRV Records using filters
list "nios_dns_sharedrecord_srv" "list_sharedrecord_srv_using_filters" {
  provider = nios
  config {
    filters = {
      name = "sharedrecord_srv.example.com"
    }
  }
}

// List specific Shared SRV Records using Extensible Attributes
list "nios_dns_sharedrecord_srv" "list_sharedrecord_srv_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Shared SRV Records with resource details included
list "nios_dns_sharedrecord_srv" "list_sharedrecord_srv_with_resource" {
  provider         = nios
  include_resource = true
}
