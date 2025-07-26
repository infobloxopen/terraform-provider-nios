// Retrieve a specific IPAM IPv6 network container using filters
data "nios_ipam_ipv6network_container" "get_network_containers_using_filters" {
  filters = {
    "network" = "10::/64"
  }
}

// Retrieve specific IPAM IPv6 network containers using Extensible Attributes
data "nios_ipam_ipv6network_container" "get_network_containers_using_extensible_attributes" {
  extattrfilters = {
    "Site" = "location-1"
  }
}

// Retrieve all IPAM IPv6 network containers
data "nios_ipam_ipv6network_container" "get_all_network_containers" {}
