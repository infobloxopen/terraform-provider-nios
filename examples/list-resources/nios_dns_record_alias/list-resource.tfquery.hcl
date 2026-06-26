// List specific Alias Records using filters
list "nios_dns_record_alias" "list_alias_records_using_filters" {
  provider = nios
  config {
    filters = {
      name = "alias-record.example.com"
    }
  }
}

// List specific Alias Records using Extensible Attributes
list "nios_dns_record_alias" "list_alias_records_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Alias Records with resource details included
list "nios_dns_record_alias" "list_alias_records_with_resource" {
  provider         = nios
  include_resource = true
}
