// Create Filter Relay Agent with Basic Fields
resource "nios_dhcp_filterrelayagent" "filter_relay_agent_basic_fields" {
  name          = "filter_relay_agent_example"
  is_circuit_id = "NOT_SET"
}

// Create Filter Relay Agent with Additional Fields
resource "nios_dhcp_filterrelayagent" "filter_relay_agent_additional_fields" {
  name                        = "filter_relay_agent_example_2"
  is_circuit_id               = "MATCHES_VALUE"
  circuit_id_name             = "Circuit-123"
  circuit_id_substring_length = 11
  circuit_id_substring_offset = 5
  is_circuit_id_substring     = true
  remote_id_name              = "Remote-456"
  remote_id_substring_length  = 10
  remote_id_substring_offset  = 3
  is_remote_id_substring      = true
  is_remote_id                = "MATCHES_VALUE"
  extattrs = {
    Site = "location-1"
  }
}
