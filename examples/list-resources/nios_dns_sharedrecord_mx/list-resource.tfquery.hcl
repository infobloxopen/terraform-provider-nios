// List specific Shared MX Records using filters
list "nios_dns_sharedrecord_mx" "list_shared_mx_records_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example_record.example.com"
    }
  }
}

// List specific Shared MX Records using Extensible Attributes
list "nios_dns_sharedrecord_mx" "list_shared_mx_records_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Shared MX Records with resource details included
list "nios_dns_sharedrecord_mx" "list_shared_mx_records_with_resource" {
  provider         = nios
  include_resource = true
}
