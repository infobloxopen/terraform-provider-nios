// Retrieve a specific RPZ CNAME record by filters
data "nios_rpz_record_cname" "get_record_using_filters" {
  filters = {
    name = "record1.rpz.example.com"
  }
}

// Retrieve specific RPZ CNAME records using Extensible Attributes
data "nios_rpz_record_cname" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ CNAME records
data "nios_rpz_record_cname" "get_all_rpz_cname_records" {}
