// Create a DHCP Option Space (Required as Parent)
resource "nios_dhcp_dhcpoptionspace" "option_space_with_basic_fields" {
  name = "example_option_space"
}

// Create a DHCP Option Definition
resource "nios_dhcp_dhcpoptiondefinition" "option_definition_with_basic_fields" {
  code  = 10
  name  = "example_option_definition"
  type  = "string"
  space = nios_dhcp_dhcpoptionspace.option_space_with_basic_fields.name
}
