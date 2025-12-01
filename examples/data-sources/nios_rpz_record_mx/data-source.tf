// Retrieve a specific RPZ MX record by filters
data "nios_rpz_record_mx" "get_rpz_mx_record_using_filters" {
  filters = {
    name = "rpz_record_mx.rpz.example.com"
  }
}

// Retrieve specific RPZ MX records using Extensible Attributes
data "nios_rpz_record_mx" "get_rpz_mx_records_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ MX records
data "nios_rpz_record_mx" "get_all_rpz_mx_records" {}
