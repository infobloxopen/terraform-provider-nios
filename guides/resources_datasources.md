The tables below list all available resources and data sources

### DHCP

| Name | Resource Description                         | Data Source Description                                        |
|----------|----------------------------------------------|----------------------------------------------------------------|
| `nios_dhcp_fixed_address` | Manages DHCP Fixed Address (IPv4) Resources  | Retrieves information about existing DHCP Fixed Addresses      |
| `nios_dhcp_range` | Manages DHCP Range (IPv4) Resources          | Retrieves information about existing DHCP Ranges               |
| `nios_dhcp_range_template` | Manages DHCP Range Template (IPv4) Resources | Retrieves information about existing DHCP Range Templates      |
| `nios_dhcp_shared_network` | Manages DHCP Shared Network (IPv4) Resources | Retrieves information about existing DHCP Shared Networks      |
| `nios_dhcp_ipv6_range_template` | Manages DHCP IPV6 Range Template Resources   | Retrieves information about existing DHCP IPV6 Range Templates |                           |                                              |                                                           |
| `nios_dhcp_ipv6dhcpoptiondefinition` | Manages DHCP IPv6 Option Definition | Retrieves information about existing IPv6 Option Definitions      |
| `nios_dhcp_ipv6dhcpoptionspace` | Manages DHCP IPv6 Option Space | Retrieves information about existing IPv6 Option Spaces      |
| `nios_dhcp_ipv6fixedaddresstemplate` | Manages DHCP IPv6 Fixed Address Template | Retrieves information about existing IPv6 Fixed Address Templates      |
| `nios_dhcp_filteroption` | Manages DHCP Filter Options | Retrieves information about existing Filter Option Configurations |
| `nios_dhcp_filterrelayagent` | Manages DHCP Filter Relay Agent | Retrieves information about existing Relay Agents |
| `nios_dhcp_ipv6filteroption` | Manages DHCP IPv6 Filter Options | Retrieves information about existing IPv6 Filter Options |
| `nios_dhcp_ipv6fixedaddress` | Manages DHCP IPv6 Fixed Addresses | Retrieves information about existing IPv6 Fixed Addresses |
| `nios_dhcp_fingerprint` | Manages DHCP Fingerprints | Retrieves information about existing Fingerprints |
| `nios_dhcp_filterfingerprint` | Manages DHCP Filter Fingerprints | Retrieves information about existing Filter Fingerprints |
| `nios_dhcp_roaminghost` | Manages DHCP Roaming Hosts | Retrieves information about existing Roaming Hosts |
| `nios_dhcp_filternac` | Manages DHCP Filter NAC | Retrieves information about existing NAC Filters |
| `nios_dhcp_macfilteraddress` | Manages DHCP MAC Filter Addresses | Retrieves information about existing MAC Filter Addresses |

### DNS

