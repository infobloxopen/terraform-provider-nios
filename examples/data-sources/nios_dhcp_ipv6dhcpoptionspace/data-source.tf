// Retrieve a specific IPv6 DHCP Option Space by filters
data "nios_dhcp_ipv6dhcpoptionspace" "get_ipv6_dhcp_option_space_using_filters" {
  filters = {
    name = "example_ipv6_dhcp_option_space"
  }
}

// Retrieve all IPv6 DHCP Option Spaces
data "nios_dhcp_ipv6dhcpoptionspace" "get_all_ipv6_dhcp_option_spaces" {}
