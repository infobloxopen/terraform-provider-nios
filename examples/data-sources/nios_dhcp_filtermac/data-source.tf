// Retrieve a specific DHCP Filter MAC by filters
data "nios_dhcp_filtermac" "get_filtermac_using_filters" {
  filters = {
    name = "mac_filter_example"
  }
}
// Retrieve specific DHCP MAC Filters using extensible attributes
data "nios_dhcp_filtermac" "get_filtermac_using_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}