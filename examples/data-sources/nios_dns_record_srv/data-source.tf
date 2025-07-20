// Retrieve a specific SRV record by filters
data "nios_dns_record_srv" "get_record_using_filters" {
  filters = {
    name = "example-srv-record.example.com"
  }
}

// Retrieve specific SRV records using Extensible Attributes
data "nios_dns_record_srv" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all SRV records
data "nios_dns_record_srv" "get_all_records" {}
