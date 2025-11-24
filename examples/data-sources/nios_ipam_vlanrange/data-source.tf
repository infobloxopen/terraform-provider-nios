// Retrieve a specific IPAM Vlan Ranges by filters
data "nios_ipam_vlanrange" "get_ipam_vlanrange_using_filters" {
  filters = {
    name = "example_vlan_range"
  }
}
// Retrieve specific IPAM Vlan Ranges using Extensible Attributes
data "nios_" "get_ipam_vlanranges_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all IPAM Vlan Ranges
data "nios_ipam_vlanrange" "get_all_ipam_vlanranges" {}