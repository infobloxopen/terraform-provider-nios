// Retrieve a specific TLSA record by filters
data "nios_dns_record_tlsa" "get_tlsa_record_using_filters" {
  filters = {
    name = nios_dns_record_tlsa.create_record_tlsa_with_basic_fields.name
  }
}

// Retrieve specific TLSA records using Extensible Attributes
data "nios_dns_record_tlsa" "get_tlsa_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all TLSA records in Default View
data "nios_dns_record_tlsa" "get_all_tlsa_records_in_default_view" {}
