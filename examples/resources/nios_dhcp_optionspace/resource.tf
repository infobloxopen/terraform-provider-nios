// Create a DHCP Option Space with Basic Fields
resource "nios_dhcp_dhcpoptionspace" "dhcp_option_space_with_basic_fields" {
  name = "example_option_space_1"
}

// Create a DHCP Option Definition
resource "nios_dhcp_dhcpoptiondefinition" "dhcp_option_definition" {
  code  = 30
  name  = "example_option_definition_1"
  type  = "string"
  space = nios_dhcp_dhcpoptionspace.dhcp_option_space_with_basic_fields.name
}

// Create a DHCP Option Space with Additional Fields
resource "nios_dhcp_dhcpoptionspace" "dhcp_option_space_with_additional_fields" {
  name    = "example_option_space_2"
  comment = "DHCP Option Space created by Terraform"
  option_definitions = [
    nios_dhcp_dhcpoptiondefinition.dhcp_option_definition.name
  ]
}
