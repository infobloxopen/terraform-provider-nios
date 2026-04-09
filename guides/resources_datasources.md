The tables below list all available resources and data sources

### DHCP

| Name | Resource Description                         | Data Source Description                                        |
|----------|----------------------------------------------|----------------------------------------------------------------|
| `nios_dhcp_fixed_address` | Manages DHCP Fixed Address (IPv4) resources  | Retrieves information about existing DHCP fixed addresses      |
| `nios_dhcp_range` | Manages DHCP Range (IPv4) resources          | Retrieves information about existing DHCP ranges               |
| `nios_dhcp_range_template` | Manages DHCP Range Template (IPv4) resources | Retrieves information about existing DHCP range templates      |
| `nios_dhcp_shared_network` | Manages DHCP Shared Network (IPv4) resources | Retrieves information about existing DHCP shared networks      |
| `nios_dhcp_ipv6_range_template` | Manages DHCP IPV6 Range Template resources   | Retrieves information about existing DHCP IPV6 Range Templates |                           |                                              |                                                           |
| `nios_dhcp_ipv6dhcpoptiondefinition` | Manages DHCP IPv6 option definition | Retrieves information about existing IPv6 option definitions      |
| `nios_dhcp_ipv6dhcpoptionspace` | Manages DHCP IPv6 option space | Retrieves information about existing IPv6 option spaces      |
| `nios_dhcp_ipv6fixedaddresstemplate` | Manages DHCP IPv6 fixed address template | Retrieves information about existing IPv6 fixed address templates      |
| `nios_dhcp_filteroption` | Manages DHCP filter options | Retrieves information about existing filter option configurations |
| `nios_dhcp_filterrelayagent` | Manages DHCP filter relay agent | Retrieves information about existing relay agent configurations |
| `nios_dhcp_ipv6filteroption` | Manages DHCP IPv6 filter options | Retrieves information about existing IPv6 filter option configurations |
| `nios_dhcp_ipv6fixedaddress` | Manages DHCP IPv6 fixed addresses | Retrieves information about existing IPv6 fixed address data |
| `nios_dhcp_fingerprint` | Manages DHCP fingerprints | Retrieves information about existing fingerprint data |
| `nios_dhcp_filterfingerprint` | Manages DHCP filter fingerprints | Retrieves information about existing filter fingerprint configurations |
| `nios_dhcp_roaminghost` | Manages DHCP roaming hosts | Retrieves information about existing roaming host configurations |
| `nios_dhcp_filternac` | Manages DHCP filter NAC | Retrieves information about existing NAC filter configurations |
| `nios_dhcp_macfilteraddress` | Manages DHCP MAC filter addresses | Retrieves information about existing MAC filter address data |

### DNS

