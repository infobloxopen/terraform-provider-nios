// Manage a Network with a Microsoft Server (required as parent)
resource "nios_ipam_network" "parent_network" {
  network      = "117.0.0.0/24"
  network_view = "ms_server"
  members = [
    {
      struct   = "msdhcpserver"
      ipv4addr = "10.34.98.68"
    }
  ]
}

// Manage a Range in the Network with a Microsoft Server (required as parent)
resource "nios_dhcp_range" "parent_range" {
  start_addr              = "117.0.0.190"
  end_addr                = "117.0.0.195"
  server_association_type = "MS_SERVER"
  ms_server               = { ipv4addr = "10.34.98.68" }
  network_view            = "ms_server"
  depends_on              = [nios_ipam_network.parent_network]
}

// Manage Microsoft Super Scope
resource "nios_microsoft_mssuperscope" "microsoft_mssuperscope_with_additional_fields" {
  // Basic Fields
  name   = "example_mssuperscope"
  ranges = [nios_dhcp_range.parent_range.ref]

  // Additional Attributes
  disable      = false
  comment      = "Super Scope Created By terraform"
  network_view = "ms_server"

  //Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}