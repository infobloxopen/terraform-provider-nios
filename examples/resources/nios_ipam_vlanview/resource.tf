// Manage IPAM Vlan Views with Basic Fields
resource "nios_ipam_vlanview" "ipam_vlanview_basic" {
  start_vlan_id = 5
  end_vlan_id   = 10
  name          = "example_vlan_view"
}

// Manage IPAM Vlan Views with Additional Fields
resource "nios_ipam_vlanview" "ipam_vlanview_with_additional_fields" {
  start_vlan_id = 50
  end_vlan_id   = 100
  name          = "example_vlan_view2"

  // Additional Fields
  comment                 = "Example VLAN View"
  allow_range_overlapping = true

  //Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}
