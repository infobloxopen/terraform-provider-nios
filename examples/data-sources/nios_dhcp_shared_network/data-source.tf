// Retrieve a specific SharedNetwork record by filters
data "nios_dhcp_sharednetwork" "get_record_using_filters" {
  filters = {
    name = "example_shared_network1"
  }
}

// Retrieve specific SharedNetwork records using Extensible Attributes
data "nios_dhcp_sharednetwork" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all SharedNetwork records
data "nios_dhcp_sharednetwork" "get_all_records_in_default_view" {}
