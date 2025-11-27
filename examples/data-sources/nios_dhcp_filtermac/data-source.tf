// Retrieve a specific DHCP Filter MAC by filters
data "nios_dhcp_filtermac" "get_filtermac_using_filters" {
  filters = {
    name = "mac_filter_example"
  }
}
// Retrieve a specific nios dhcp macfilter by extensible attribute filters
data "nios_dhcp_filtermac" "get_filtermac_using_extattrfilters" {
  extattrfilters = {
    Site = "location-1"
  }
}