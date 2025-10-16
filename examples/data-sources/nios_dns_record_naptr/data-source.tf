// Retrieve a specific NAPTR record by filters
data "nios_dns_record_naptr" "get_record_with_filter" {
  filters = {
    name = "naptr_record.example.com"
  }
}

// Retrieve specific NAPTR records using Extensible Attributes
data "nios_dns_record_naptr" "get_record_with_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all NAPTR records 
data "nios_dns_record_naptr" "get_all_naptr_records" {}
