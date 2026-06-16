// List specific AAAA Records using filters
list "nios_dns_record_aaaa" "list_aaaa_records_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example_record.example.com"
    }
  }
}

// List specific AAAA Records using Extensible Attributes
list "nios_dns_record_aaaa" "list_aaaa_records_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List AAAA Records with resource details included
list "nios_dns_record_aaaa" "list_aaaa_records_with_resource" {
  provider         = nios
  include_resource = true
}
