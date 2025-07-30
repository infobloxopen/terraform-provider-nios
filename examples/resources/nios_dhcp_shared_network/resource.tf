// Create Shared Network with Basic Fields
resource "nios_dhcp_shared_network" "shared_network_basic_fields" {
  name = "example_shared_network1"
  networks = [
    {
      ref = "network/ZG5zLm5ldHdvcmskMjEuMjEuMTQuMC8yNC8w:21.21.14.0/24/default"
    },
    {
      ref = "network/ZG5zLm5ldHdvcmskMjEuMjEuMTMuMC8yNC8w:21.21.13.0/24/default"
    }
  ]
  network_view = "default"
  extattrs = {
    Site = "location-1"
  }
}

// Create Shared Network with Additional Fields
resource "nios_dhcp_shared_network" "shared_network_additional_fields" {
  name = "example_shared_network2"
  networks = [
    {
      ref = "network/ZG5zLm5ldHdvcmskMTUuMTQuMS4wLzI0LzA:15.14.1.0/24/default"
    },
    {
      ref = "network/ZG5zLm5ldHdvcmskMTYuMC4wLjAvMjQvMA:16.0.0.0/24/default"
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
  use_logic_filter_rules = true
  # logic_filter_rules = [
  #   {
  #     filter = "option_filter"
  #     type   = "Option"
  #   }
  # ]
  comment                    = "Shared network with additional fields"
  ddns_server_always_updates = true
  ddns_use_option81          = true
  use_ddns_use_option81      = true
}
