// Retrieve a specific RPZ CNAME IP address record by filters
data "nios_rpz_record_cname_ipaddress" "get_record_using_filters" {
  filters = {
    name = "11.0.0.4.rpzip.example.com"
  }
}

// Retrieve specific RPZ CNAME IP address records using Extensible Attributes
data "nios_rpz_record_cname_ipaddress" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ CNAME IP address records
data "nios_rpz_record_cname_ipaddress" "get_all_rpz_cname_records" {}
