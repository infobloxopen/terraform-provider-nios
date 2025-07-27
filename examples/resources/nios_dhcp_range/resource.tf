//Create a Network Range with basic fields 
resource "nios_dhcp_range" "range_basic_fields" {
  start_addr = "10.0.0.170"
  end_addr   = "10.0.0.180"
}

//Create a Network Range with additional parameters 
resource "nios_dhcp_range" "range_additional_fields" {
  start_addr        = "10.0.0.190"
  end_addr          = "10.0.0.200"
  comment           = "Full range for testing"
  network_view      = "default"
  name              = "range object"
  always_update_dns = true


  // Email and notification settings
  email_list     = ["admin@example.com", "network@example.com"]
  use_email_list = true

  // Water mark settings
  high_water_mark       = 95
  high_water_mark_reset = 85
  low_water_mark        = 10
  low_water_mark_reset  = 20

  deny_all_clients = true
  deny_bootp       = true
  use_deny_bootp   = true

  //DHCP Threshold
  enable_dhcp_thresholds     = true
  use_enable_dhcp_thresholds = true

  extattrs = {
    "Site" = "location-1"
  }

  //filter rules 
  fingerprint_filter_rules = [{
    filter     = "finger_print_filter"
    permission = "Allow"
  }]
  mac_filter_rules = [{
    filter     = "mac_filter"
    permission = "Deny"
    }
  ]
  nac_filter_rules = [{
    filter     = "nac_filter"
    permission = "Deny"
    }
  ]
  option_filter_rules = [
    {
      filter     = "option_filter"
      permission = "Allow"
    }
  ]
  relay_agent_filter_rules = [
    {
      filter     = "relay_agent_filter"
      permission = "Allow"
    }
  ]
  logic_filter_rules = [
    {
      filter = "mac_filter"
      type   = "MAC"
    }
  ]
  use_logic_filter_rules = true

  ignore_dhcp_option_list_request     = true
  use_ignore_dhcp_option_list_request = true

  nextserver     = "next_server.com"
  use_nextserver = true

  //failover association 
  server_association_type = "FAILOVER"
  failover_association    = "failover_association"

  ignore_id     = "MACADDR"
  use_ignore_id = true

  known_clients = "Allow"
}
