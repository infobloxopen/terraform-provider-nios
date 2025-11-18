// Create an IPv6 DHCP Option Space (Required as Parent)
resource "nios_dhcp_ipv6dhcpoptionspace" "ipv6_dhcp_option_space" {
  name              = "example_ipv6_dhcp_option_space"
  enterprise_number = 5473
}

// Create an IPv6 DHCP Option Definition with type string
resource "nios_dhcp_ipv6dhcpoptiondefinition" "ipv6_dhcp_option_definition" {
  name  = "example_ipv6_dhcp_option_definition"
  code  = 1234
  space = nios_dhcp_ipv6dhcpoptionspace.ipv6_dhcp_option_space.name
  type  = "string"
}

// Create an IPv6 DHCP Option Definition with type IP Address
resource "nios_dhcp_ipv6dhcpoptiondefinition" "ipv6_dhcp_option_definition_2" {
  name  = "example_ipv6_dhcp_option_definition_ipv6addr"
  code  = 1235
  space = nios_dhcp_ipv6dhcpoptionspace.ipv6_dhcp_option_space.name
  type  = "ip-address"
}
