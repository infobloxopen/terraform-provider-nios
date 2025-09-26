// Retrieve a specific Host record by filters
data "nios_record_host" "get_record_using_filters" {
  filters = {
    name = "host1.example.com"
  }
}

// Retrieve specific Host records using Extensible Attributes
data "nios_record_host" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Host records
data "nios_record_host" "get_all_records_in_default_view" {}
