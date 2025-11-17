// Retrieve a specific RPZ A record by filters
data "nios_rpz_record_rpz_a" "get_record_using_filters" {
  filters = {
    name = "record1.rpz.example.com"
  }
}

// Retrieve specific RPZ A records using Extensible Attributes
data "nios_rpz_record_rpz_a" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ A records
data "nios_rpz_record_rpz_a" "get_all_rpz_a_records" {}

