// Retrieve a specific Range Template by filters
data "nios_dhcp_rangetemplate" "get_range_template_using_filters" {
  filters = {
    name = "example_range_template"
  }
}

// Retrieve specific Range Templates using Extensible Attributes
data "nios_dhcp_rangetemplate" "get_range_template_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Range Templates
data "nios_dhcp_rangetemplate" "get_all_range_templates" {}
