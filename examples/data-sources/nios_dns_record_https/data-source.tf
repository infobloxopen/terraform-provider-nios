// Retrieve a specific HTTPS record by filters
data "nios_dns_record_https" "get_record_using_filters" {
  filters = {
    name = "example-https-record.example.com"
  }
}

// Retrieve specific HTTPS records using Extensible Attributes
data "nios_dns_record_https" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all HTTPS records
data "nios_dns_record_https" "get_all_https_records" {}