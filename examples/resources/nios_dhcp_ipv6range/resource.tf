// Create an IPv6 Network (Required as Parent)
resource "nios_ipam_ipv6network" "parent_ipv6network" {
  network = "15::/64"
}

// Create DHCP IPv6 Range with Address Type set to "ADDRESS" (default)
resource "nios_dhcp_ipv6range" "create_ipv6range_with__default_address_type" {
  network      = nios_ipam_ipv6network.parent_ipv6network.network
  start_addr   = "15::10"
  end_addr     = "15::20"
  address_type = "ADDRESS"
}

// Create DHCP IPv6 Range with Address Type set to "PREFIX"
resource "nios_dhcp_ipv6range" "create_ipv6range_with_address_type_prefix" {
  network           = nios_ipam_ipv6network.parent_ipv6network.network
  ipv6_start_prefix = "15:0:0:1000::"
  ipv6_end_prefix   = "15:0:0:1fff::"
  ipv6_prefix_bits  = 80
  address_type      = "PREFIX"
}

// Create DHCP IPv6 Range with Address Type set to "BOTH"
resource "nios_dhcp_ipv6range" "create_ipv6range_with_address_type_both" {
  network           = nios_ipam_ipv6network.parent_ipv6network.network
  start_addr        = "15::25"
  end_addr          = "15::35"
  ipv6_start_prefix = "15:0:0:2000::"
  ipv6_end_prefix   = "15:0:0:2fff::"
  ipv6_prefix_bits  = 80
  address_type      = "BOTH"
}

// Create DHCP IPv6 Range with Additional Fields
resource "nios_dhcp_ipv6range" "create_ipv6range_with_additional_fields" {
  network    = nios_ipam_ipv6network.parent_ipv6network.network
  start_addr = "15::40"
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
