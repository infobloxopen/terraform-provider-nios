// Retrieve a specific IPAM IPv6 network using filters
data "nios_ipam_ipv6network" "get_ipv6network_using_filters" {
  filters = {
    network = "10::/64"
  }
}

// Retrieve specific IPAM IPv6 networks using Extensible Attributes
data "nios_ipam_ipv6network" "get_ipv6network_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all IPAM IPv6 networks
data "nios_ipam_ipv6network" "get_all_ipv6networks" {}
