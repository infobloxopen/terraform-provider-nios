// Retrieve a specific DHCP MAC Filter Address by filters
data "nios_dhcp_macfilteraddress" "get_mac_filter_address_using_filters" {
  filters = {
    filter = "mac_filter_example"
  }
}
// Retrieve specific DHCP MAC Filter Addresses using extensible attributes
data "nios_dhcp_macfilteraddress" "get_mac_filter_address_using_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DHCP MAC Filter Addresses
data "nios_dhcp_macfilteraddress" "get_all_mac_filter_addresses" {
}
