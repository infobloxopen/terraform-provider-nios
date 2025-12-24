// Retrieve a specific DHCP Filter Relay Agent by filters
data "nios_dhcp_filterrelayagent" "get_filterrelayagent_using_filters" {
  filters = {
    name = "filter_relay_agent_example"
  }
}
// Retrieve specific DHCP Filter Relay Agents using extensible attributes
data "nios_dhcp_filterrelayagent" "get_filterrelayagent_using_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DHCP Filter Relay Agents
data "nios_dhcp_filterrelayagent" "get_all_filterrelayagents" {
}
