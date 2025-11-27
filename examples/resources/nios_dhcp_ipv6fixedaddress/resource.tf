// Create an IPV6 Network (Required as Parent)
resource "nios_ipam_ipv6network" "parent_network" {
  network      = "2001:db8:abcd:1234::/64"
  network_view = "default"
  comment      = "Parent network for DHCP fixed addresses"
}

// Create an IPV6 Fixed Address with Basic Fields
resource "nios_dhcp_ipv6fixedaddress" "create_ipv6_fixed_address_basic" {
  ipv6addr = "2001:db8:abcd:1234::1"
  duid     = "01:01:00:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
  network  = nios_ipam_ipv6network.parent_network.network
}

// Create an IPV6 Fixed Address with Additional Fields with PREFIX address type
resource "nios_dhcp_ipv6fixedaddress" "create_ipv6_fixed_address_additional1" {
  // Basic Fields
  address_type    = "PREFIX"
  ipv6prefix      = "2001:db8:abcd:1234::"
  ipv6prefix_bits = 64
  match_client    = "MAC_ADDRESS"
  mac_address     = "01:6a:7b:8c:9d:5e"
  network_view    = "default"

  // Additional Fields
  comment = "IPV6 Fixed Address created with additional fields"

  options = [
    {
      name  = "domain-name"
      num   = 15
      value = "example.com"
    },
    {
      name  = "dhcp-renewal-time"
      num   = 58
      value = "720"
    }
  ]
  use_options = true
  // Extensible Attributes
  extattrs = {
    Site = "location-2"
  }
  depends_on = [nios_ipam_ipv6network.parent_network]
}

// Create an IPV6 Fixed Address with Additional Fields with BOTH address type
resource "nios_dhcp_ipv6fixedaddress" "create_ipv6_fixed_address_additional2" {
  // Basic Fields
  address_type    = "BOTH"
  ipv6addr        = "2001:db8:abcd:1234::3"
  ipv6prefix      = "2001:db8:abcd:1234::"
  ipv6prefix_bits = 64
  match_client    = "MAC_ADDRESS"
  mac_address     = "00:6a:7b:8c:9d:5e"
  network_view    = "default"

  // Additional Fields
  preferred_lifetime      = 2400
  use_preferred_lifetime  = true
  valid_lifetime          = 4800
  use_valid_lifetime      = true
  domain_name             = "example.com"
  use_domain_name         = true
  domain_name_servers     = ["2001:4860:4860::8888", "2001:4860:4860::8844"]
  use_domain_name_servers = true
  use_logic_filter_rules  = true
  logic_filter_rules = [
    {
      filter = "ipv6_option_filter"
      type   = "Option"
    }
  ]
  depends_on = [nios_ipam_ipv6network.parent_network]
}

// Create an IPV6 Fixed Address using function call to retrieve ipv4addr
resource "nios_dhcp_ipv6fixedaddress" "create_ipv6_fixed_address_with_func_call" {
  duid = "00:01:01:01:1d:2b:3c:4d:00:0c:29:ab:cd:ef"
  func_call = {
    attribute_name  = "ipv6addr"
    object_function = "next_available_ip"
    result_field    = "ips"
    object          = "ipv6network"
    object_parameters = {
      network      = "2001:db8:abcd:1234::/64"
      network_view = "default"
    }
  }
  comment    = "Fixed Address created with ipv4addr retrieved via function call"
  depends_on = [nios_ipam_ipv6network.parent_network]
}
