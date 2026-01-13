// Retrieve a specific RPZ CNAME client IP address record by filters
data "nios_rpz_record_cname_clientipaddress" "get_record_using_filters" {
  filters = {
    name = "12.0.0.1.rpzclientipaddress.example.com"
  }
}

// Retrieve specific RPZ CNAME client IP address records using Extensible Attributes
data "nios_rpz_record_cname_clientipaddress" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ CNAME client IP address records
data "nios_rpz_record_cname_clientipaddress" "get_all_rpz_cname_clientipaddress_records" {}
