// Retrieve a specific IPAM network container using filters
data "nios_ipam_network" "get_network_using_filters" {
  filters = {
    "network" = "10.0.0.0/24"
  }
}

// Retrieve specific IPAM network using Extensible Attributes
data "nios_ipam_network" "get_network_using_extensible_attributes" {
  extattrfilters = {
    "Site" = "location-1"
  }
}

// Retrieve all IPAM network
data "nios_ipam_network" "get_all_network" {}
