// Retrieve a specific IPv6 DHCP Option Definition by filters
data "nios_dhcp_ipv6dhcpoptiondefinition" "get_ipv6_dhcp_option_definition_using_filters" {
  filters = {
    name = "example_ipv6_dhcp_option_definition"
  }
}

// Retrieve all IPv6 DHCP Option Definitions
data "nios_dhcp_ipv6dhcpoptiondefinition" "get_all_ipv6_dhcp_option_definitions" {}
