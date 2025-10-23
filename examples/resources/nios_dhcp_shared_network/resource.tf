// Create IPV4 Networks (required as parents)
resource "nios_ipam_network" "parent_network1" {
  network      = "21.21.14.0/24"
  network_view = "default"
  comment      = "Parent network for shared network 1"
}

resource "nios_ipam_network" "parent_network2" {
  network      = "21.21.13.0/24"
  network_view = "default"
  comment      = "Parent network for shared network 1"
}

resource "nios_ipam_network" "parent_network3" {
  network      = "15.14.1.0/24"
  network_view = "default"
  comment      = "Parent network for shared network 2"
}

resource "nios_ipam_network" "parent_network4" {
  network      = "16.0.0.0/24"
  network_view = "default"
  comment      = "Parent network for shared network 2"
}

// Create Shared Network with Basic Fields
resource "nios_dhcp_shared_network" "shared_network_basic_fields" {
  name = "example_shared_network1"
  networks = [
    {
      ref = nios_ipam_network.parent_network1.ref
    },
    {
      ref = nios_ipam_network.parent_network2.ref
    }
  ]
  network_view = "default"
  extattrs = {
    Site = "location-1"
  }
  depends_on = [
    nios_ipam_network.parent_network1,
    nios_ipam_network.parent_network2
  ]
}

// Create Shared Network with Additional Fields
resource "nios_dhcp_shared_network" "shared_network_additional_fields" {
  name = "example_shared_network2"
  networks = [
    {
      ref = nios_ipam_network.parent_network3.ref
    },
    {
      ref = nios_ipam_network.parent_network4.ref
    }
  ]
  ignore_mac_addresses = ["66:77:88:99:aa:bb", "00:11:22:33:44:55"]
  use_options          = true
  options = [
    {
      name  = "domain-name-servers"
      num   = "6"
      value = "11.22.1.2"
    },
    {
      name  = "time-offset"
      num   = "2"
      value = "1000"
    },
    {
      name  = "domain-name"
      num   = "15"
      value = "aa.bb.com"
    },
  ]
  comment                    = "Shared network with additional fields"
  ddns_server_always_updates = true
  ddns_use_option81          = true
  use_ddns_use_option81      = true
  depends_on = [
    nios_ipam_network.parent_network3,
    nios_ipam_network.parent_network4
  ]
}
