// List specific CAA Records using filters
list "nios_dns_record_caa" "list_caa_records_using_filters" {
  provider = nios
  config {
    filters = {
      name = "caa-record.example.com"
    }
  }
}

// List specific CAA Records using Extensible Attributes
list "nios_dns_record_caa" "list_caa_records_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List CAA Records with resource details included
list "nios_dns_record_caa" "list_caa_records_with_resource" {
  provider         = nios
  include_resource = true
}
