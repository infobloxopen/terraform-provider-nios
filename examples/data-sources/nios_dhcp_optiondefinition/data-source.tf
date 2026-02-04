// Retrieve a specific DHCP Option Definition by filters
data "nios_dhcp_optiondefinition" "get_dhcp_option_definition_using_filters" {
  filters = {
    name  = "example_option_definition"
    space = "example_option_space"
  }
}

// Retrieve all DHCP Option Definitions
data "nios_dhcp_optiondefinition" "get_all_dhcp_option_definitions" {}
