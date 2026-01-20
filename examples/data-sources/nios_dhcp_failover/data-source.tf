// Retrieve a specific DHCP failover using filters
data "nios_dhcp_failover" "dhcpfailover_by_name" {
  filters = {
    name = "dhcp_failover_1"
  }
}

// Retrieve specific DHCP failovers using Extensible Attributes
data "nios_dhcp_failover" "dhcpfailover_by_extattrs" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all the DHCP failovers
data "nios_dhcp_failover" "all_dhcpfailovers" {}
