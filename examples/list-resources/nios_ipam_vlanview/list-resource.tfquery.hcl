// List specific VLAN Views using filters
list "nios_ipam_vlanview" "list_vlanviews_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example_vlan_view"
    }
  }
}

// List specific VLAN Views using Extensible Attributes
list "nios_ipam_vlanview" "list_vlanviews_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List VLAN Views with resource details included
list "nios_ipam_vlanview" "list_vlanviews_with_resource" {
  provider         = nios
  include_resource = true
}
 