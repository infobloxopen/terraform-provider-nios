// Retrieve a specific Fixed Addresses by filters
data "nios_dhcp_fixed_address" "get_record_using_filters" {
  filters = {
    name = "example_fixed_address"
  }
}

// Retrieve specific Fixed Addresses using Extensible Attributes
data "nios_dhcp_fixed_address" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Fixed Addresses
data "nios_dhcp_fixed_address" "get_all_records_in_default_view" {}
