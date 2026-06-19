// List specific Unknown Records using filters
list "nios_dns_record_unknown" "list_unknown_records_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example-unknown.example.com"
    }
  }
}

// List specific Unknown Records using Extensible Attributes
list "nios_dns_record_unknown" "list_unknown_records_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Unknown Records with resource details included
list "nios_dns_record_unknown" "list_unknown_records_with_resource" {
  provider         = nios
  include_resource = true
}
