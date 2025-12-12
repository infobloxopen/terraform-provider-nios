// Create a DHCP Roaming Host with Basic Fields
resource "nios_dhcp_roaminghost" "dhcp_roaminghost_basic" {
  name         = "roaming-host-1"
  match_client = "MAC_ADDRESS"
  mac          = "00:11:22:33:44:55"
}

// Create a DHCP Roaming Host with Address Type set to IPv6 and ipv6_match_option as V6_MAC_ADDRESS
resource "nios_dhcp_roaminghost" "dhcp_roaminghost_with_additional_fields" {
  name              = "roaming-host-2"
  address_type      = "IPV6"
  ipv6_match_option = "V6_MAC_ADDRESS"
  ipv6_mac_address  = "00:a1:b2:c3:44:55"
}

// Create a DHCP Roaming Host with Address Type IPv4 and match_client as CLIENT_ID
resource "nios_dhcp_roaminghost" "dhcp_roaminghost_ipv4_client_id" {
  name                   = "roaming-host-3"
  address_type           = "IPV4"
  match_client           = "CLIENT_ID"
  dhcp_client_identifier = "client-identifier-123"
}

// Create a DHCP Roaming Host with Address Type IPv6 and ipv6_match_option as DUID
resource "nios_dhcp_roaminghost" "dhcp_roaminghost_ipv6_duid" {
  name              = "roaming-host-4"
  address_type      = "IPV6"
  ipv6_match_option = "DUID"
  ipv6_duid         = "00:03:00:01:26:9e:5b:4c:00:11:22:33:44:55"
}

// Create a DHCP Roaming Host with Address Type set to BOTH and various additional fields
resource "nios_dhcp_roaminghost" "dhcp_roaminghost_address_type_both" {
  name              = "roaming-host-5"
  address_type      = "BOTH"
  match_client      = "MAC_ADDRESS"
  mac               = "66:77:88:99:aa:bb"
  ipv6_match_option = "V6_MAC_ADDRESS"
  ipv6_mac_address  = "00:b1:c2:d3:e4:f5"

  // Additional fields
  comment = "Roaminghost managed by Terraform"
  extattrs = {
    Site = "location-1"
  }

  options = [
    {
      num : 51
      value : "7200"
    },
    {
      name : "subnet-mask"
      value : "1.1.1.1"
    },
  ]
  use_options = true

  bootserver     = "example.com"
  use_bootserver = true
}
