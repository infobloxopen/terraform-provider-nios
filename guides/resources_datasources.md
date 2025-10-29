The tables below list all available resources and data sources

### DHCP

| Name | Resource Description                         | Data Source Description                                        |
|----------|----------------------------------------------|----------------------------------------------------------------|
| `nios_dhcp_fixed_address` | Manages DHCP fixed address (IPv4) resources  | Retrieves information about existing DHCP fixed addresses      |
| `nios_dhcp_range` | Manages DHCP range (IPv4) resources          | Retrieves information about existing DHCP ranges               |
| `nios_dhcp_range_template` | Manages DHCP range template (IPv4) resources | Retrieves information about existing DHCP range templates      |
| `nios_dhcp_shared_network` | Manages DHCP shared network (IPv4) resources | Retrieves information about existing DHCP shared networks      |
| `nios_dhcp_ipv6_range_template` | Manages DHCP IPV6 Range Template resources   | Retrieves information about existing DHCP IPV6 Range Templates |                           |                                              |                                                           |

### DNS

| Name                                 | Resource Description                      | Data Source Description                                              |
|--------------------------------------|-------------------------------------------|----------------------------------------------------------------------|
| `nios_dns_view`                      | Manages DNS views                         | Retrieves information about existing DNS views                       |
| `nios_dns_zone_auth`                 | Manages authoritative DNS zones           | Retrieves information about existing authoritative DNS zones         |
| `nios_dns_zone_delegated`            | Manages delegated DNS zones               | Retrieves information about existing delegated DNS zones             |
| `nios_dns_zone_forward`              | Manages forwarding DNS zones              | Retrieves information about existing forwarding DNS zones            |
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


### DTC

| Name | Resource Description | Data Source Description |
|----------|-------------|------------|
| `nios_dtc_lbdn` | Manages DTC LBDN resources | Retrieves information about existing DTC LBDNs |
| `nios_dtc_pool` | Manages DTC pool resources | Retrieves information about existing DTC pools |
| `nios_dtc_server` | Manages DTC server resources | Retrieves information about existing DTC servers |

### IPAM

| Name                              | Resource Description                 | Data Source Description                                           |
|-----------------------------------|--------------------------------------|-------------------------------------------------------------------|
| `nios_ipam_network_view`          | Manages IPAM network views           | Retrieves information about existing IPAM network views           |
| `nios_ipam_network`               | Manages IPAM networks                | Retrieves information about existing IPAM networks                |
| `nios_ipam_network_container`     | Manages IPAM network containers      | Retrieves information about existing IPAM network containers      |
| `nios_ipam_ipv6network`           | Manages IPAM IPv6 networks           | Retrieves information about existing IPAM IPv6 networks           |
| `nios_ipam_ipv6network_container` | Manages IPAM IPv6 network containers | Retrieves information about existing IPAM IPv6 network containers |
| `nios_ipam_bulk_hostname_template` | Manages IPAM Bulk Hostname Templates | Retrieves information about existing IPAM Bulk Hostname templates |

### CLOUD

| Name                                | Resource Description                 | Data Source Description                                           |
|-------------------------------------|--------------------------------------|-------------------------------------------------------------------|
| `nios_cloud_aws_route53_task_group` | Manages AWS Users                | Retrieves information about existing AWS Users                |
| `nios_cloud_aws_user`               | Manages AWS Route 53 Task Groups | Retrieves information about existing AWS Route 53 Task Groups |

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

### Misc

| Name                    | Resource Description  | Data Source Description                            |
|-------------------------|-----------------------|----------------------------------------------------|
| `nios_misc_ruleset`     | Manages Rule Sets     | Retrieves information about existing Rule Sets     |
| `nios_misc_bfdtemplate` | Manages BFD Templates | Retrieves information about existing DFD Templates |

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