| Name                                 | Resource Description                      | Data Source Description                                              |
|--------------------------------------|-------------------------------------------|----------------------------------------------------------------------|
| `nios_dns_view`                      | Manages DNS Views                         | Retrieves information about existing DNS views                       |
| `nios_dns_zone_auth`                 | Manages Authoritative DNS zones           | Retrieves information about existing authoritative DNS zones         |
| `nios_dns_zone_delegated`            | Manages Delegated DNS zones               | Retrieves information about existing delegated DNS zones             |
| `nios_dns_zone_forward`              | Manages Forwarding DNS zones              | Retrieves information about existing forwarding DNS zones            |
| `nios_dns_record_a`                  | Manages DNS A records                     | Retrieves information about existing DNS A records                   |
| `nios_dns_record_aaaa`               | Manages DNS AAAA records                  | Retrieves information about existing DNS AAAA records                |
| `nios_dns_record_alias`              | Manages DNS ALIAS records                 | Retrieves information about existing DNS ALIAS records               |
| `nios_dns_record_cname`              | Manages DNS CNAME records                 | Retrieves information about existing DNS CNAME records               |
| `nios_dns_record_mx`                 | Manages DNS MX records                    | Retrieves information about existing DNS MX records                  |
| `nios_dns_record_ns`                 | Manages DNS NS records                    | Retrieves information about existing DNS NS records                  |
| `nios_dns_record_ptr`                | Manages DNS PTR records                   | Retrieves information about existing DNS PTR records                 |
| `nios_dns_record_srv`                | Manages DNS SRV records                   | Retrieves information about existing DNS SRV records                 |
| `nios_dns_record_txt`                | Manages DNS TXT records                   | Retrieves information about existing DNS TXT records                 |
| `nios_dns_record_dname`              | Manages DNS DNAME records                 | Retrieves information about existing DNS DNAME records               |
| `nios_dns_record_naptr`              | Manages DNS NAPTR records                 | Retrieves information about existing DNS NAPTR records               |
| `nios_dns_record_tlsa`               | Manages DNS TLSA records                  | Retrieves information about existing DNS TLSA records                |
| `nios_dns_record_caa`                | Manages DNS CAA records                   | Retrieves information about existing DNS CAA records                 |
| `nios_dns_record_unknown`            | Manages DNS Unknown records               | Retrieves information about existing DNS Unknown records             |
| `nios_dns_zone_rp`                   | Manages DNS Response Policy Zones         | Retrieves information about existing DNS Response Policy Zones       |
| `nios_dns_zone_stub`                 | Manages DNS DNS Stub Zones                | Retrieves information about existing DNS Stub Zones                  |
| `nios_dns_nsgroup`                   | Manages DNS NS Groups                     | Retrieves information about existing DNS NS Groups                   |
| `nios_dns_nsgroup_delegation`        | Manages DNS NS Group Delegations          | Retrieves information about existing DNS NS Group Delegations        |
| `nios_dns_nsgroup_forwardingmember`  | Manages DNS NS Group Forwarding Members   | Retrieves information about existing DNS NS Group Forwarding Member  |
| `nios_dns_nsgroup_forwardstubserver` | Manages DNS NS Group Forward Stub Servers | Retrieves information about existing DNS NS Group Forward Stub Servers |
| `nios_dns_nsgroup_stubmember`        | Manages DNS NS Group Stub Members         | Retrieves information about existing DNS NS Group Stub Members       |
| `nios_ip_allocation`                 | Manages an IP allocation                  |                                                                      |
| `nios-ip_association`                | Manages an IP association                 |                                                                      |
| `nios_host_record`                   |                                           | Retrieves information about existing Host Records                    |
| `nios_dns_sharedrecordgroup`         | Manages Shared Record Group               | Retrieves information about existing Shared Record Groups            |
| `nios_dns_sharedrecord_txt`          | Manages Shared Record TXT                 | Retrieves information about existing DNS Shared TXT Records          |


### DTC

| Name | Resource Description         | Data Source Description                          |
|----------|------------------------------|--------------------------------------------------|
| `nios_dtc_lbdn` | Manages DTC LBDN resources   | Retrieves information about existing DTC LBDNs   |
| `nios_dtc_pool` | Manages DTC Pool resources   | Retrieves information about existing DTC Pools   |
| `nios_dtc_server` | Manages DTC Server resources | Retrieves information about existing DTC Servers |
| `nios_dtc_record_a` | Manages DTC A records | Retrieves information about existing DTC A record data |
| `nios_dtc_record_aaaa` | Manages DTC AAAA records | Retrieves information about existing DTC AAAA record data |
| `nios_dtc_record_cname` | Manages DTC CNAME records | Retrieves information about existing DTC CNAME record data |
| `nios_dtc_record_naptr` | Manages DTC NAPTR records | Retrieves information about existing DTC NAPTR record data |
| `nios_dtc_record_srv` | Manages DTC SRV records | Retrieves information about existing DTC SRV record data |
| `nios_dtc_topology_rule` |  | Retrieves information about existing DTC topology rule configurations |

### RPZ

