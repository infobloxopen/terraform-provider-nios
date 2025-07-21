// Retrieve a specific Range template record by filters
data "nios_dhcp_rangetemplate" "get_record_using_filters" {
  filters = {
    name = "example_range_template"
  }
}

// Retrieve specific Range template records using Extensible Attributes
data "nios_dhcp_rangetemplate" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Range template records
data "nios_dhcp_rangetemplate" "get_all_records" {}
