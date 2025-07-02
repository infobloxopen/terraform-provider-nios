resource "nios_ipam_networkcontainer" "example_container" {
  network      = "10.0.0.0/24"
  network_view = "default"
  comment      = "Created by Terraform"

  // Optional: Configure extensible attributes
  extattrs = {
    "Site" = "location-1"
  }
}

resource "nios_ipam_networkcontainer" "complete_example" {
  // Required attributes
  network = "11.0.0.0/24"

  // Basic configuration
  network_view = "default"
  comment      = "Complete network container example with all possible writable attributes"

  // Authority and boot settings
  authority = true
  # auto_create_reversezone = false
  bootfile       = "pxelinux.0"
  bootserver     = "192.168.1.10"
  use_authority  = true
  use_bootfile   = true
  use_bootserver = true

  // DDNS settings
  ddns_domainname                 = "example.com"
  ddns_generate_hostname          = true
  ddns_ttl                        = 3600
  ddns_update_fixed_addresses     = true
  ddns_use_option81               = true
  use_ddns_domainname             = true
  use_ddns_generate_hostname      = true
  use_ddns_ttl                    = true
  use_ddns_update_fixed_addresses = true
  use_ddns_use_option81           = true

  // DHCP settings
  deny_bootp                          = false
  enable_ddns                         = true
  enable_dhcp_thresholds              = true
  enable_email_warnings               = true
  enable_pxe_lease_time               = true
  enable_snmp_warnings                = true
  ignore_dhcp_option_list_request     = false
  lease_scavenge_time                 = 86400
  pxe_lease_time                      = 43200
  recycle_leases                      = true
  update_dns_on_lease_renewal         = true
  use_deny_bootp                      = true
  use_enable_ddns                     = true
  use_enable_dhcp_thresholds          = true
  use_ignore_dhcp_option_list_request = true
  use_lease_scavenge_time             = true
  use_pxe_lease_time                  = true
  use_recycle_leases                  = true
  use_update_dns_on_lease_renewal     = true

  // Discovery settings (removed read-only discovery_engine_type)
  same_port_control_discovery_blackout = false

  use_discovery_basic_polling_settings = true

  // Network management
  mgm_private     = true
  nextserver      = "192.168.1.11"
  unmanaged       = false
  use_mgm_private = true
  use_nextserver  = true

  // Email and notification settings
  email_list               = ["admin@example.com", "network@example.com"]
  ipam_email_addresses     = ["ipam@example.com"]
  use_email_list           = true
  use_ipam_email_addresses = true

  // Water mark settings
  high_water_mark       = 95
  high_water_mark_reset = 85
  low_water_mark        = 10
  low_water_mark_reset  = 20

  // Ignore settings (fixed ignore_id value)
  ignore_id            = "MACADDR" // Valid values: "CLIENT", "MACADDR", "NONE"
  ignore_mac_addresses = ["00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"]
  use_ignore_id        = true

  // Operation flags
  delete_reason     = "Network reorganization"
  restart_if_needed = true


  // Discovery basic poll settings (nested object)
  discovery_basic_poll_settings = {
    auto_arp_refresh_before_switch_port_polling : true,
    cli_collection : true,
    complete_ping_sweep : false,
    credential_group : "default",
    device_profile : false,
    netbios_scanning : false,
    polling_frequency_modifier : "1",
    port_scanning : false,
    smart_subnet_ping_sweep : false,
    snmp_collection : true,
    switch_port_data_collection_polling : "PERIODIC",
    switch_port_data_collection_polling_interval : 3600,
    use_global_polling_frequency_modifier : true
  }

  // Discovery blackout setting (nested object)
  discovery_blackout_setting = {
    enable_blackout = false
  }
  use_blackout_setting = true

  // IPAM threshold settings (nested object)
  ipam_threshold_settings = {
    reset_value : 90,
    trigger_value : 100
  }
  use_ipam_threshold_settings = true

  // IPAM trap settings (nested object) 
  ipam_trap_settings = {
    enable_email_warnings : true,
    enable_snmp_warnings : true
  }
  use_ipam_trap_settings = true

  # // DHCP options (list of nested objects)
  # options = [
  #   {
  #     name         = "routers"
  #     num          = 3
  #     value        = "10.0.0.1"
  #     use_option   = true
  #     vendor_class = "DHCP"
  #   },
  #   {
  #     name         = "domain-name-servers"
  #     num          = 6
  #     value        = "8.8.8.8,8.8.4.4"
  #     use_option   = true
  #     vendor_class = "DHCP"
  #   }
  # ]
  # use_options = true

  # // Subscribe settings (nested object)
  # subscribe_settings = {
  #   enabled_attributes = ["USERNAME", "SSID", "VLAN"]
  #   mapped_ea_attributes = [
  #     {
  #       name      = "Username"
  #       mapped_ea = "User"
  #     },
  #     {
  #       name      = "SSID"
  #       mapped_ea = "WiFiNetwork"
  #     }
  #   ]
  # }
  # use_subscribe_settings = true

  // Extensible attributes
  extattrs = {
    "Site" = "DataCenter1"
  }
}


resource "nios_ipam_networkcontainer" "example_func_call" {
  func_call = {
    attribute_name  = "network"
    object_function = "next_available_network"
    result_field    = "networks"
    object          = "networkcontainer"
    object_parameters = {
      network      = "10.0.0.0/24"
      network_view = "default"
    }
    parameters = {
      cidr = 28
    }
  }
  comment = "Network container created with function call"
  depends_on = [
    nios_ipam_networkcontainer.example_container
  ]
}