| Name                                 | Resource Description                      | Data Source Description                                              |
|--------------------------------------|-------------------------------------------|----------------------------------------------------------------------|
| `nios_dns_view`                      | Manages DNS Views                         | Retrieves information about existing DNS Views                       |
| `nios_dns_zone_auth`                 | Manages Authoritative DNS Zones           | Retrieves information about existing Authoritative DNS Zones         |
| `nios_dns_zone_delegated`            | Manages Delegated DNS Zones               | Retrieves information about existing Delegated DNS Zones             |
| `nios_dns_zone_forward`              | Manages Forwarding DNS Zones              | Retrieves information about existing Forwarding DNS Zones            |
| `nios_dns_record_a`                  | Manages DNS A Records                     | Retrieves information about existing DNS A Records                   |
| `nios_dns_record_aaaa`               | Manages DNS AAAA Records                  | Retrieves information about existing DNS AAAA Records                |
| `nios_dns_record_alias`              | Manages DNS ALIAS Records                 | Retrieves information about existing DNS ALIAS Records               |
| `nios_dns_record_cname`              | Manages DNS CNAME Records                 | Retrieves information about existing DNS CNAME Records               |
| `nios_dns_record_mx`                 | Manages DNS MX Records                    | Retrieves information about existing DNS MX Records                  |
| `nios_dns_record_ns`                 | Manages DNS NS Records                    | Retrieves information about existing DNS NS Records                  |
| `nios_dns_record_ptr`                | Manages DNS PTR Records                   | Retrieves information about existing DNS PTR Records                 |
| `nios_dns_record_srv`                | Manages DNS SRV Records                   | Retrieves information about existing DNS SRV Records                 |
| `nios_dns_record_txt`                | Manages DNS TXT Records                   | Retrieves information about existing DNS TXT Records                 |
| `nios_dns_record_dname`              | Manages DNS DNAME Records                 | Retrieves information about existing DNS DNAME Records               |
| `nios_dns_record_naptr`              | Manages DNS NAPTR Records                 | Retrieves information about existing DNS NAPTR Records               |
| `nios_dns_record_tlsa`               | Manages DNS TLSA Records                  | Retrieves information about existing DNS TLSA Records                |
| `nios_dns_record_caa`                | Manages DNS CAA Records                   | Retrieves information about existing DNS CAA Records                 |
| `nios_dns_record_unknown`            | Manages DNS Unknown Records               | Retrieves information about existing DNS Unknown Records             |
| `nios_dns_zone_rp`                   | Manages DNS Response Policy Zones         | Retrieves information about existing DNS Response Policy Zones       |
| `nios_dns_zone_stub`                 | Manages DNS Stub Zones                | Retrieves information about existing DNS Stub Zones                  |
| `nios_dns_nsgroup`                   | Manages DNS NS Groups                     | Retrieves information about existing DNS NS Groups                   |
| `nios_dns_nsgroup_delegation`        | Manages DNS NS Group Delegations          | Retrieves information about existing DNS NS Group Delegations        |
| `nios_dns_nsgroup_forwardingmember`  | Manages DNS NS Group Forwarding Members   | Retrieves information about existing DNS NS Group Forwarding Member  |
| `nios_dns_nsgroup_forwardstubserver` | Manages DNS NS Group Forward Stub Servers | Retrieves information about existing DNS NS Group Forward Stub Servers |
| `nios_dns_nsgroup_stubmember`        | Manages DNS NS Group Stub Members         | Retrieves information about existing DNS NS Group Stub Members       |
| `nios_ip_allocation`                 | Manages an IP Allocation                  |                                                                      |
| `nios-ip_association`                | Manages an IP Association                 |                                                                      |
| `nios_host_record`                   |                                           | Retrieves information about existing Host Records                    |
| `nios_dns_sharedrecordgroup`         | Manages Shared Record Group               | Retrieves information about existing Shared Record Groups            |
| `nios_dns_sharedrecord_txt`          | Manages Shared Record TXT                 | Retrieves information about existing DNS Shared TXT Records          |
| `nios_dns_record_https`              | Manages DNS HTTPS Records                 | Retrieves information about existing HTTPS Records               |
| `nios_dns_record_svcb`               | Manages DNS SVCB Records                  | Retrieves information about existing SVCB Records                |


### DTC

| Name | Resource Description         | Data Source Description                          |
|----------|------------------------------|--------------------------------------------------|
| `nios_dtc_lbdn` | Manages DTC LBDN Resources   | Retrieves information about existing DTC LBDNs   |
| `nios_dtc_pool` | Manages DTC Pool Resources   | Retrieves information about existing DTC Pools   |
| `nios_dtc_server` | Manages DTC Server Resources | Retrieves information about existing DTC Servers |
| `nios_dtc_record_a` | Manages DTC A Records | Retrieves information about existing DTC A Records |
| `nios_dtc_record_aaaa` | Manages DTC AAAA Records | Retrieves information about existing DTC AAAA Records |
| `nios_dtc_record_cname` | Manages DTC CNAME Records | Retrieves information about existing DTC CNAME Records |
| `nios_dtc_record_naptr` | Manages DTC NAPTR Records | Retrieves information about existing DTC NAPTR Records |
| `nios_dtc_record_srv` | Manages DTC SRV Records | Retrieves information about existing DTC SRV Records |
| `nios_dtc_topology_rule` |  | Retrieves information about existing DTC Topology Rules |

