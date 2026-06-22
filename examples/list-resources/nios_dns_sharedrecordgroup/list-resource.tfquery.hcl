// List specific Shared Record Groups using filters
list "nios_dns_sharedrecordgroup" "list_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example-shared-record-group"
    }
  }
}

// List specific Shared Record Groups using Extensible Attributes
list "nios_dns_sharedrecordgroup" "list_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Shared Record Groups with resource details included
list "nios_dns_sharedrecordgroup" "list_with_resource" {
  provider         = nios
  include_resource = true
}
