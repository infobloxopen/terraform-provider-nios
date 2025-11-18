// Manage IPAM Vlan Views (Required as Parent)
resource "nios_ipam_vlanview" "ipam_vlanview_parent" {
  start_vlan_id = 5
  end_vlan_id   = 10
  name          = "example_vlan_view"
}

// Create network for function call (required as parent)
resource "nios_ipam_network" "example_network" {
  network      = "85.85.0.0/16"
  network_view = "default"
  comment      = "Network for A record IP allocation"
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

// Manage IPAM Vlan using function call to retrieve next available Vlan id
resource "nios_ipam_vlan" "ipam_vlan_with_func_call" {
  name   = "example_vlan_function_call"
  parent = nios_ipam_vlanview.ipam_vlanview_parent.ref
  func_call = {
    attribute_name  = "id"
    object_function = "next_available_vlan_id"
    result_field    = "id"
    object          = "network"
    object_parameters = {
      network      = "85.85.0.0/16"
      network_view = "default"
    }
  }
}
