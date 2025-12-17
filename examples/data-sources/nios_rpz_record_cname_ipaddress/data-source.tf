// Retrieve a specific RPZ CNAME ipaddress record by filters and edit to match existing record
data "nios_rpz_record_rpz_cname_ipaddress" "get_record_using_filters" {
  filters = {
    name = "11.0.0.4.rpzip.example.com"
  }
}

// Retrieve specific RPZ CNAME ipaddress records using Extensible Attributes
data "nios_rpz_record_rpz_cname_ipaddress" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ CNAME ipaddress records
data "nios_rpz_record_rpz_cname_ipaddress" "get_all_rpz_cname_records" {}
