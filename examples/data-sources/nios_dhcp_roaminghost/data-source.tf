// Retrieve a specific DHCP Roaming Host by filters
data "nios_dhcp_roaminghost" "get_dhcp_roaminghost_using_filters" {
  filters = {
    name = "roaming-host-1"
  }
}
// Retrieve specific DHCP Roaming Hosts using Extensible Attributes
data "nios_dhcp_roaminghost" "get_dhcp_roaminghosts_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DHCP Roaming Hosts
data "nios_dhcp_roaminghost" "get_all_dhcp_roaminghosts" {}
