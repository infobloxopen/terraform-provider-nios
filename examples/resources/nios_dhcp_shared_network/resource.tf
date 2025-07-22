// Create SharedNetwork Record with Basic Fields
resource "nios_dhcp_sharednetwork" "shared_network_basic_fields" {
  name = "example_shared_network1"
  networks = ["network/ZG5zLm5ldHdvcmskMTEuMTEuMTIuMC8yNC8w:11.11.12.0/24/default",
  "network/ZG5zLm5ldHdvcmskMTIuMTIuMTEuMC8yNC8w:12.12.11.0/24/default"]
  network_view = "default"
  extattrs = {
    Site = "location-1"
  }
}

// Create SharedNetwork Record using function call to retrieve ipv4addr
resource "nios_dhcp_sharednetwork" "shared_network_additional_fields" {
  name = "example_shared_network2"
  networks = ["network/ZG5zLm5ldHdvcmskMTQuMTQuMS4wLzI0LzA:14.14.1.0/24/default",
  "network/ZG5zLm5ldHdvcmskMTUuMTQuMS4wLzI0LzA:15.14.1.0/24/default"]
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
  logic_filter_rules = [
    {
      filter = "option_filter"
      type   = "Option"
    }
  ]
  comment = "shared network with additional fields"
}
