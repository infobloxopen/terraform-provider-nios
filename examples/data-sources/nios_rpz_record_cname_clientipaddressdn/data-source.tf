// Retrieve a specific RPZ CNAME Client IP Address DN record by filters
data "nios_rpz_record_cname_clientipaddressdn" "get_record_using_filters" {
  filters = {
    name = "10.10.0.1.rpz.example.com"
  }
}

// Retrieve specific RPZ CNAME Client IP Address DN records using Extensible Attributes
data "nios_rpz_record_cname_clientipaddressdn" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ CNAME Client IP Address DN records
data "nios_rpz_record_cname_clientipaddressdn" "get_all_rpz_cname_clientipaddressdn_records" {}
