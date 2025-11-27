// Create Two IPv6 Networks (Required as Parent)
resource "nios_ipam_ipv6network" "parent_ipv6network1" {
  network = "2024:db8:1::/64"
}

resource "nios_ipam_ipv6network" "parent_ipv6network2" {
  network = "2024:db8:2::/64"
}

// Manage DHCP Ipv6 Shared Network with Basic Fields
resource "nios_dhcp_ipv6sharednetwork" "dhcp_ipv6sharednetwork_basic_fields" {
  name = "shared_network_1"
  networks = [{
    ref = nios_ipam_ipv6network.parent_ipv6network1.ref
  }]
}

// Manage Ipv6 Shared Network with Additional Fields
resource "nios_dhcp_ipv6sharednetwork" "dhcp_ipv6sharednetwork_with_additional_fields" {
  name = "shared_network_2"
  networks = [{
    ref = nios_ipam_ipv6network.parent_ipv6network2.ref
  }]

  // Additional Fields
  comment         = "Ipv6 Shared Network created by Terraform"
  domain_name     = "example.com"
  use_domain_name = true
  options = [
    {
      name  = "dhcp-lease-time"
      num   = "51"
      value = "50000"
    }
  ]
  use_options        = true
  valid_lifetime     = 50000
  use_valid_lifetime = true
  extattrs = {
    "Site" = "location-1"
  }
}
