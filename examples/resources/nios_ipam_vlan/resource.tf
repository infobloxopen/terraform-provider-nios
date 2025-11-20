// Manage IPAM Vlan Views (Required as Parent)
resource "nios_ipam_vlanview" "ipam_vlanview_parent" {
  start_vlan_id = 5
  end_vlan_id   = 10
  name          = "example_vlan_view"
}

// Manage  IPAM Vlan with Basic Fields
resource "nios_ipam_vlan" "ipam_vlan_basic" {
  id     = 6
  name   = "example_vlan"
  parent = nios_ipam_vlanview.ipam_vlanview_parent.ref
}

// Manage  IPAM Vlan with Additional Fields
resource "nios_ipam_vlan" "ipam_vlan_with_additional_fields" {
  id     = 7
  name   = "example_vlan_additional"
  parent = nios_ipam_vlanview.ipam_vlanview_parent.ref

  // Additional Fields
  comment     = "Example VLAN"
  contact     = "Infoblox"
  department  = "Engineering"
  description = "This is an example VLAN"
  reserved    = false

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}
