// List specific PTR Records using filters
list "nios_dns_record_ptr" "list_ptr_records_using_filters" {
  provider = nios
  config {
    filters = {
      ptrdname = "example_record1.example.com"
    }
  }
}

// List specific PTR Records using Extensible Attributes
list "nios_dns_record_ptr" "list_ptr_records_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List PTR Records with resource details included
list "nios_dns_record_ptr" "list_ptr_records_with_resource" {
  provider         = nios
  include_resource = true
}
