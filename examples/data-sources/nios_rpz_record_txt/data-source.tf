// Retrieve a specific RPZ TXT record by filters
data "nios_rpz_record_txt" "get_record_using_filters" {
  filters = {
    name = "record1.rpz.example.com"
  }
}

// Retrieve specific RPZ TXT records using Extensible Attributes
data "nios_rpz_record_txt" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ TXT records
data "nios_rpz_record_txt" "get_all_rpz_txt_records" {}
