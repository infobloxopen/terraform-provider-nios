//Create a Network Range with basic fields 
resource "nios_dhcp_range" "create_range_with_basic_fields" {
  start_addr = "10.0.0.170"
  end_addr   = "10.0.0.180"
}

//Create a Network Range with additional fields 
resource "nios_dhcp_range" "create_range_with_additional_fields" {
  start_addr        = "10.0.0.190"
  end_addr          = "10.0.0.200"
  comment           = "DHCP Range with additional fields"
  network_view      = "default"
  name              = "range object"
  always_update_dns = true

  extattrs = {
    "Site" = "location-1"
  }

  //filter rules 
  fingerprint_filter_rules = [{
    filter     = "finger_print_filter"
    permission = "Allow"
  }]

  nextserver     = "next_server.com"
  use_nextserver = true

  //failover association 
  server_association_type = "FAILOVER"
  failover_association    = "failover_association"

  ignore_id     = "MACADDR"
  use_ignore_id = true

}
