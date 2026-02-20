// List specific A Records using filters
list "nios_dns_record_a" "list_records_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example_record.example.com"
    }
  }
}

// List specific A Records using Extensible Attributes
list "nios_dns_record_a" "list_records_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List A records with resource details included
list "nios_dns_record_a" "list_records_with_resource" {
  provider         = nios
  include_resource = true
}