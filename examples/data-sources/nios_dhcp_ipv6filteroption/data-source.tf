// Retrieve a specific DHCP IPv6 Filter Option by filters
data "nios_dhcp_ipv6filteroption" "get_dhcp_ipv6filteroptions_using_filters" {
  filters = {
    name = "example_ipv6_filter_option_1"
  }
}

// Retrieve specific DHCP IPv6 Option Filters using Extensible Attributes
data "nios_dhcp_ipv6filteroption" "get_dhcp_ipv6filteroption_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DHCP IPv6 Option Filters
data "nios_dhcp_ipv6filteroption" "get_all_dhcp_ipv6filteroptions" {}