| Name | Resource Description         | Data Source Description                          |
|----------|------------------------------|--------------------------------------------------|
| `nios_dns_record_rpz_a_ipaddress` | Manages RPZ A IP address records | Retrieves information about existing RPZ A IP address data |
| `nios_dns_record_rpz_aaaa` | Manages RPZ AAAA records | Retrieves information about existing RPZ AAAA record data |
| `nios_dns_record_rpz_aaaa_ipaddress` | Manages RPZ AAAA IP address records | Retrieves information about existing RPZ AAAA IP address data |
| `nios_dns_record_rpz_cname` | Manages RPZ CNAME records | Retrieves information about existing RPZ CNAME record data |
| `nios_dns_record_rpz_cname_clientipaddress` | Manages RPZ CNAME client IP address records | Retrieves information about existing RPZ CNAME client IP address data |
| `nios_dns_record_rpz_cname_clientipaddressdn` | Manages RPZ CNAME client IP address DN records | Retrieves information about existing RPZ CNAME client IP address DN data |
| `nios_dns_record_rpz_cname_ipaddressdn` | Manages RPZ CNAME IP address DN records | Retrieves information about existing RPZ CNAME IP address DN data |
| `nios_dns_record_rpz_txt` | Manages RPZ TXT records | Retrieves information about existing RPZ TXT record data |

### IPAM

| Name                              | Resource Description                 | Data Source Description                                           |
|-----------------------------------|--------------------------------------|-------------------------------------------------------------------|
| `nios_ipam_network_view`          | Manages IPAM Network Views           | Retrieves information about existing IPAM network views           |
| `nios_ipam_network`               | Manages IPAM Networks                | Retrieves information about existing IPAM networks                |
| `nios_ipam_network_container`     | Manages IPAM Network Containers      | Retrieves information about existing IPAM network containers      |
| `nios_ipam_ipv6network`           | Manages IPAM IPv6 Networks           | Retrieves information about existing IPAM IPv6 networks           |
| `nios_ipam_ipv6network_container` | Manages IPAM IPv6 Network Containers | Retrieves information about existing IPAM IPv6 network containers |
| `nios_ipam_bulk_hostname_template` | Manages IPAM Bulk Hostname Templates | Retrieves information about existing IPAM Bulk Hostname templates |
| `nios_ipam_ipv6networktemplate` | Manages IPAM IPv6 network templates | Retrieves information about existing IPv6 network template configurations |
| `nios_ipam_superhost` | Manages IPAM super hosts | Retrieves information about existing super host data |

### CLOUD

| Name                                | Resource Description                 | Data Source Description                                           |
|-------------------------------------|--------------------------------------|-------------------------------------------------------------------|
| `nios_cloud_aws_route53_task_group` | Manages AWS Users                | Retrieves information about existing AWS Users                |
| `nios_cloud_aws_user`               | Manages AWS Route 53 Task Groups | Retrieves information about existing AWS Route 53 Task Groups |

### MICROSOFT

| Name                                | Resource Description                 | Data Source Description                                           |
|-------------------------------------|--------------------------------------|-------------------------------------------------------------------|
| `nios_microsoft_superscope` | Manages Microsoft super scopes | Retrieves information about existing super scope configurations |
| `nios_microsoft_adsites` | Manages Microsoft AD sites | Retrieves information about existing AD site data |
| `nios_microsoft_server` | Manages Microsoft servers | Retrieves information about existing Microsoft server configurations |

### SECURITY

| Name                               | Resource Description         | Data Source Description                                                  |
|------------------------------------|------------------------------|--------------------------------------------------------------------------|
| `nios_security_admin_user`         | Manages Admin Users          | Retrieves information about existing Admin Users                         |
| `nios_security_admin_role`         | Manages Admin Roles          | Retrieves information about existing Admin ROles                         |
| `nios_security_admin_group`        | Manages Admin Groups         | Retrieves information about existing Admin Groups                        |
| `nios_security_permission`         | Manages Security Permissions | Retrieves information about existing Security Permissions                |
| `nios_security_ftpuser`            | Manages Security FTP User    | Retrieves information about existing Security FTP Users                  |
| `nios_security_snmp_user`          | Manages Security SNMP Users  | Retrieves information about existing Security SNMPUsers                  |
| `security_certificate_authservice` | Manages Security Certificate Authentication Services | Retrieves information about existing Certificate Authentication Services |
| `nios_security_radius_authservice` | Manages RADIUS authentication services | Retrieves information about existing RADIUS auth service configurations |
| `nios_security_tacacsplus_authservice` | Manages TACACS+ authentication services | Retrieves information about existing TACACS+ auth service data |
| `nios_security_ldap_authservice` | Manages LDAP authentication services | Retrieves information about existing LDAP auth service configurations |
| `nios_security_saml_authservice` | Manages SAML authentication services | Retrieves information about existing SAML auth service data |

