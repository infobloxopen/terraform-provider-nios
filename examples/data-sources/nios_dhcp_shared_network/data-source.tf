// Retrieve a specific SharedNetwork record by filters
data "nios_dhcp_sharednetwork" "get_shared_network_using_filters" {
  filters = {
    name = "example_shared_network1"
  }
}

// Retrieve specific SharedNetworks using Extensible Attributes
data "nios_dhcp_sharednetwork" "get_shared_network_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all SharedNetworks
data "nios_dhcp_sharednetwork" "get_all_shared_networks_in_default_view" {}
