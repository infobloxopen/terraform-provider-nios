// Retrieve a DHCP failover by filter using name
data "nios_dhcp_dhcpfailover" "dhcpfailover_by_name" {
  filters = {
    name = "dhcp_failover_1"
  }
}

// Retrieve a DHCP failover by filter using ext-attrs
data "nios_dhcp_dhcpfailover" "dhcpfailover_by_extattrs" {
  extattrfilters = {
    Site = "NewYork"
  }
}

// Retrieve all the DHCP failovers
data "nios_dhcp_dhcpfailover" "all_dhcpfailovers" {
  // No filters applied, retrieves all DHCP failovers
}
