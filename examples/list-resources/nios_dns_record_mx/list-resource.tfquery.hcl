// List specific MX Records using filters
list "nios_dns_record_mx" "list_mx_records_using_filters" {
  provider = nios
  config {
    filters = {
      name = "mx-record.example.com"
    }
  }
}

// List specific MX Records using Extensible Attributes
list "nios_dns_record_mx" "list_mx_records_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List MX Records with resource details included
list "nios_dns_record_mx" "list_mx_records_with_resource" {
  provider         = nios
  include_resource = true
}
