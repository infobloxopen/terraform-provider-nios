// Retrieve a specific PTR record by filters
data "nios_dns_record_ptr" "get_record_using_filters" {
  filters = {
    ipv4addr = "10.20.1.2"
  }
}

// Retrieve specific PTR records using Extensible Attributes
data "nios_dns_record_ptr" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all PTR records
data "nios_dns_record_ptr" "get_all_records" {}
