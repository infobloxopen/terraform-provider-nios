// Retrieve a specific DHCP IPv6 Range by filters
data "nios_dhcp_ipv6range" "get_ipv6range_with_filter" {
  filters = {
    start_addr = "15::10"
    end_addr   = "15::20"
  }
}

// Retrieve specific IPv6 DHCP Ranges using Extensible Attributes
data "nios_dhcp_ipv6range" "get_ipv6ranges_with_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all IPv6 DHCP Ranges
data "nios_dhcp_ipv6range" "get_all_ipv6ranges" {}
