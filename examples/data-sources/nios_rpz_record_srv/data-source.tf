// Retrieve a specific RPZ SRV record by filters
data "nios_rpz_record_srv" "get_record_using_filters" {
  filters = {
    name = "record1.rpz.example.com"
  }
}

// Retrieve specific RPZ SRV records using Extensible Attributes
data "nios_rpz_record_srv" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ SRV records
data "nios_rpz_record_srv" "get_all_rpz_srv_records" {}
