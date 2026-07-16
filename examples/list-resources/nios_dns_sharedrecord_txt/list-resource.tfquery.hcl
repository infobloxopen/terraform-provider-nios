// List specific Shared TXT Records using filters
list "nios_dns_sharedrecord_txt" "list_sharedrecord_txt_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example-shared-record-txt"
    }
  }
}

// List specific Shared TXT Records using Extensible Attributes
list "nios_dns_sharedrecord_txt" "list_sharedrecord_txt_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Shared TXT Records with resource details included
list "nios_dns_sharedrecord_txt" "list_sharedrecord_txt_with_resource" {
  provider         = nios
  include_resource = true
}
