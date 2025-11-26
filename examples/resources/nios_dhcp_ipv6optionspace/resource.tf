// Create an Ipv6 DHCP Option Space with Basic Fields
resource "nios_dhcp_ipv6optionspace" "ipv6_dhcp_option_space_with_basic_fields" {
  name              = "example_ipv6_dhcp_option_space"
  enterprise_number = 5473
}

// Create an Ipv6 DHCP Option Definition
resource "nios_dhcp_ipv6optiondefinition" "ipv6_dhcp_option_definition" {
  name  = "example_option_definition"
  code  = 1234
  space = nios_dhcp_ipv6optionspace.ipv6_dhcp_option_space_with_basic_fields.name
  type  = "string"
}

// Create an Ipv6 DHCP Option Space with Additional Fields using pre-created Option Definition
resource "nios_dhcp_ipv6optionspace" "ipv6_dhcp_option_space_with_additional_fields" {
  name               = "example_ipv6_dhcp_option_space_2"
  enterprise_number  = 5678
  option_definitions = [nios_dhcp_ipv6optiondefinition.ipv6_dhcp_option_definition.name]
  comment            = "IPv6 DHCP Option Space created by Terraform"
}
