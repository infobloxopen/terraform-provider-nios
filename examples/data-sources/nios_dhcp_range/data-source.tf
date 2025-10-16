// Retrieve a specific DHCP range by filters
data "nios_dhcp_range" "get_range_with_filter" {
  filters = {
    start_addr = "10.0.0.170"
  }
}

// Retrieve specific DHCP ranges using Extensible Attributes
data "nios_dhcp_range" "get_range_with_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DHCP ranges
data "nios_dhcp_range" "get_all_ranges" {}
