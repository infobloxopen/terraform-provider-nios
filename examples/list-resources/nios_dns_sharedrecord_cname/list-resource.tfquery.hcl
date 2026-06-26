// List specific Shared CNAME Records using filters
list "nios_dns_sharedrecord_cname" "list_shared_cname_records_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example-shared-record-cname"
    }
  }
}

// List specific Shared CNAME Records using Extensible Attributes
list "nios_dns_sharedrecord_cname" "list_shared_cname_records_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Shared CNAME Records with resource details included
list "nios_dns_sharedrecord_cname" "list_shared_cname_records_with_resource" {
  provider         = nios
  include_resource = true
}
