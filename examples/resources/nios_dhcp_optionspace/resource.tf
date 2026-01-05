// Create a DHCP Option Space with Basic Fields
resource "nios_dhcp_optionspace" "dhcp_option_space_with_basic_fields" {
  name = "example_option_space_1"
}

//  NOTE: option_definitions is a computed (read-only) field and cannot be set here.
//  It will be automatically populated when option definitions are created that
//  reference this space (nios_dhcp_ioptiondefinition.dhcp_option_space_with_basic_fields) as shown below.

// Create a DHCP Option Definition
resource "nios_dhcp_optiondefinition" "dhcp_option_definition" {
  code  = 30
  name  = "example_option_definition_1"
  type  = "string"
  space = nios_dhcp_optionspace.dhcp_option_space_with_basic_fields.name
}