### RPZ

| Name | Resource Description         | Data Source Description                          |
|----------|------------------------------|--------------------------------------------------|
| `nios_dns_record_rpz_a_ipaddress` | Manages RPZ A IP Address Records | Retrieves information about existing RPZ A IP Address Records |
| `nios_dns_record_rpz_aaaa` | Manages RPZ AAAA Records | Retrieves information about existing RPZ AAAA Records |
| `nios_dns_record_rpz_aaaa_ipaddress` | Manages RPZ AAAA IP Address Records | Retrieves information about existing RPZ AAAA IP Address Records |
| `nios_dns_record_rpz_cname` | Manages RPZ CNAME Records | Retrieves information about existing RPZ CNAME Records |
| `nios_dns_record_rpz_cname_clientipaddress` | Manages RPZ CNAME Client IP Address Records | Retrieves information about existing RPZ CNAME Client IP Address Records |
| `nios_dns_record_rpz_cname_clientipaddressdn` | Manages RPZ CNAME Client IP Address DN Records | Retrieves information about existing RPZ CNAME Client IP Address DN Records |
| `nios_dns_record_rpz_cname_ipaddressdn` | Manages RPZ CNAME IP Address DN Records | Retrieves information about existing RPZ CNAME IP Address DN Records |
| `nios_dns_record_rpz_txt` | Manages RPZ TXT Records | Retrieves information about existing RPZ TXT Records |

### IPAM

| Name                              | Resource Description                 | Data Source Description                                           |
|-----------------------------------|--------------------------------------|-------------------------------------------------------------------|
| `nios_ipam_network_view`          | Manages IPAM Network Views           | Retrieves information about existing IPAM Network Views           |
| `nios_ipam_network`               | Manages IPAM Networks                | Retrieves information about existing IPAM Networks                |
| `nios_ipam_network_container`     | Manages IPAM Network Containers      | Retrieves information about existing IPAM Network Containers      |
| `nios_ipam_ipv6network`           | Manages IPAM IPv6 Networks           | Retrieves information about existing IPAM IPv6 Networks           |
| `nios_ipam_ipv6network_container` | Manages IPAM IPv6 Network Containers | Retrieves information about existing IPAM IPv6 Network Containers |
| `nios_ipam_bulk_hostname_template` | Manages IPAM Bulk Hostname Templates | Retrieves information about existing IPAM Bulk Hostname Templates |
| `nios_ipam_ipv6networktemplate` | Manages IPAM IPv6 Network Templates | Retrieves information about existing IPv6 Network Templates |
| `nios_ipam_superhost` | Manages IPAM Super Hosts | Retrieves information about existing Super Hosts |

### CLOUD

| Name                                | Resource Description                 | Data Source Description                                           |
|-------------------------------------|--------------------------------------|-------------------------------------------------------------------|
| `nios_cloud_aws_route53_task_group` | Manages AWS Route 53 Task Groups | Retrieves information about existing AWS Route 53 Task Groups |
| `nios_cloud_aws_user`               | Manages AWS Users                | Retrieves information about existing AWS Users                |

### MICROSOFT

| Name                                | Resource Description                 | Data Source Description                                           |
|-------------------------------------|--------------------------------------|-------------------------------------------------------------------|
| `nios_microsoft_superscope` | Manages Microsoft Super Scopes | Retrieves information about existing Super Scopes  |
| `nios_microsoft_adsites` | Manages Microsoft AD Sites | Retrieves information about existing AD Sites |
| `nios_microsoft_server` | Manages Microsoft Servers | Retrieves information about existing Microsoft Servers |

### SECURITY

