// Retrieve a specific IPAM network container using filters
data "nios_ipam_ipv6network" "get_ipv6network_using_filters" {
  filters = {
    "network" = "10::/64"
  }
}

// Retrieve specific IPAM network containers using Extensible Attributes
data "nios_ipam_ipv6network" "get_ipv6network_using_extensible_attributes" {
  extattrfilters = {
    "Site" = "location-1"
  }
}

// Retrieve all IPAM network containers
data "nios_ipam_ipv6network" "get_all_ipv6network" {}
