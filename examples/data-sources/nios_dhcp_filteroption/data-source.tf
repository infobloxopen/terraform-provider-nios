// Retrieve a specific DHCP Filteroption by filters
data "nios_dhcp_filteroption" "get_filteroption_using_filters" {
  filters = {
    name = "filteroption_example"
  }
}
// Retrieve specific DHCP Filteroptions using extensible attributes
data "nios_dhcp_filteroption" "get_filteroption_using_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}
// Retrieve all DHCP Filteroptions
data "nios_dhcp_filteroption" "get_all_filteroptions" {
}
