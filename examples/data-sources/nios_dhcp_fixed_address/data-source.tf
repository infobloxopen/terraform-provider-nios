// Retrieve a specific Fixed Addresses by filters
data "nios_dhcp_fixed_address" "get_fixed_address_using_filters" {
  filters = {
    name = "example_fixed_address"
  }
}

// Retrieve specific Fixed Addresses using Extensible Attributes
data "nios_dhcp_fixed_address" "get_fixed_address_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Search for a fixed address by Microsoft Server
data "nios_dhcp_fixed_address" "get_fixed_address_using_microsoft_server" {
  body = {
    ms_server = {
      struct   = "msdhcpserver"
      ipv4addr = "1.1.1.1" // Specify the IP address of the Microsoft DHCP server
    }
  }
}

// Retrieve all Fixed Addresses
data "nios_dhcp_fixed_address" "get_all_fixed_address" {}
