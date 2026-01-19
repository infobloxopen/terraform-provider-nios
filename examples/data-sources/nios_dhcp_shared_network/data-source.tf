// Retrieve a specific Shared Network record by filters
data "nios_dhcp_shared_network" "get_shared_network_using_filters" {
  filters = {
    name = "example_shared_network1"
  }
}

// Retrieve specific Shared Networks using Extensible Attributes
data "nios_dhcp_shared_network" "get_shared_network_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Shared Networks
data "nios_dhcp_shared_network" "get_all_shared_networks" {}
