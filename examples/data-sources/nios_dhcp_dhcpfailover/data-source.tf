// Retrieve a specific DHCP failover using filters
data "nios_dhcp_dhcpfailover" "dhcpfailover_by_name" {
  filters = {
    name = "dhcp_failover_1"
  }
}

// Retrieve specific DHCP failovers using Extensible Attributes
data "nios_dhcp_dhcpfailover" "dhcpfailover_by_extattrs" {
  extattrfilters = {
    Site = "NewYork"
  }
}

// Retrieve all the DHCP failovers
data "nios_dhcp_dhcpfailover" "all_dhcpfailovers" {}
