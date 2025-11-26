// Retrieve a specific IPV6 Fixed Addresses by filters
data "nios_dhcp_ipv6fixedaddress" "get_ipv6_fixed_address_using_filters" {
  filters = {
    ipv4addr = "16.0.0.20"
  }
}

// Retrieve specific IPV6 Fixed Addresses using Extensible Attributes
data "nios_dhcp_ipv6fixedaddress" "get_ipv6_fixed_address_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Search for a IPV6 Fixed aAddress by Microsoft Server
data "nios_dhcp_ipv6fixedaddress" "get_ipv6_fixed_address_using_microsoft_server" {
  body = {
    ms_server = {
      struct   = "msdhcpserver"
      ipv4addr = "1.1.1.1" // Specify the IP address of the Microsoft DHCP server
    }
  }
}

// Retrieve all IPV6 Fixed Addresses
data "nios_dhcp_ipv6fixedaddress" "get_all_ipv6_fixed_address" {}
