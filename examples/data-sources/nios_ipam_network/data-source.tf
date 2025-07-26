// Retrieve a specific IPAM network container using filters
data "nios_ipam_network" "get_network_using_filters" {
  filters = {
    "network" = "10.0.0.0/24"
  }
}

// Retrieve specific IPAM network containers using Extensible Attributes
data "nios_ipam_networkcontainer" "get_network_containers_using_extensible_attributes" {
  extattrfilters = {
    "Site" = "location-1"
  }
}

// Retrieve all IPAM network containers
data "nios_ipam_networkcontainer" "get_all_network_containers" {}
