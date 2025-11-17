// Retrieve a specific IPAM Vlan by filters
data "nios_ipam_vlan" "get_ipam_vlan_using_filters" {
  filters = {
    name = "example_vlan"
  }
}
// Retrieve specific IPAM Vlan using Extensible Attributes
data "nios_" "get__using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all IPAM Vlan
data "nios_ipam_vlan" "get_all_ipam_vlan" {}
