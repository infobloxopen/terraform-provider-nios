// create Dhcp Failover with Basic Fields
resource "nios_dhcp_dhcpfailover" "dhcpfailover_1" {
  name                  = "dhcp_failover_1"
  primary               = "infoblox.localdomain"
  secondary             = "infoblox.member1"
  primary_server_type   = "GRID"
  secondary_server_type = "GRID"
}

// create  Dhcp Failover with Additional Fields
resource "nios_dhcp_dhcpfailover" "dhcpfailover_2" {
  name                  = "dhcp_failover_2"
  primary               = "infoblox.member2"
  secondary             = "10.197.36.31"
  primary_server_type   = "GRID"
  secondary_server_type = "EXTERNAL"
  failover_port         = 62001
  use_failover_port     = true
  extattrs = {
    Site = "location-1"
  }
}
