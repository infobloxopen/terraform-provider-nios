// Retrieve a specific Fixed Address Template by filters
data "nios_dhcp_fixedaddresstemplate" "get_fixed_address_template_using_filters" {
  filters = {
    name = "example_fixed_address_template_1"
  }
}

// Retrieve specific Fixed Address Templates using Extensible Attributes
data "nios_dhcp_fixedaddresstemplate" "get_fixed_address_template_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Fixed Address Templates
data "nios_dhcp_fixedaddresstemplate" "get_all_fixed_address_templates" {}
