// Retrieve a specific CNAME record by filters
data "nios_dns_record_cname" "get_record_using_filters" {
  filters = {
    name = "example_record.example.com"
  }
}

// Retrieve specific CNAME records using Extensible Attributes
data "nios_dns_record_cname" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all CNAME records
data "nios_dns_record_cname" "get_all_records_in_default_view" {}
