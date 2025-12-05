// Retrieve a specific DHCP Roaminghost by filters
data "nios_dhcp_roaminghost" "get_dhcp_roaminghost_using_filters" {
  filters = {
    name = "roaming-host-1"
  }
}
// Retrieve specific DHCP Roaminghosts using Extensible Attributes
data "nios_dhcp_roaminghost" "get_dhcp_roaminghost_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DHCP Roaminghosts
data "nios_dhcp_roaminghost" "get_all_dhcp_roaminghost" {}
