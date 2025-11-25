// Manage IPAM Vlan Views (Required as Parent)
resource "nios_ipam_vlanview" "ipam_vlanview_parent" {
  start_vlan_id = 1
  end_vlan_id   = 100
  name          = "example_vlan_view"
}

// Manage IPAM Vlan Range with Basic Fields
resource "nios_ipam_vlanrange" "ipam_vlanrange_basic" {
  start_vlan_id = 5
  end_vlan_id   = 10
  name          = "example_vlan_range"
  vlan_view     = nios_ipam_vlanview.ipam_vlanview_parent.ref
}

// Manage IPAM Vlan Range with Additional Fields
resource "nios_ipam_vlanrange" "ipam_vlanrange_with_additional_fields" {
  start_vlan_id = 50
  end_vlan_id   = 100
  name          = "example_vlan_range2"
  vlan_view     = nios_ipam_vlanview.ipam_vlanview_parent.ref

  // Additional Fields
  comment          = "Example VLAN Range"
  pre_create_vlan  = true
  vlan_name_prefix = "vlan_range_"

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}
