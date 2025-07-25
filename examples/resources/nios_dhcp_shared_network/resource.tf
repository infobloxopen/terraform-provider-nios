// Create SharedNetwork Record with Basic Fields
resource "nios_dhcp_sharednetwork" "shared_network_basic_fields" {
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

// Create SharedNetwork Record using function call to retrieve ipv4addr
resource "nios_dhcp_sharednetwork" "shared_network_additional_fields" {
  name = "example_shared_network2"
  networks = [
    {
      ref = "network/ZG5zLm5ldHdvcmskMTUuMTQuMS4wLzI0LzA:15.14.1.0/24/default"
    },
    {
      ref = "network/ZG5zLm5ldHdvcmskMTYuMC4wLjAvMjQvMA:16.0.0.0/24/default"
    }
  ]
  use_options = true
  options = [
    {
      name  = "domain-name-servers"
      value = "11.22.1.2"
    },
    {
      name  = "time-offset"
      value = "1000"
    },
    {
      name  = "domain-name"
      value = "aa.bb.com"
    },
  ]
  use_logic_filter_rules = true
  logic_filter_rules = [
    {
      filter = "option_filter"
      type   = "Option"
    }
  ]
  comment                    = "shared network with additional fields"
  ddns_server_always_updates = true
  ddns_use_option81          = true
  use_ddns_use_option81      = true
}
