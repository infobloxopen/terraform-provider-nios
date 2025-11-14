// Retrieve a specific DHCP Option Definition by filters
data "nios_dhcp_dhcpoptiondefinition" "get_dhcp_option_definition_using_filters" {
  filters = {
    name  = "example_option_definition"
    code  = 10
    space = "example_option_space"
    type  = "string"
  }
}

// Retrieve all IPv6 DHCP Option Definitions
data "nios_dhcp_dhcpoptiondefinition" "get_all_dhcp_option_definitions" {}
