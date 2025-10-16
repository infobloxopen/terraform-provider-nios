// Retrieve a specific DNAME record by filters
data "nios_dns_record_dname" "get_record_using_filters" {
  filters = {
    target = "example-dname-1.com"
  }
}

// Retrieve specific DNAME records using Extensible Attributes
data "nios_dns_record_dname" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DNAME records
data "nios_dns_record_dname" "get_all_dname_records" {}
