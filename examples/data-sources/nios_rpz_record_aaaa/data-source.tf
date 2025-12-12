// Retrieve a specific RPZ AAAA record by filters
data "nios_rpz_record_aaaa" "get_record_using_filters" {
  filters = {
    name = "record1.rpz.example.com"
  }
}

// Retrieve specific RPZ AAAA records using Extensible Attributes
data "nios_rpz_record_aaaa" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ AAAA records
data "nios_rpz_record_aaaa" "get_all_rpz_aaaa_records" {}