| Name                               | Resource Description         | Data Source Description                                                  |
|------------------------------------|------------------------------|--------------------------------------------------------------------------|
| `nios_security_admin_user`         | Manages Admin Users          | Retrieves information about existing Admin Users                         |
| `nios_security_admin_role`         | Manages Admin Roles          | Retrieves information about existing Admin Roles                         |
| `nios_security_admin_group`        | Manages Admin Groups         | Retrieves information about existing Admin Groups                        |
| `nios_security_permission`         | Manages Security Permissions | Retrieves information about existing Security Permissions                |
| `nios_security_ftpuser`            | Manages Security FTP User    | Retrieves information about existing Security FTP Users                  |
| `nios_security_snmp_user`          | Manages Security SNMP Users  | Retrieves information about existing Security SNMP Users                  |
| `security_certificate_authservice` | Manages Security Certificate Authentication Services | Retrieves information about existing Certificate Authentication Services |
| `nios_security_radius_authservice` | Manages RADIUS Authentication Services | Retrieves information about existing RADIUS Authentication Services |
| `nios_security_tacacsplus_authservice` | Manages TACACS+ Authentication Services | Retrieves information about existing TACACS+ Authentication Services |
| `nios_security_ldap_authservice` | Manages LDAP Authentication Services | Retrieves information about existing LDAP Authentication Services |
| `nios_security_saml_authservice` | Manages SAML Authentication Services | Retrieves information about existing SAML Authentication Services |

### Misc

| Name                    | Resource Description  | Data Source Description                            |
|-------------------------|-----------------------|----------------------------------------------------|
| `nios_misc_ruleset`     | Manages Rule Sets     | Retrieves information about existing Rule Sets     |
| `nios_misc_bfdtemplate` | Manages BFD Templates | Retrieves information about existing BFD Templates |
| `nios_misc_tftp_filedir` | Manages TFTP File Directories | Retrieves information about existing TFTP File Directories |
| `nios_misc_dxl_endpoint` | Manages DXL Endpoints | Retrieves information about existing DXL Endpoints |

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
| `nios_grid_extensibleattributedef` | Manages Grid Extensible Attribute Definitions | Retrieves information about existing Grid Extensible Attribute Definitions |
| `nios_grid_upgradegroup`           | Manages Grid Upgrade Groups                              | Retrieves information about existing Grid Upgrade Groups                              |
| `nios_grid_servicerestart_group`   | Manages Grid Service Restart Groups                      | Retrieves information about existing Grid Service Restart Groups                      |
| `nios_grid_distributionschedule`   | Manages Grid Distribution Schedule                     | Retrieves information about current Distribution Schedule                       |
| `nios_grid_member` | Manages Grid Members | Retrieves information about existing Grid Members |
| `nios_grid_upgradeschedule` | Manages Grid Upgrade Schedule | Retrieves information about current Upgrade Schedule |
| `nios_grid_join` | Joins a member to an Infoblox Grid |  |

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
| `nios_notification_syslog_endpoint` | Manages Syslog Endpoints | Retrieves information about existing Syslog Endpoints |

### PARENTAL CONTROL

| Name                              | Resource Description                | Data Source Description                                          |
|-----------------------------------|-------------------------------------|------------------------------------------------------------------|
| `nios_parentalcontrol_subscriberrecord` | Manages Subscriber Records | Retrieves information about existing Subscriber Records |
| `nios_parentalcontrol_subscribersite` | Manages Subscriber Sites | Retrieves information about existing Subscriber Sites |
| `nios_parentalcontrol_blockingpolicy` | Manages Parental Control Blocking Policies | Retrieves information about existing Parental Control Blocking Policies |
| `nios_parentalcontrol_avp` | Manages Parental Control AVPs | Retrieves information about existing Parental Control AVPs |

### RIR

| Name                              | Resource Description                | Data Source Description                                          |
|-----------------------------------|-------------------------------------|------------------------------------------------------------------|
| `nios_rir_organization` | Manages RIR Organizations | Retrieves information about existing RIR Organizations |
