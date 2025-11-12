// Create a DHCP Option Space with Basic Fields
resource "nios_dhcp_dhcpoptionspace" "dhcp_option_space_with_basic_fields" {
  name = "example_option_space_1"
}

// Create a DHCP Option Space with Additional Fields
resource "nios_dhcp_dhcpoptionspace" "dhcp_option_space_with_additional_fields" {
  name    = "example_option_space_2"
  comment = "DHCP Option Space created by Terraform"
}
