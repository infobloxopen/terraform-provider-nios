# Changelog

## Version 1.0.0

### Newly Supported Resources and Datasources

#### ACL

- `nios_acl_namedacl` : Manage named ACLs and retrieve existing named ACL configurations. ([#111](https://github.com/infobloxopen/terraform-provider-nios/pull/111))

#### Cloud

- `nios_cloud_aws_route53_task_group` : Manage AWS Route53 task groups and retrieve existing task group configurations. ([#112](https://github.com/infobloxopen/terraform-provider-nios/pull/112))
- `nios_cloud_aws_user` : Manage AWS users and retrieve existing AWS user data. ([#105](https://github.com/infobloxopen/terraform-provider-nios/pull/105))

#### Discovery

- `nios_discovery_credentialgroup` : Manage discovery credential groups and retrieve existing credential group data. ([#144](https://github.com/infobloxopen/terraform-provider-nios/pull/144))
- `nios_discovery_vdiscovery_task` : Manage vdiscovery tasks and retrieve existing vdiscovery task configurations. ([#169](https://github.com/infobloxopen/terraform-provider-nios/pull/169))

#### DNS

- `nios_dns_nsgroup` : Manage name server groups and retrieve existing NS group data. ([#95](https://github.com/infobloxopen/terraform-provider-nios/pull/95))
- `nios_dns_nsgroup_delegation` : Manage name server group delegations and retrieve existing delegation data. ([#97](https://github.com/infobloxopen/terraform-provider-nios/pull/97))
- `nios_dns_nsgroup_forwardingmember` : Manage forwarding members in name server groups and retrieve existing forwarding member data. ([#122](https://github.com/infobloxopen/terraform-provider-nios/pull/122))
- `nios_dns_nsgroup_forwardstubserver` : Manage forward stub servers in name server groups and retrieve existing stub server data. ([#136](https://github.com/infobloxopen/terraform-provider-nios/pull/136))
- `nios_dns_nsgroup_stubmember` : Manage stub members in name server groups and retrieve existing stub member data. ([#130](https://github.com/infobloxopen/terraform-provider-nios/pull/130))
- `nios_dns_record_dname` : Manage DNS DNAME records and retrieve existing DNAME record data. ([#101](https://github.com/infobloxopen/terraform-provider-nios/pull/101))
- `nios_dns_record_naptr` : Manage DNS NAPTR records and retrieve existing NAPTR record data. ([#102](https://github.com/infobloxopen/terraform-provider-nios/pull/102))
- `nios_dns_record_tlsa` : Manage DNS TLSA records and retrieve existing TLSA record data. ([#110](https://github.com/infobloxopen/terraform-provider-nios/pull/110))
- `nios_dns_record_caa` : Manage DNS CAA records and retrieve existing CAA record data. ([#120](https://github.com/infobloxopen/terraform-provider-nios/pull/120))
- `nios_dns_record_unknown` : Manage DNS unknown type records and retrieve existing unknown record data. ([#134](https://github.com/infobloxopen/terraform-provider-nios/pull/134))
- `nios_dns_zone_stub` : Manage DNS stub zones and retrieve existing stub zone data. ([#103](https://github.com/infobloxopen/terraform-provider-nios/pull/103))
- `nios_dns_zone_rp` : Manage DNS reverse proxy zones and retrieve existing RP zone data. ([#164](https://github.com/infobloxopen/terraform-provider-nios/pull/164))
- `nios_ip_allocation` : Manage Allocation and deallocation of an IP address from a network. ([#143](https://github.com/infobloxopen/terraform-provider-nios/pull/143))
- `nios_ip_association` : Manage Association and disassociation of an IP address with a VM. ([#143](https://github.com/infobloxopen/terraform-provider-nios/pull/143))
- `nios_record_host` : Retrieves exisiting host record data. ([#143](https://github.com/infobloxopen/terraform-provider-nios/pull/143))

#### Grid

- `nios_grid_distributionschedule` : Manage grid distribution schedules and retrieve existing schedule data. ([#124](https://github.com/infobloxopen/terraform-provider-nios/pull/124))
- `nios_grid_extensibleattributedef` : Manage extensible attribute definitions and retrieve existing attribute definitions. ([#116](https://github.com/infobloxopen/terraform-provider-nios/pull/116))
- `nios_grid_servicerestart_group` : Manage service restart groups and retrieve existing restart group configurations. ([#149](https://github.com/infobloxopen/terraform-provider-nios/pull/149))
- `nios_grid_natgroup` : Manage NAT groups and retrieve existing NAT group data. ([#117](https://github.com/infobloxopen/terraform-provider-nios/pull/117))
- `nios_grid_upgradegroup` : Manage upgrade groups and retrieve existing upgrade group configurations. ([#123](https://github.com/infobloxopen/terraform-provider-nios/pull/123))

#### IPAM

- `nios_ipam_bulk_hostname_template` : Manage bulk hostname templates and retrieve existing template configurations. ([#96](https://github.com/infobloxopen/terraform-provider-nios/pull/96))

#### Miscellaneous

- `nios_misc_bfdtemplate` : Manage BFD templates and retrieve existing BFD template data. ([#121](https://github.com/infobloxopen/terraform-provider-nios/pull/121))
- `nios_misc_ruleset` : Manage rulesets and retrieve existing ruleset configurations. ([#98](https://github.com/infobloxopen/terraform-provider-nios/pull/98))

#### Notification

- `nios_notification_rest_endpoint` : Manage REST notification endpoints and retrieve existing endpoint data. ([#171](https://github.com/infobloxopen/terraform-provider-nios/pull/171))
- `nios_notification_rule` : Manage notification rules and retrieve existing notification rule configurations. ([#154](https://github.com/infobloxopen/terraform-provider-nios/pull/154))

#### Security

- `nios_security_admin_user` : Manage administrator users and retrieve existing admin user configurations. ([#93](https://github.com/infobloxopen/terraform-provider-nios/pull/93))
- `nios_security_admin_role` : Manage administrator roles and retrieve existing admin role data. ([#94](https://github.com/infobloxopen/terraform-provider-nios/pull/94))
- `nios_security_admin_group` : Manage administrator groups and retrieve existing admin group data. ([#140](https://github.com/infobloxopen/terraform-provider-nios/pull/140))
- `nios_security_permission` : Manage permissions and retrieve existing permission configurations. ([#139](https://github.com/infobloxopen/terraform-provider-nios/pull/139))
- `nios_security_ftpuser` : Manage FTP users and retrieve existing FTP user data. ([#146](https://github.com/infobloxopen/terraform-provider-nios/pull/146))
- `nios_security_certificate_authservice` : Manage certificate authentication services and retrieve existing certificate data. ([#153](https://github.com/infobloxopen/terraform-provider-nios/pull/153))
- `nios_security_snmpuser` : Manage SNMP users and retrieve existing SNMP user configurations. ([#137](https://github.com/infobloxopen/terraform-provider-nios/pull/137))

#### SmartFolder

- `nios_smartfolder_global` : Manage global smart folders and retrieve existing global smart folder configurations. ([#119](https://github.com/infobloxopen/terraform-provider-nios/pull/119))
- `nios_smartfolder_personal` : Manage personal smart folders and retrieve existing personal smart folder data. ([#104](https://github.com/infobloxopen/terraform-provider-nios/pull/104))

## Version 0.0.1

### Newly Supported Resources and Datasources

#### DNS

- `nios_dns_view` : Manage DNS views and retrieve existing view configurations. ([#67](https://github.com/infobloxopen/terraform-provider-bloxone/pull/67))
- `nios_dns_zone_auth` : Manage authoritative DNS zones and retrieve existing zone data. ([#57](https://github.com/infobloxopen/terraform-provider-bloxone/pull/57))
- `nios_dns_zone_delegated` : Manage delegated DNS zones and retrieve existing delegation data. ([#62](https://github.com/infobloxopen/terraform-provider-bloxone/pull/62))
- `nios_dns_zone_forward` : Manage forwarding DNS zones and retrieve existing forward zone data. ([#33](https://github.com/infobloxopen/terraform-provider-bloxone/pull/33))
- `nios_dns_record_a` : Manage DNS A records and retrieve existing A record data. ([#1](https://github.com/infobloxopen/terraform-provider-bloxone/pull/1))
- `nios_dns_record_aaaa` : Manage DNS AAAA records and retrieve existing AAAA record data. ([#23](https://github.com/infobloxopen/terraform-provider-bloxone/pull/23))
- `nios_dns_record_alias` : Manage DNS ALIAS records and retrieve existing ALIAS record data. ([#46](https://github.com/infobloxopen/terraform-provider-bloxone/pull/46))
- `nios_dns_record_cname` : Manage DNS CNAME records and retrieve existing CNAME record data. ([#50](https://github.com/infobloxopen/terraform-provider-bloxone/pull/50))
- `nios_dns_record_mx` : Manage DNS MX records and retrieve existing MX record data. ([#61](https://github.com/infobloxopen/terraform-provider-bloxone/pull/61))
- `nios_dns_record_ns` : Manage DNS NS records and retrieve existing NS record data. ([#55](https://github.com/infobloxopen/terraform-provider-bloxone/pull/55))
- `nios_dns_record_ptr` : Manage DNS PTR records and retrieve existing PTR record data. ([#45](https://github.com/infobloxopen/terraform-provider-bloxone/pull/45))
- `nios_dns_record_srv` : Manage DNS SRV records and retrieve existing SRV record data. ([#53](https://github.com/infobloxopen/terraform-provider-bloxone/pull/53))
- `nios_dns_record_txt` : Manage DNS TXT records and retrieve existing TXT record data. ([#53](https://github.com/infobloxopen/terraform-provider-bloxone/pull/53))

#### DHCP

- `nios_dhcp_fixed_address` : Manage DHCP fixed address resources and retrieve existing fixed address data. ([#51](https://github.com/infobloxopen/terraform-provider-bloxone/pull/51))
- `nios_dhcp_range` : Manage DHCP range resources and retrieve existing DHCP range data. ([#63](https://github.com/infobloxopen/terraform-provider-bloxone/pull/63))
- `nios_dhcp_range_template` : Manage DHCP range templates and retrieve existing template data. ([#52](https://github.com/infobloxopen/terraform-provider-bloxone/pull/52))
- `nios_dhcp_shared_network` : Manage DHCP shared networks and retrieve existing shared network data. ([#64](https://github.com/infobloxopen/terraform-provider-bloxone/pull/64))

#### IPAM

- `nios_ipam_network_view` : Manage IPAM network views and retrieve existing view data. ([#68](https://github.com/infobloxopen/terraform-provider-bloxone/pull/68))
- `nios_ipam_network` : Manage IPAM networks and retrieve existing network data. ([#44](https://github.com/infobloxopen/terraform-provider-bloxone/pull/44))
- `nios_ipam_network_container` : Manage IPAM network containers and retrieve existing container data. ([#38](https://github.com/infobloxopen/terraform-provider-bloxone/pull/38))
- `nios_ipam_ipv6network` : Manage IPAM IPv6 networks and retrieve existing IPv6 network data. ([#69](https://github.com/infobloxopen/terraform-provider-bloxone/pull/69))
- `nios_ipam_ipv6network_container` : Manage IPAM IPv6 network containers and retrieve existing IPv6 container data. ([#56](https://github.com/infobloxopen/terraform-provider-bloxone/pull/56))

#### DTC

- `nios_dtc_lbdn` : Manage DTC LBDN resources and retrieve existing LBDN configurations. ([#27](https://github.com/infobloxopen/terraform-provider-bloxone/pull/27))
- `nios_dtc_pool` : Manage DTC pools and retrieve existing pool data. ([#28](https://github.com/infobloxopen/terraform-provider-bloxone/pull/28))
- `nios_dtc_server` : Manage DTC servers and retrieve existing server data. ([#32](https://github.com/infobloxopen/terraform-provider-bloxone/pull/32))
