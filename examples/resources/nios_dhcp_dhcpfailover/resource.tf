// create a dhcp failover with required fields
resource "nios_dhcp_dhcpfailover" "dhcpfailover_1" {
  name                  = "dhcp_failover_1"
  primary               = "infoblox.172_28_83_59"
  secondary             = "infoblox.172_28_82_171"
  primary_server_type   = "GRID"
  secondary_server_type = "GRID"
}

// create a dhcp failover with ext-attrs
resource "nios_dhcp_dhcpfailover" "dhcpfailover_2" {
  name                  = "dhcp_failover_2"
  primary               = "infoblox.member"
  secondary             = "10.197.36.31"
  primary_server_type   = "GRID"
  secondary_server_type = "EXTERNAL"
  extattrs = {
    Site = "NewYork"
  }
}
