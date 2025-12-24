// Create a DHCP IPv6 Filter Option with Basic Fields
resource "nios_dhcp_ipv6filteroption" "dhcp_ipv6filteroption_with_basic_fields" {
  name = "example_ipv6_filter_option_1"
}

// Create a DHCP IPv6 Filter Option with Additional Fields
resource "nios_dhcp_ipv6filteroption" "dhcp_ipv6filteroption_with_additional_fields" {
  name = "example_ipv6_filter_option_2"

  // Additional Fields
  comment    = "IPv6  Filter Option created via Terraform"
  expression = "(option dhcp6.server-id=\"server-id\" AND option dhcp6.vendor-class=\"DHCPv6\")"
  lease_time = 7200
  option_list = [
    {
      name         = "dhcp6.name-servers"
      num          = 23
      value        = "fc00::,2001:db8::"
      vendor_class = "DHCPv6"
    },
    {
      name         = "dhcp6.remote-id"
      num          = 37
      value        = "remote-id"
      vendor_class = "DHCPv6"
    }
  ]

  //Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}
