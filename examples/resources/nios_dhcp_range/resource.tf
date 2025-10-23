// Create Ipv4 network (required as parent )
resource "nios_ipam_network" "range_parent_network" {
  network      = "10.0.0.0/24"
  network_view = "default"
  comment      = "Parent network for DHCP ranges"
}

//Create a Network Range with basic fields 
resource "nios_dhcp_range" "create_range_with_basic_fields" {
  start_addr = "10.0.0.170"
  end_addr   = "10.0.0.180"
  depends_on = [nios_ipam_network.range_parent_network]
}

//Create a Network Range with additional fields 
resource "nios_dhcp_range" "create_range_with_additional_fields" {
  start_addr        = "10.0.0.190"
  end_addr          = "10.0.0.200"
  comment           = "DHCP Range with additional fields"
  network_view      = "default"
  name              = "range object"
  always_update_dns = true

  options = [
    {
      name  = "time-offset"
      num   = 2
      value = "50"
    },
    {
      name  = "dhcp-lease-time"
      num   = 51
      value = "7200"
    }
  ]
  use_options = true

  extattrs = {
    "Site" = "location-1"
  }

  nextserver     = "next_server.com"
  use_nextserver = true

  ignore_id     = "MACADDR"
  use_ignore_id = true

  depends_on = [nios_ipam_network.range_parent_network]
}
