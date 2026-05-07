# Scenario: Full ignore configuration
name                                = "full-ignore-config"
ignore_mac_addresses                = ["00:00:00:00:00:01", "00:00:00:00:00:02"]
ignore_client_identifier            = true
use_ignore_client_identifier = false
//use_ignore_client_identifier        = true
//ignore_id                           = "NONE"
use_ignore_id                       = false
use_bootfile                        = true
ignore_dhcp_option_list_request     = true
use_ignore_dhcp_option_list_request = true
comment                             = "Full ignore configuration"
nextserver = "10.0.0.1"
use_nextserver = true
