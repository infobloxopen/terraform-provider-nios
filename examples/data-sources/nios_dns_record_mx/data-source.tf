// Retrieve a specific MX record by filters
data "nios_dns_record_mx" "get_record_with_filter" {
  filters = {
    name = "mx_record.example.com"
  }
}

// Retrieve specific MX records using Extensible Attributes
data "nios_dns_record_mx" "get_record_with_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve MX records 
data "nios_dns_record_mx" "get_all_mx_records" {}
