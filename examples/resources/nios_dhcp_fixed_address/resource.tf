// Create Fixed Address with Basic Fields
resource "nios_dhcp_fixed_address" "create_address_basic" {
  ipv4addr     = "16.0.0.10"
  match_client = "MAC_ADDRESS"
  mac          = "00:1a:2b:3c:4d:5e"

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}

// Create Fixed Address with Additional Fields
resource "nios_dhcp_fixed_address" "create_address_additional" {
  // Basic Fields
  ipv4addr     = "16.0.0.20"
  match_client = "MAC_ADDRESS"
  mac          = "00:6a:7b:8c:9d:5e"

  // Additional Fields
  agent_remote_id = 30
  comment         = "Fixed Address created with additional fields"

  bootfile     = "pxelinux.0"
  use_bootfile = true

  enable_ddns     = true
  use_enable_ddns = true

  pxe_lease_time     = 3600
  use_pxe_lease_time = true

  device_location = "APJ"
  device_type     = "Server"

  options = [
    {
      name  = "time-offset"
      value = "50"
    },
    {
      name  = "dhcp-lease-time"
      value = "7200"
    }
  ]

  // Extensible Attributes
  extattrs = {
    Site = "location-2"
  }
}

// Create Fixed Address using function call to retrieve ipv4addr
resource "nios_dhcp_fixed_address" "create__address_with_func_call" {
  match_client     = "CIRCUIT_ID"
  agent_circuit_id = 250
  func_call = {
    attribute_name  = "ipv4addr"
    object_function = "next_available_ip"
    result_field    = "ips"
    object          = "network"
    object_parameters = {
      network      = "16.0.0.0/24"
      network_view = "default"
    }
  }
  comment = "Fixed Address created with ipv4addr retrieved via function call"
}
