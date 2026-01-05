// Manage an IPV4 Network (Required as Parent)
resource "nios_ipam_network" "parent_network" {
  network      = "21.0.0.0/24"
  network_view = "default"
  comment      = "Parent network for DHCP fixed addresses"
}

// Manage a Fixed Address (Required as Parent)
resource "nios_dhcp_fixed_address" "parent_fixed_address" {
  ipv4addr     = "21.0.0.10"
  match_client = "MAC_ADDRESS"
  mac          = "00:1f:2b:3c:4d:5e"

  depends_on = [nios_ipam_network.parent_network]
}

// Manage an Auth Zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_auth_zone" {
  fqdn = "example_auth_as_parent.com"
  view = "default"
}

// Manage a Record A (Required as Parent)
resource "nios_dns_record_a" "parent_record_a" {
  name     = "parent-record.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  ipv4addr = "10.20.1.2"

  depends_on = [nios_dns_zone_auth.parent_auth_zone]
}

// Manage an IPAM Super Host with Basic Fields
resource "nios_ipam_superhost" "ipam_super_host_basic" {
  name = "example_super_host"
}

// Manage an IPAM Super Host with Additional Fields
resource "nios_ipam_superhost" "ipam_super_host_with_additional_fields" {
  name = "example_super_host2"

  // Additional Fields
  comment                   = "This is a Super Host example"
  delete_associated_objects = true
  disabled                  = false
  dns_associated_objects    = [nios_dns_record_a.parent_record_a.ref]
  dhcp_associated_objects   = [nios_dhcp_fixed_address.parent_fixed_address.ref]

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}
