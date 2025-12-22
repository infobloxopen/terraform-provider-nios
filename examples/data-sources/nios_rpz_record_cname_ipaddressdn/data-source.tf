// Retrieve a specific RPZ CNAME IP Address DN record by filters
data "nios_rpz_record_cname_ipaddressdn" "get_record_using_filters" {
  filters = {
    name = "record1.rpz.example.com"
  }
}

// Retrieve specific RPZ CNAME IP Address DN records using Extensible Attributes
data "nios_rpz_record_cname_ipaddressdn" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ CNAME IP Address DN records
data "nios_rpz_record_cname_ipaddressdn" "get_all_rpz_cname_ipaddressdn_records" {}
