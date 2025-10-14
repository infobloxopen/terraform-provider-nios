// Retrieve a specific TXT record by filters
data "nios_dns_record_txt" "get_record_using_filters" {
  filters = {
    name = "example-txt-record.example.com"
  }
}

// Retrieve specific TXT records using Extensible Attributes
data "nios_dns_record_txt" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all TXT records
data "nios_dns_record_txt" "get_all_txt_records" {}
