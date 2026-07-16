// List specific VLANs using filters
list "nios_ipam_vlan" "list_vlans_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example_vlan"
    }
  }
}

// List specific VLANs using Extensible Attributes
list "nios_ipam_vlan" "list_vlans_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List VLANs with resource details included
list "nios_ipam_vlan" "list_vlans_with_resource" {
  provider         = nios
  include_resource = true
}
