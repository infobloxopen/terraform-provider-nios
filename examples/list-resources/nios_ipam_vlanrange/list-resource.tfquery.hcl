// List specific VLAN Ranges using filters
list "nios_ipam_vlanrange" "list_vlanranges_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example_vlan_range"
    }
  }
}

// List specific VLAN Ranges using Extensible Attributes
list "nios_ipam_vlanrange" "list_vlanranges_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List VLAN Ranges with resource details included
list "nios_ipam_vlanrange" "list_vlanranges_with_resource" {
  provider         = nios
  include_resource = true
}

