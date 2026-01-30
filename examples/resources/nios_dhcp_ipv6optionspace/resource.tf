// Create an Ipv6 DHCP Option Space with Basic Fields
resource "nios_dhcp_ipv6optionspace" "ipv6_dhcp_option_space_with_basic_fields" {
  name              = "example_ipv6_dhcp_option_space"
  enterprise_number = 5473
}

//  NOTE: option_definitions is a computed (read-only) field and cannot be set here.
//  It will be automatically populated when option definitions are created that
//  reference this space (nios_dhcp_ipv6optiondefinition.ipv6_dhcp_option_space_with_basic_fields) as shown below.

// Create an Ipv6 DHCP Option Definition in the above created Option Space
resource "nios_dhcp_ipv6optiondefinition" "ipv6_dhcp_option_definition" {
  name  = "example_option_definition"
  code  = 1234
  space = nios_dhcp_ipv6optionspace.ipv6_dhcp_option_space_with_basic_fields.name
  type  = "string"
}
