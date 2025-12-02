// Retrieve a specific RPZ NAPTR record by filters
data "nios_rpz_record_naptr" "get_record_using_filters" {
  filters = {
    name = "record1.rpz.example.com"
  }
}

// Retrieve specific RPZ NAPTR records using Extensible Attributes
data "nios_rpz_record_naptr" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ NAPTR records
data "nios_rpz_record_naptr" "get_all_rpz_naptr_records" {}
