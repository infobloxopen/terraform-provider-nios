// Create an IPv6 DHCP Option Definition in the default IPv6 Option Space (DHCPv6)
resource "nios_dhcp_ipv6optiondefinition" "ipv6_dhcp_option_definition_default_space" {
  name = "dhcp6.example_ipv6_option_definition_default_space"
  code = 4321
  type = "string"
}

// Create an IPv6 DHCP Option Space (Required as Parent)
resource "nios_dhcp_ipv6optionspace" "ipv6_dhcp_option_space" {
  name              = "example_ipv6_dhcp_option_space"
  enterprise_number = 5473
}

// Create an IPv6 DHCP Option Definition with type string in the created Option Space
resource "nios_dhcp_ipv6optiondefinition" "ipv6_dhcp_option_definition" {
  name  = "example_ipv6_dhcp_option_definition"
  code  = 1234
  space = nios_dhcp_ipv6optionspace.ipv6_dhcp_option_space.name
  type  = "string"
}

// Create an IPv6 DHCP Option Definition with type IP Address in the created Option Space
resource "nios_dhcp_ipv6optiondefinition" "ipv6_dhcp_option_definition_2" {
  name  = "example_ipv6_dhcp_option_definition_ipv6addr"
  code  = 1235
  space = nios_dhcp_ipv6optionspace.ipv6_dhcp_option_space.name
  type  = "ip-address"
}
