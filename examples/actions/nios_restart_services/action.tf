resource "nios_ipam_network" "example" { // ( Required As Parent )
  network      = "10.0.0.0/24"
  network_view = "default"
}

action "nios_restart_services" "dhcpv4" {
  config {
    services       = ["DHCPV4"]
    restart_option = "RESTART_IF_NEEDED"
    members        = ["infoblox.member1"]
    user_name      = "admin"
    mode           = "SEQUENTIAL"
  }
}

# Restart the DHCPv4 service on the target member after the DHCP range is updated.
resource "nios_dhcp_range" "example" {
  start_addr = "10.0.0.170"
  end_addr   = "10.0.0.180"
  comment    = "Primary DHCP range"

  depends_on = [nios_ipam_network.example]

  lifecycle {
    action_trigger {
      events  = [after_update]
      actions = [action.nios_restart_services.dhcpv4]
    }
  }
}
