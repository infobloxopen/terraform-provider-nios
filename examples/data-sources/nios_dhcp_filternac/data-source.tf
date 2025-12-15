// Retrieve a specific DHCP Filter NAC by filters
data "nios_dhcp_filternac" "get_filternac_using_filters" {
  filters = {
    name = "nac_filter_example"
  }
}
// Retrieve specific DHCP NAC Filters using extensible attributes
data "nios_dhcp_filternac" "get_filternac_using_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}
// Retrieve all DHCP NAC Filters
data "nios_dhcp_filternac" "get_all_filternacs" {
}
