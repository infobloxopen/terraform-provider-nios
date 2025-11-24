// Create an IPv6 Network (Required as Parent)
resource "nios_ipam_ipv6network" "parent_ipv6network" {
  network = "15::/64"
}

// Create DHCP IPv6 Range with Basic Fields
resource "nios_dhcp_ipv6range" "create_ipv6range_with_basic_fields" {
  network    = nios_ipam_ipv6network.parent_ipv6network.network
  start_addr = "15::10"
  end_addr   = "15::20"
}

// Create DHCP IPv6 Range with Additional Fields
resource "nios_dhcp_ipv6range" "create_ipv6range_with_additional_fields" {
  network    = nios_ipam_ipv6network.parent_ipv6network.network
  start_addr = "15::30"
  end_addr   = "15::50"

  // Additional Fields
  comment = "DHCP IPv6 Range created by Terraform"
  extattrs = {
    Site = "location-1"
  }
  exclude = [{
    start_address = "15::40",
    end_address   = "15::45",
    comment       = "Exclude range 15::40 - 15::45",
  }]
}
