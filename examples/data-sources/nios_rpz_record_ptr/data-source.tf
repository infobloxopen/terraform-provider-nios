// Retrieve a specific RPZ PTR record by filters
data "nios_rpz_record_ptr" "get_record_using_filters" {
  filters = {
    ptrdname = "record1.rpz.example.com"
  }
}

// Retrieve specific RPZ PTR records using Extensible Attributes
data "nios_rpz_record_ptr" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ PTR records
data "nios_rpz_record_ptr" "get_all_rpz_ptr_records" {}
