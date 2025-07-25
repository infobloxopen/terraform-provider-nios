---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nios_ipam_network_container Data Source - nios"
subcategory: "IPAM"
description: |-
  Retrieves Information about existing Network Containers.
---

# nios_ipam_network_container (Data Source)

Retrieves Information about existing Network Containers.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `extattrfilters` (Map of String) External Attribute Filters are used to return a more specific list of results by filtering on external attributes. If you specify multiple filters, the results returned will have only resources that match all the specified filters.
- `filters` (Map of String) Filter are used to return a more specific list of results. Filters can be used to match resources by specific attributes, e.g. name. If you specify multiple filters, the results returned will have only resources that match all the specified filters.
- `max_results` (Number) Maximum number of objects to be returned. Defaults to 1000.
- `paging` (Number) Enable (1) or disable (0) paging for the data source query. When enabled, the system retrieves results in pages, allowing efficient handling of large result sets. Paging is enabled by default.

### Read-Only

- `result` (Attributes List) (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Optional:

- `authority` (Boolean) Authority for the DHCP network container.
- `auto_create_reversezone` (Boolean) This flag controls whether reverse zones are automatically created when the network is added.
- `bootfile` (String) The boot server IPv4 Address or name in FQDN format for the network container. You can specify the name and/or IP address of the boot server that the host needs to boot.
- `bootserver` (String) The bootserver address for the network container. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format.
- `cloud_info` (Attributes) Structure containing all cloud API related information for this object. (see [below for nested schema](#nestedatt--result--cloud_info))
- `comment` (String) Comment for the network container; maximum 256 characters.
- `ddns_domainname` (String) The dynamic DNS domain name the appliance uses specifically for DDNS updates for this network container.
- `ddns_generate_hostname` (Boolean) If this field is set to True, the DHCP server generates a hostname and updates DNS with it when the DHCP client request does not contain a hostname.
- `ddns_server_always_updates` (Boolean) This field controls whether the DHCP server is allowed to update DNS, regardless of the DHCP client requests. Note that changes for this field take effect only if ddns_use_option81 is True.
- `ddns_ttl` (Number) The DNS update Time to Live (TTL) value of a DHCP network container object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.
- `ddns_update_fixed_addresses` (Boolean) By default, the DHCP server does not update DNS when it allocates a fixed address to a client. You can configure the DHCP server to update the A and PTR records of a client with a fixed address. When this feature is enabled and the DHCP server adds A and PTR records for a fixed address, the DHCP server never discards the records.
- `ddns_use_option81` (Boolean) The support for DHCP Option 81 at the network container level.
- `delete_reason` (String) The reason for deleting the RIR registration request.
- `deny_bootp` (Boolean) If set to True, BOOTP settings are disabled and BOOTP requests will be denied.
- `discovery_basic_poll_settings` (Attributes) The discovery basic poll settings for this network container. (see [below for nested schema](#nestedatt--result--discovery_basic_poll_settings))
- `discovery_blackout_setting` (Attributes) The discovery blackout setting for this network container. (see [below for nested schema](#nestedatt--result--discovery_blackout_setting))
- `discovery_member` (String) The member that will run discovery for this network container.
- `email_list` (List of String) The e-mail lists to which the appliance sends DHCP threshold alarm e-mail messages.
- `enable_ddns` (Boolean) The dynamic DNS updates flag of a DHCP network container object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.
- `enable_dhcp_thresholds` (Boolean) Determines if DHCP thresholds are enabled for the network container.
- `enable_discovery` (Boolean) Determines whether a discovery is enabled or not for this network container. When this is set to False, the network container discovery is disabled.
- `enable_email_warnings` (Boolean) Determines if DHCP threshold warnings are sent through email.
- `enable_immediate_discovery` (Boolean) Determines if the discovery for the network container should be immediately enabled.
- `enable_pxe_lease_time` (Boolean) Set this to True if you want the DHCP server to use a different lease time for PXE clients.
- `enable_snmp_warnings` (Boolean) Determines if DHCP threshold warnings are send through SNMP.
- `extattrs` (Map of String) Extensible attributes associated with the object.
- `federated_realms` (Attributes List) This field contains the federated realms associated to this network container. (see [below for nested schema](#nestedatt--result--federated_realms))
- `func_call` (Attributes) Function call to be executed. (see [below for nested schema](#nestedatt--result--func_call))
- `high_water_mark` (Number) The percentage of DHCP network container usage threshold above which network container usage is not expected and may warrant your attention. When the high watermark is reached, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.
- `high_water_mark_reset` (Number) The percentage of DHCP network container usage below which the corresponding SNMP trap is reset. A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the high watermark value.
- `ignore_dhcp_option_list_request` (Boolean) If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested.
- `ignore_id` (String) Indicates whether the appliance will ignore DHCP client IDs or MAC addresses.
- `ignore_mac_addresses` (List of String) A list of MAC addresses the appliance will ignore.
- `ipam_email_addresses` (List of String) The e-mail lists to which the appliance sends IPAM threshold alarm e-mail messages.
- `ipam_threshold_settings` (Attributes) The IPAM threshold settings for this network container. (see [below for nested schema](#nestedatt--result--ipam_threshold_settings))
- `ipam_trap_settings` (Attributes) The IPAM trap settings for this network container. (see [below for nested schema](#nestedatt--result--ipam_trap_settings))
- `lease_scavenge_time` (Number) An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day).
- `logic_filter_rules` (Attributes List) This field contains the logic filters to be applied on the this network container. This list corresponds to the match rules that are written to the dhcpd configuration file. (see [below for nested schema](#nestedatt--result--logic_filter_rules))
- `low_water_mark` (Number) The percentage of DHCP network container usage below which the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.
- `low_water_mark_reset` (Number) The percentage of DHCP network container usage threshold below which network container usage is not expected and may warrant your attention. When the low watermark is crossed, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low watermark reset value must be higher than the low watermark value.
- `mgm_private` (Boolean) This field controls whether this object is synchronized with the Multi-Grid Master. If this field is set to True, objects are not synchronized.
- `ms_ad_user_data` (Attributes) (see [below for nested schema](#nestedatt--result--ms_ad_user_data))
- `network` (String) The IPv4 Address of the record.
- `network_view` (String) The name of the network view in which this network resides.
- `nextserver` (String) The name in FQDN and/or IPv4 Address of the next server that the host needs to boot.
- `options` (Attributes List) An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. (see [below for nested schema](#nestedatt--result--options))
- `port_control_blackout_setting` (Attributes) The port control blackout setting for this network container. (see [below for nested schema](#nestedatt--result--port_control_blackout_setting))
- `pxe_lease_time` (Number) The PXE lease time value of a DHCP Network container object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.
- `recycle_leases` (Boolean) If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the leases are permanently deleted.
- `remove_subnets` (Boolean) Remove subnets delete option. Determines whether all child objects should be removed alongside with the network container or child objects should be assigned to another parental container. By default child objects are deleted with the network container.
- `restart_if_needed` (Boolean) Restarts the member service.
- `rir_organization` (String) The RIR organization assoicated with the network container.
- `rir_registration_action` (String) The RIR registration action.
- `rir_registration_status` (String) The registration status of the network container in RIR.
- `same_port_control_discovery_blackout` (Boolean) If the field is set to True, the discovery blackout setting will be used for port control blackout setting.
- `send_rir_request` (Boolean) Determines whether to send the RIR registration request.
- `subscribe_settings` (Attributes) (see [below for nested schema](#nestedatt--result--subscribe_settings))
- `unmanaged` (Boolean) Determines whether the network container is unmanaged or not.
- `update_dns_on_lease_renewal` (Boolean) This field controls whether the DHCP server updates DNS when a DHCP lease is renewed.
- `use_authority` (Boolean) Use flag for: authority
- `use_blackout_setting` (Boolean) Use flag for: discovery_blackout_setting , port_control_blackout_setting, same_port_control_discovery_blackout
- `use_bootfile` (Boolean) Use flag for: bootfile
- `use_bootserver` (Boolean) Use flag for: bootserver
- `use_ddns_domainname` (Boolean) Use flag for: ddns_domainname
- `use_ddns_generate_hostname` (Boolean) Use flag for: ddns_generate_hostname
- `use_ddns_ttl` (Boolean) Use flag for: ddns_ttl
- `use_ddns_update_fixed_addresses` (Boolean) Use flag for: ddns_update_fixed_addresses
- `use_ddns_use_option81` (Boolean) Use flag for: ddns_use_option81
- `use_deny_bootp` (Boolean) Use flag for: deny_bootp
- `use_discovery_basic_polling_settings` (Boolean) Use flag for: discovery_basic_poll_settings
- `use_email_list` (Boolean) Use flag for: email_list
- `use_enable_ddns` (Boolean) Use flag for: enable_ddns
- `use_enable_dhcp_thresholds` (Boolean) Use flag for: enable_dhcp_thresholds
- `use_enable_discovery` (Boolean) Use flag for: discovery_member , enable_discovery
- `use_ignore_dhcp_option_list_request` (Boolean) Use flag for: ignore_dhcp_option_list_request
- `use_ignore_id` (Boolean) Use flag for: ignore_id
- `use_ipam_email_addresses` (Boolean) Use flag for: ipam_email_addresses
- `use_ipam_threshold_settings` (Boolean) Use flag for: ipam_threshold_settings
- `use_ipam_trap_settings` (Boolean) Use flag for: ipam_trap_settings
- `use_lease_scavenge_time` (Boolean) Use flag for: lease_scavenge_time
- `use_logic_filter_rules` (Boolean) Use flag for: logic_filter_rules
- `use_mgm_private` (Boolean) Use flag for: mgm_private
- `use_nextserver` (Boolean) Use flag for: nextserver
- `use_options` (Boolean) Use flag for: options
- `use_pxe_lease_time` (Boolean) Use flag for: pxe_lease_time
- `use_recycle_leases` (Boolean) Use flag for: recycle_leases
- `use_subscribe_settings` (Boolean) Use flag for: subscribe_settings
- `use_update_dns_on_lease_renewal` (Boolean) Use flag for: update_dns_on_lease_renewal
- `use_zone_associations` (Boolean) Use flag for: zone_associations
- `zone_associations` (Attributes List) The list of zones associated with this network. (see [below for nested schema](#nestedatt--result--zone_associations))

Read-Only:

- `discover_now_status` (String) Discover now status for this network container.
- `discovery_engine_type` (String) The network discovery engine type.
- `endpoint_sources` (List of String) The endpoints that provides data for the DHCP Network Container object.
- `extattrs_all` (Map of String) Extensible attributes associated with the object , including default attributes.
- `last_rir_registration_update_sent` (Number) The timestamp when the last RIR registration update was sent.
- `last_rir_registration_update_status` (String) Last RIR registration update status.
- `mgm_private_overridable` (Boolean) This field is assumed to be True unless filled by any conforming objects, such as Network, IPv6 Network, Network Container, IPv6 Network Container, and Network View. This value is set to False if mgm_private is set to True in the parent object.
- `network_container` (String) The network container to which this network belongs, if any.
- `ref` (String) The reference to the object.
- `rir` (String) The registry (RIR) that allocated the network container address space.
- `utilization` (Number) The network container utilization in percentage.

<a id="nestedatt--result--cloud_info"></a>
### Nested Schema for `result.cloud_info`

Optional:

- `delegated_member` (Attributes) (see [below for nested schema](#nestedatt--result--cloud_info--delegated_member))

Read-Only:

- `authority_type` (String) Type of authority over the object.
- `delegated_root` (String) Indicates the root of the delegation if delegated_scope is SUBTREE or RECLAIMING. This is not set otherwise.
- `delegated_scope` (String) Indicates the scope of delegation for the object. This can be one of the following: NONE (outside any delegation), ROOT (the delegation point), SUBTREE (within the scope of a delegation), RECLAIMING (within the scope of a delegation being reclaimed, either as the delegation point or in the subtree).
- `mgmt_platform` (String) Indicates the specified cloud management platform.
- `owned_by_adaptor` (Boolean) Determines whether the object was created by the cloud adapter or not.
- `tenant` (String) Reference to the tenant object associated with the object, if any.
- `usage` (String) Indicates the cloud origin of the object.

<a id="nestedatt--result--cloud_info--delegated_member"></a>
### Nested Schema for `result.cloud_info.delegated_member`

Optional:

- `ipv4addr` (String) The IPv4 Address of the Grid Member.
- `ipv6addr` (String) The IPv6 Address of the Grid Member.
- `name` (String) The Grid member name



<a id="nestedatt--result--discovery_basic_poll_settings"></a>
### Nested Schema for `result.discovery_basic_poll_settings`

Optional:

- `auto_arp_refresh_before_switch_port_polling` (Boolean) Determines whether auto ARP refresh before switch port polling is enabled or not.
- `cli_collection` (Boolean) Determines whether CLI collection is enabled or not.
- `complete_ping_sweep` (Boolean) Determines whether complete ping sweep is enabled or not.
- `credential_group` (String) Credential group.
- `device_profile` (Boolean) Determines whether device profile is enabled or not.
- `netbios_scanning` (Boolean) Determines whether netbios scanning is enabled or not.
- `polling_frequency_modifier` (String) Polling Frequency Modifier.
- `port_scanning` (Boolean) Determines whether port scanning is enabled or not.
- `smart_subnet_ping_sweep` (Boolean) Determines whether smart subnet ping sweep is enabled or not.
- `snmp_collection` (Boolean) Determines whether SNMP collection is enabled or not.
- `switch_port_data_collection_polling` (String) A switch port data collection polling mode.
- `switch_port_data_collection_polling_interval` (Number) Indicates the interval for switch port data collection polling.
- `switch_port_data_collection_polling_schedule` (Attributes) (see [below for nested schema](#nestedatt--result--discovery_basic_poll_settings--switch_port_data_collection_polling_schedule))
- `use_global_polling_frequency_modifier` (Boolean) Use Global Polling Frequency Modifier.

<a id="nestedatt--result--discovery_basic_poll_settings--switch_port_data_collection_polling_schedule"></a>
### Nested Schema for `result.discovery_basic_poll_settings.switch_port_data_collection_polling_schedule`

Optional:

- `day_of_month` (Number) The day of the month for the scheduled task.
- `disable` (Boolean) If set to True, the scheduled task is disabled.
- `every` (Number) The number of frequency to wait before repeating the scheduled task.
- `frequency` (String) The frequency for the scheduled task.
- `hour_of_day` (Number) The hour of day for the scheduled task.
- `minutes_past_hour` (Number) The minutes past the hour for the scheduled task.
- `month` (Number) The month for the scheduled task.
- `recurring_time` (Number) The recurring time for the schedule in Epoch seconds format. This field is obsolete and is preserved only for backward compatibility purposes. Please use other applicable fields to define the recurring schedule. DO NOT use recurring_time together with these fields. If you use recurring_time with other fields to define the recurring schedule, recurring_time has priority over year, hour_of_day, and minutes_past_hour and will override the values of these fields, although it does not override month and day_of_month. In this case, the recurring time value might be different than the intended value that you define.
- `repeat` (String) Indicates if the scheduled task will be repeated or run only once.
- `time_zone` (String) The time zone for the schedule.
- `weekdays` (List of String) Days of the week when scheduling is triggered.
- `year` (Number) The year for the scheduled task.



<a id="nestedatt--result--discovery_blackout_setting"></a>
### Nested Schema for `result.discovery_blackout_setting`

Optional:

- `blackout_duration` (Number) The blackout duration in seconds; minimum value is 1 minute.
- `blackout_schedule` (Attributes) (see [below for nested schema](#nestedatt--result--discovery_blackout_setting--blackout_schedule))
- `enable_blackout` (Boolean) Determines whether a blackout is enabled or not.

<a id="nestedatt--result--discovery_blackout_setting--blackout_schedule"></a>
### Nested Schema for `result.discovery_blackout_setting.blackout_schedule`

Optional:

- `day_of_month` (Number) The day of the month for the scheduled task.
- `disable` (Boolean) If set to True, the scheduled task is disabled.
- `every` (Number) The number of frequency to wait before repeating the scheduled task.
- `frequency` (String) The frequency for the scheduled task.
- `hour_of_day` (Number) The hour of day for the scheduled task.
- `minutes_past_hour` (Number) The minutes past the hour for the scheduled task.
- `month` (Number) The month for the scheduled task.
- `recurring_time` (Number) The recurring time for the schedule in Epoch seconds format. This field is obsolete and is preserved only for backward compatibility purposes. Please use other applicable fields to define the recurring schedule. DO NOT use recurring_time together with these fields. If you use recurring_time with other fields to define the recurring schedule, recurring_time has priority over year, hour_of_day, and minutes_past_hour and will override the values of these fields, although it does not override month and day_of_month. In this case, the recurring time value might be different than the intended value that you define.
- `repeat` (String) Indicates if the scheduled task will be repeated or run only once.
- `time_zone` (String) The time zone for the schedule.
- `weekdays` (List of String) Days of the week when scheduling is triggered.
- `year` (Number) The year for the scheduled task.



<a id="nestedatt--result--federated_realms"></a>
### Nested Schema for `result.federated_realms`

Optional:

- `id` (String) The federated realm id
- `name` (String) The federated realm name


<a id="nestedatt--result--func_call"></a>
### Nested Schema for `result.func_call`

Required:

- `attribute_name` (String) The attribute to be called.

Optional:

- `object` (String) The object to be called.
- `object_function` (String) The function to be called.
- `object_parameters` (Map of String) The parameters for the object.
- `parameters` (Map of String) The parameters for the function.
- `result_field` (String) The result field of the function.


<a id="nestedatt--result--ipam_threshold_settings"></a>
### Nested Schema for `result.ipam_threshold_settings`

Optional:

- `reset_value` (Number) Indicates the percentage point which resets the email/SNMP trap sending.
- `trigger_value` (Number) Indicates the percentage point which triggers the email/SNMP trap sending.


<a id="nestedatt--result--ipam_trap_settings"></a>
### Nested Schema for `result.ipam_trap_settings`

Optional:

- `enable_email_warnings` (Boolean) Determines whether sending warnings by email is enabled or not.
- `enable_snmp_warnings` (Boolean) Determines whether sending warnings by SNMP is enabled or not.


<a id="nestedatt--result--logic_filter_rules"></a>
### Nested Schema for `result.logic_filter_rules`

Optional:

- `filter` (String) The filter name.
- `type` (String) The filter type. Valid values are: * MAC * NAC * Option


<a id="nestedatt--result--ms_ad_user_data"></a>
### Nested Schema for `result.ms_ad_user_data`

Read-Only:

- `active_users_count` (Number) The number of active users.


<a id="nestedatt--result--options"></a>
### Nested Schema for `result.options`

Optional:

- `name` (String) Name of the DHCP option.
- `num` (Number) The code of the DHCP option.
- `use_option` (Boolean) Only applies to special options that are displayed separately from other options and have a use flag. These options are: * routers * router-templates * domain-name-servers * domain-name * broadcast-address * broadcast-address-offset * dhcp-lease-time * dhcp6.name-servers
- `value` (String) Value of the DHCP option
- `vendor_class` (String) The name of the space this DHCP option is associated to.


<a id="nestedatt--result--port_control_blackout_setting"></a>
### Nested Schema for `result.port_control_blackout_setting`

Optional:

- `blackout_duration` (Number) The blackout duration in seconds; minimum value is 1 minute.
- `blackout_schedule` (Attributes) (see [below for nested schema](#nestedatt--result--port_control_blackout_setting--blackout_schedule))
- `enable_blackout` (Boolean) Determines whether a blackout is enabled or not.

<a id="nestedatt--result--port_control_blackout_setting--blackout_schedule"></a>
### Nested Schema for `result.port_control_blackout_setting.blackout_schedule`

Optional:

- `day_of_month` (Number) The day of the month for the scheduled task.
- `disable` (Boolean) If set to True, the scheduled task is disabled.
- `every` (Number) The number of frequency to wait before repeating the scheduled task.
- `frequency` (String) The frequency for the scheduled task.
- `hour_of_day` (Number) The hour of day for the scheduled task.
- `minutes_past_hour` (Number) The minutes past the hour for the scheduled task.
- `month` (Number) The month for the scheduled task.
- `recurring_time` (Number) The recurring time for the schedule in Epoch seconds format. This field is obsolete and is preserved only for backward compatibility purposes. Please use other applicable fields to define the recurring schedule. DO NOT use recurring_time together with these fields. If you use recurring_time with other fields to define the recurring schedule, recurring_time has priority over year, hour_of_day, and minutes_past_hour and will override the values of these fields, although it does not override month and day_of_month. In this case, the recurring time value might be different than the intended value that you define.
- `repeat` (String) Indicates if the scheduled task will be repeated or run only once.
- `time_zone` (String) The time zone for the schedule.
- `weekdays` (List of String) Days of the week when scheduling is triggered.
- `year` (Number) The year for the scheduled task.



<a id="nestedatt--result--subscribe_settings"></a>
### Nested Schema for `result.subscribe_settings`

Optional:

- `enabled_attributes` (List of String) The list of Cisco ISE attributes allowed for subscription.
- `mapped_ea_attributes` (Attributes List) The list of NIOS extensible attributes to Cisco ISE attributes mappings. (see [below for nested schema](#nestedatt--result--subscribe_settings--mapped_ea_attributes))

<a id="nestedatt--result--subscribe_settings--mapped_ea_attributes"></a>
### Nested Schema for `result.subscribe_settings.mapped_ea_attributes`

Optional:

- `mapped_ea` (String) The name of the extensible attribute definition object the Cisco ISE attribute that is enabled for subscription is mapped on.
- `name` (String) The Cisco ISE attribute name that is enabled for publishsing from a Cisco ISE endpoint.



<a id="nestedatt--result--zone_associations"></a>
### Nested Schema for `result.zone_associations`

Optional:

- `fqdn` (String) The FQDN of the authoritative forward zone.
- `is_default` (Boolean) True if this is the default zone.
- `view` (String) The view to which the zone belongs. If a view is not specified, the default view is used.