### Misc

| Name                    | Resource Description  | Data Source Description                            |
|-------------------------|-----------------------|----------------------------------------------------|
| `nios_misc_ruleset`     | Manages Rule Sets     | Retrieves information about existing Rule Sets     |
| `nios_misc_bfdtemplate` | Manages BFD Templates | Retrieves information about existing BFD Templates |
| `nios_misc_tftp_filedir` | Manages TFTP file directories | Retrieves information about existing TFTP file directory configurations |
| `nios_misc_dxl_endpoint` | Manages DXL endpoints | Retrieves information about existing DXL endpoint data |

### SMARTFOLDER

| Name                        | Resource Description           | Data Source Description                                     |
|-----------------------------|--------------------------------|-------------------------------------------------------------|
| `nios_smartfolder_personal` | Manages Personal Smart Folders | Retrieves information about existing Personal Smart Folders |
| `nios_smartfolder_global`   | Manages Global Smart Folders   | Retrieves information about existing Global Smart Folders   |

### ACL

| Name                | Resource Description | Data Source Description |
|---------------------|----------|------------|
| `nios_acl_namedacl` | Manages Named Access Control Lists | Retrieves information about existing Named Access Control Lists |

### GRID

| Name                               | Resource Description                                     | Data Source Description                                                               |
|------------------------------------|----------------------------------------------------------|---------------------------------------------------------------------------------------|
| `nios_grid_natgroup`               | Manages Grid NAT Groups                                  | Retrieves information about existing Grid NAT Groups                                  |
| `nios_grid_extensibleattributedef` | Manages Grid Manages an Extensible Attribute definitions | Retrieves information about existing Grid Manages an Extensible Attribute definitions |
| `nios_grid_upgradegroup`           | Manages Grid Upgrade Groups                              | Retrieves information about existing Grid Upgrade Groups                              |
| `nios_grid_servicerestart_group`   | Manages Grid Service Restart Groups                      | Retrieves information about existing Grid Service Restart Groups                      |
| `nios_grid_distributionschedule`   | Manages Grid Distribution Schedules                      | Retrieves information about existing Grid Distribution Schedules                      |
| `nios_grid_member` | Manages grid members | Retrieves information about existing member configurations |
| `nios_grid_upgradeschedule` | Manages grid upgrade schedules | Retrieves information about existing upgrade schedule configurations |

### DISCOVERY

| Name | Resource Description                | Data Source Description                                           |
|----------|-------------------------------------|-------------------------------------------------------------------|
| `nios_discovery_credentialgroup` | Manages Discovery Credential Groups | Retrieves information about existing Discovery Credential Groups  |
| `nios_discovery_vdiscovery_task` | Manages Discovery vDiscovery Tasks  | Retrieves information about existing Discovery vDiscovery Tasks   |

### NOTIFICATION

| Name                              | Resource Description                | Data Source Description                                          |
|-----------------------------------|-------------------------------------|------------------------------------------------------------------|
| `nios_notification_rule`          | Manages Notification Rules          | Retrieves information about existing Notification Rules          |
| `nios_notification_rest_endpoint` | Manages Notification Rest Endpoints | Retrieves information about existing Notification Rest Endpoints |
| `nios_notification_syslog_endpoint` | Manages syslog notification endpoints | Retrieves information about existing syslog endpoint configurations |

### PARENTAL CONTROL

| Name                              | Resource Description                | Data Source Description                                          |
|-----------------------------------|-------------------------------------|------------------------------------------------------------------|
| `nios_parentalcontrol_subscriberrecord` | Manages subscriber records | Retrieves information about existing subscriber record data |
| `nios_parentalcontrol_subscribersite` | Manages subscriber sites | Retrieves information about existing subscriber site configurations |
| `nios_parentalcontrol_blockingpolicy` | Manages blocking policies | Retrieves information about existing blocking policy data |
| `nios_parentalcontrol_avp` | Manages parental control AVP | Retrieves information about existing AVP configurations |

### RIR

| Name                              | Resource Description                | Data Source Description                                          |
|-----------------------------------|-------------------------------------|------------------------------------------------------------------|
| `nios_rir_organization` | Manages RIR organizations | Retrieves information about existing RIR organization data |
