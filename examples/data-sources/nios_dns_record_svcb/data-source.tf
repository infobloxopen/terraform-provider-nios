// Retrieve a specific SVCB record by filters
data "nios_dns_record_svcb" "get_record_using_filters" {
  filters = {
    name = "example-svcb-record.example.com"
  }
}

// Retrieve specific SVCB records using Extensible Attributes
data "nios_dns_record_svcb" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all SVCB records
data "nios_dns_record_svcb" "get_all_svcb_records" {}
