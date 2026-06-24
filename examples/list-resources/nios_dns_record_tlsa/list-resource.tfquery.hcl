// List specific TLSA Records using filters
list "nios_dns_record_tlsa" "list_tlsa_records_using_filters" {
  provider = nios
  config {
    filters = {
      name = "record-tlsa.example.com"
    }
  }
}

// List specific TLSA Records using Extensible Attributes
list "nios_dns_record_tlsa" "list_tlsa_records_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List TLSA Records with resource details included
list "nios_dns_record_tlsa" "list_tlsa_records_with_resource" {
  provider         = nios
  include_resource = true
}
