// Retrieve a specific DHCP Ipv6 Shared Network by filters
data "nios_dhcp_ipv6sharednetwork" "get_dhcp_ipv6sharednetwork_using_filters" {
  filters = {
    name = "shared_network_1"
  }
}

// Retrieve specific DHCP Ipv6 Shared Networks using Extensible Attributes
data "nios_dhcp_ipv6sharednetwork" "get_dhcp_ipv6sharednetworks_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DHCP Ipv6 Shared Networks
data "nios_dhcp_ipv6sharednetwork" "get_all_dhcp_ipv6sharednetworks" {}
