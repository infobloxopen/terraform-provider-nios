// List specific Shared A Records using filters
list "nios_dns_sharedrecord_a" "list_shared_a_records_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example_record.example.com"
    }
  }
}

// List specific Shared A Records using Extensible Attributes
list "nios_dns_sharedrecord_a" "list_shared_a_records_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Shared A Records with resource details included
list "nios_dns_sharedrecord_a" "list_shared_a_records_with_resource" {
  provider         = nios
  include_resource = true
}
