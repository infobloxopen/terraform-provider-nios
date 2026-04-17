# Changelog

## Version 2.0.0

### Newly Supported Resources and Datasources

#### DHCP

- `nios_dhcp_failover` : Manage DHCP failover and retrieve existing failover configurations. ([#287](https://github.com/infobloxopen/terraform-provider-nios/pull/287))
- `nios_dhcp_filteroption` : Manage DHCP filter options and retrieve existing filter option configurations. ([#338](https://github.com/infobloxopen/terraform-provider-nios/pull/338))
- `nios_dhcp_filterrelayagent` : Manage DHCP filter relay agent and retrieve existing relay agent configurations. ([#320](https://github.com/infobloxopen/terraform-provider-nios/pull/320))
- `nios_dhcp_filtermac` : Manage DHCP filter MAC and retrieve existing MAC filter configurations. ([#267](https://github.com/infobloxopen/terraform-provider-nios/pull/267))
- `nios_dhcp_filternac` : Manage DHCP filter NAC and retrieve existing NAC filter configurations. ([#308](https://github.com/infobloxopen/terraform-provider-nios/pull/308))
- `nios_dhcp_fingerprint` : Manage DHCP fingerprints and retrieve existing fingerprint data. ([#311](https://github.com/infobloxopen/terraform-provider-nios/pull/311))
- `nios_dhcp_filterfingerprint` : Manage DHCP filter fingerprints and retrieve existing filter fingerprint configurations. ([#311](https://github.com/infobloxopen/terraform-provider-nios/pull/311))
- `nios_dhcp_fixedaddresstemplate` : Manage DHCP fixed address templates and retrieve existing fixed address template data. ([#261](https://github.com/infobloxopen/terraform-provider-nios/pull/261))
- `nios_dhcp_ipv6filteroption` : Manage DHCP IPv6 filter options and retrieve existing IPv6 filter option configurations. ([#330](https://github.com/infobloxopen/terraform-provider-nios/pull/330))
- `nios_dhcp_ipv6fixedaddress` : Manage DHCP IPv6 fixed addresses and retrieve existing IPv6 fixed address data. ([#277](https://github.com/infobloxopen/terraform-provider-nios/pull/277))
- `nios_dhcp_ipv6range` : Manage DHCP IPv6 ranges and retrieve existing IPv6 range data. ([#273](https://github.com/infobloxopen/terraform-provider-nios/pull/273))
- `nios_dhcp_ipv6sharednetwork` : Manage DHCP IPv6 shared networks and retrieve existing IPv6 shared network data. ([#282](https://github.com/infobloxopen/terraform-provider-nios/pull/282))
- `nios_dhcp_macfilteraddress` : Manage DHCP MAC filter addresses and retrieve existing MAC filter address data. ([#309](https://github.com/infobloxopen/terraform-provider-nios/pull/309))
- `nios_dhcp_optiondefinition` : Manage DHCP option definitions and retrieve existing option definition data. ([#252](https://github.com/infobloxopen/terraform-provider-nios/pull/252))
- `nios_dhcp_optionspace` : Manage DHCP option spaces and retrieve existing option space data. ([#252](https://github.com/infobloxopen/terraform-provider-nios/pull/252))
- `nios_dhcp_roaminghost` : Manage DHCP roaming hosts and retrieve existing roaming host configurations. ([#297](https://github.com/infobloxopen/terraform-provider-nios/pull/297))

#### DNS

- `nios_dns_record_https` : Manage DNS HTTPS records and retrieve existing HTTPS record data. ([#421](https://github.com/infobloxopen/terraform-provider-nios/pull/421))
- `nios_dns_record_svcb` : Manage DNS SVCB records and retrieve existing SVCB record data. ([#419](https://github.com/infobloxopen/terraform-provider-nios/pull/419))
- `nios_dns_sharedrecord_a` : Manage DNS Shared A records and retrieve existing Shared A record data. ([#259](https://github.com/infobloxopen/terraform-provider-nios/pull/259))
- `nios_dns_sharedrecord_aaaa` : Manage DNS Shared AAAA records and retrieve existing Shared AAAA record data. ([#269](https://github.com/infobloxopen/terraform-provider-nios/pull/269))
- `nios_dns_sharedrecord_cname` : Manage DNS Shared CNAME records and retrieve existing Shared CNAME record data. ([#263](https://github.com/infobloxopen/terraform-provider-nios/pull/263))
- `nios_dns_sharedrecord_mx` : Manage DNS Shared MX records and retrieve existing Shared MX record data. ([#246](https://github.com/infobloxopen/terraform-provider-nios/pull/246))
- `nios_dns_sharedrecord_srv` : Manage DNS Shared SRV records and retrieve existing Shared SRV record data. ([#254](https://github.com/infobloxopen/terraform-provider-nios/pull/254))

#### DTC

- `nios_dtc_monitor_http` : Manage DTC HTTP monitors and retrieve existing HTTP monitor configurations. ([#268](https://github.com/infobloxopen/terraform-provider-nios/pull/268))
- `nios_dtc_monitor_icmp` : Manage DTC ICMP monitors and retrieve existing ICMP monitor configurations. ([#274](https://github.com/infobloxopen/terraform-provider-nios/pull/274))
- `nios_dtc_monitor_pdp` : Manage DTC PDP monitors and retrieve existing PDP monitor configurations. ([#283](https://github.com/infobloxopen/terraform-provider-nios/pull/283))
- `nios_dtc_monitor_sip` : Manage DTC SIP monitors and retrieve existing SIP monitor configurations. ([#296](https://github.com/infobloxopen/terraform-provider-nios/pull/296))
- `nios_dtc_monitor_snmp` : Manage DTC SNMP monitors and retrieve existing SNMP monitor configurations. ([#251](https://github.com/infobloxopen/terraform-provider-nios/pull/251))
- `nios_dtc_monitor_tcp` : Manage DTC TCP monitors and retrieve existing TCP monitor configurations. ([#275](https://github.com/infobloxopen/terraform-provider-nios/pull/275))
- `nios_dtc_record_a` : Manage DTC A records and retrieve existing DTC A record data. ([#307](https://github.com/infobloxopen/terraform-provider-nios/pull/307))
- `nios_dtc_record_aaaa` : Manage DTC AAAA records and retrieve existing DTC AAAA record data. ([#315](https://github.com/infobloxopen/terraform-provider-nios/pull/315))
- `nios_dtc_record_cname` : Manage DTC CNAME records and retrieve existing DTC CNAME record data. ([#314](https://github.com/infobloxopen/terraform-provider-nios/pull/314))
- `nios_dtc_record_naptr` : Manage DTC NAPTR records and retrieve existing DTC NAPTR record data. ([#319](https://github.com/infobloxopen/terraform-provider-nios/pull/319))
- `nios_dtc_record_srv` : Manage DTC SRV records and retrieve existing DTC SRV record data. ([#310](https://github.com/infobloxopen/terraform-provider-nios/pull/310))
- `nios_dtc_topology` : Manage DTC topology and retrieve existing topology configurations. ([#245](https://github.com/infobloxopen/terraform-provider-nios/pull/245))
- `nios_dtc_topology_rule` : Retrieve existing DTC topology rule configurations. ([#331](https://github.com/infobloxopen/terraform-provider-nios/pull/331))

#### RPZ

- `nios_dns_record_rpz_a` : Manage RPZ A records and retrieve existing RPZ A record data. ([#256](https://github.com/infobloxopen/terraform-provider-nios/pull/256))
- `nios_dns_record_rpz_a_ipaddress` : Manage RPZ A IP address records and retrieve existing RPZ A IP address data. ([#327](https://github.com/infobloxopen/terraform-provider-nios/pull/327))
- `nios_dns_record_rpz_aaaa` : Manage RPZ AAAA records and retrieve existing RPZ AAAA record data. ([#305](https://github.com/infobloxopen/terraform-provider-nios/pull/305))
- `nios_dns_record_rpz_aaaa_ipaddress` : Manage RPZ AAAA IP address records and retrieve existing RPZ AAAA IP address data. ([#326](https://github.com/infobloxopen/terraform-provider-nios/pull/326))
- `nios_dns_record_rpz_cname` : Manage RPZ CNAME records and retrieve existing RPZ CNAME record data. ([#321](https://github.com/infobloxopen/terraform-provider-nios/pull/321))
- `nios_dns_record_rpz_cname_clientipaddress` : Manage RPZ CNAME client IP address records and retrieve existing RPZ CNAME client IP address data. ([#342](https://github.com/infobloxopen/terraform-provider-nios/pull/342))
- `nios_dns_record_rpz_cname_clientipaddressdn` : Manage RPZ CNAME client IP address DN records and retrieve existing RPZ CNAME client IP address DN data. ([#334](https://github.com/infobloxopen/terraform-provider-nios/pull/334))
- `nios_dns_record_rpz_cname_ipaddress` : Manage RPZ CNAME IP address records and retrieve existing RPZ CNAME IP address data. ([#289](https://github.com/infobloxopen/terraform-provider-nios/pull/289))
- `nios_dns_record_rpz_cname_ipaddressdn` : Manage RPZ CNAME IP address DN records and retrieve existing RPZ CNAME IP address DN data. ([#334](https://github.com/infobloxopen/terraform-provider-nios/pull/334))
- `nios_dns_record_rpz_mx` : Manage RPZ MX records and retrieve existing RPZ MX record data. ([#276](https://github.com/infobloxopen/terraform-provider-nios/pull/276))
- `nios_dns_record_rpz_naptr` : Manage RPZ NAPTR records and retrieve existing RPZ NAPTR record data. ([#280](https://github.com/infobloxopen/terraform-provider-nios/pull/280))
- `nios_dns_record_rpz_ptr` : Manage RPZ PTR records and retrieve existing RPZ PTR record data. ([#278](https://github.com/infobloxopen/terraform-provider-nios/pull/278))
- `nios_dns_record_rpz_srv` : Manage RPZ SRV records and retrieve existing RPZ SRV record data. ([#295](https://github.com/infobloxopen/terraform-provider-nios/pull/295))
- `nios_dns_record_rpz_txt` : Manage RPZ TXT records and retrieve existing RPZ TXT record data. ([#300](https://github.com/infobloxopen/terraform-provider-nios/pull/300))

#### Grid

- `nios_grid_member` : Manage grid members and retrieve existing member configurations. ([#425](https://github.com/infobloxopen/terraform-provider-nios/pull/425))
- `nios_grid_upgradeschedule` : Manage grid upgrade schedules and retrieve existing upgrade schedule configurations. ([#391](https://github.com/infobloxopen/terraform-provider-nios/pull/391))
- `nios_grid_join` : Join a member to an Infoblox Grid. ([#434](https://github.com/infobloxopen/terraform-provider-nios/pull/434))

#### IPAM

- `nios_ipam_ipv6networktemplate` : Manage IPAM IPv6 network templates and retrieve existing IPv6 network template configurations. ([#339](https://github.com/infobloxopen/terraform-provider-nios/pull/339))
- `nios_ipam_networktemplate` : Manage IPAM network templates and retrieve existing network template data. ([#288](https://github.com/infobloxopen/terraform-provider-nios/pull/288))
- `nios_ipam_superhost` : Manage IPAM super hosts and retrieve existing super host data. ([#323](https://github.com/infobloxopen/terraform-provider-nios/pull/323))
- `nios_ipam_vlan` : Manage IPAM VLANs and retrieve existing VLAN data. ([#255](https://github.com/infobloxopen/terraform-provider-nios/pull/255))
- `nios_ipam_vlanrange` : Manage IPAM VLAN ranges and retrieve existing VLAN range data. ([#272](https://github.com/infobloxopen/terraform-provider-nios/pull/272))
- `nios_ipam_vlanview` : Manage IPAM VLAN views and retrieve existing VLAN view data. ([#250](https://github.com/infobloxopen/terraform-provider-nios/pull/250))

#### Microsoft

- `nios_microsoft_mssuperscope` : Manage Microsoft super scopes and retrieve existing super scope configurations. ([#401](https://github.com/infobloxopen/terraform-provider-nios/pull/401))
- `nios_microsoft_msserver_adsites_site` : Manage Microsoft AD sites and retrieve existing AD site data. ([#402](https://github.com/infobloxopen/terraform-provider-nios/pull/402))
- `nios_microsoft_msserver` : Manage Microsoft servers and retrieve existing Microsoft server configurations. ([#395](https://github.com/infobloxopen/terraform-provider-nios/pull/395))

#### Parental Control

- `nios_parentalcontrol_subscriberrecord` : Manage subscriber records and retrieve existing subscriber record data. ([#407](https://github.com/infobloxopen/terraform-provider-nios/pull/407))
- `nios_parentalcontrol_subscribersite` : Manage subscriber sites and retrieve existing subscriber site configurations. ([#408](https://github.com/infobloxopen/terraform-provider-nios/pull/408))
- `nios_parentalcontrol_blockingpolicy` : Manage blocking policies and retrieve existing blocking policy data. ([#392](https://github.com/infobloxopen/terraform-provider-nios/pull/392))
- `nios_parentalcontrol_avp` : Manage parental control AVP and retrieve existing AVP configurations. ([#385](https://github.com/infobloxopen/terraform-provider-nios/pull/385))

#### RIR

- `nios_rir_organization` : Manage RIR organizations and retrieve existing RIR organization data. ([#378](https://github.com/infobloxopen/terraform-provider-nios/pull/378))

#### Security

- `nios_security_radius_authservice` : Manage RADIUS authentication services and retrieve existing RADIUS auth service configurations. ([#405](https://github.com/infobloxopen/terraform-provider-nios/pull/405))
- `nios_security_tacacsplus_authservice` : Manage TACACS+ authentication services and retrieve existing TACACS+ auth service data. ([#406](https://github.com/infobloxopen/terraform-provider-nios/pull/406))
- `nios_security_ldap_authservice` : Manage LDAP authentication services and retrieve existing LDAP auth service configurations. ([#396](https://github.com/infobloxopen/terraform-provider-nios/pull/396))
- `nios_security_saml_authservice` : Manage SAML authentication services and retrieve existing SAML auth service data. ([#393](https://github.com/infobloxopen/terraform-provider-nios/pull/393))

#### Miscellaneous

- `nios_misc_tftpfiledir` : Manage TFTP file directories and retrieve existing TFTP file directory configurations. ([#411](https://github.com/infobloxopen/terraform-provider-nios/pull/411))
- `nios_misc_dxl_endpoint` : Manage DXL endpoints and retrieve existing DXL endpoint data. ([#389](https://github.com/infobloxopen/terraform-provider-nios/pull/389))
- `nios_misc_syslog_endpoint` : Manage syslog notification endpoints and retrieve existing syslog endpoint configurations. ([#438](https://github.com/infobloxopen/terraform-provider-nios/pull/438))


### Modules

- **NIOS Grid Member AWS Module** : Terraform module for deploying NIOS Grid EC2 instances on AWS. ([#434](https://github.com/infobloxopen/terraform-provider-nios/pull/434))

### Enhancements

- Added support for NIOS 9.1.0 with WAPI version v2.14 and UUID-based API calls.
- Updated BFD Template by removing unsupported authentication fields (`authentication_key`, `authentication_key_id`, and `authentication_type`) to maintain parity with WAPI v2.14 ([#418](https://github.com/infobloxopen/terraform-provider-nios/pull/418))
- Enhanced CNAME Record with `rr_precondition_instructions` field enabling conditional record creation. When A or AAAA records with the same name exist, this field allows automatic deletion of conflicting records based on specified conditions before creating the CNAME record. ([#418](https://github.com/infobloxopen/terraform-provider-nios/pull/418))
- Enhanced Alias Record with CAA enum support under `target_type`. ([#418](https://github.com/infobloxopen/terraform-provider-nios/pull/418))
- Updated DTC Topology with structured `destination` field replacing `destination_link`. ([#418](https://github.com/infobloxopen/terraform-provider-nios/pull/418))
- Updated Admin Group field naming: renamed `set_analytics_database_dump` to `set_query_logging_warnings` in nested command fields. ([#418](https://github.com/infobloxopen/terraform-provider-nios/pull/418))
- Added comprehensive IPv4, IPv6, CIDR, and domain name validation support across multiple resource types.([#345](https://github.com/infobloxopen/terraform-provider-nios/pull/345))
- Improved handling of unknown values in ValidateConfig to support variable resolution. ([#375](https://github.com/infobloxopen/terraform-provider-nios/pull/375))

### Fixes

- IPAM: Fixed VLAN type validation error in Network and IPv6 Network Container objects when reading network objects containing VLAN information. ([#380](https://github.com/infobloxopen/terraform-provider-nios/pull/380))
- Fixed numerous bugs across DHCP, DNS, DTC, IPAM, Grid, Security, and other modules improving overall stability and reliability.

## Version 1.1.0

### Newly Supported Resources and Datasources

#### DHCP

- `nios_dhcp_ipv6dhcpoptiondefinition` : Manage DHCP IPv6 option definition and retrieve existing IPv6 option definitions. ([#242](https://github.com/infobloxopen/terraform-provider-nios/pull/242))
- `nios_dhcp_ipv6dhcpoptionspace` : Manage DHCP IPv6 option space and retrieve existing IPv6 option spaces. ([#242](https://github.com/infobloxopen/terraform-provider-nios/pull/242))
- `nios_dhcp_ipv6fixedaddresstemplate` : Manage DHCP IPv6 fixed address template and retrieve existing IPv6 fixed address templates. ([#247](https://github.com/infobloxopen/terraform-provider-nios/pull/247))

#### DNS

- `nios_dns_sharedrecordgroup` : Manage Shared Record Group and retrieve existing Shared Record Groups. ([#244](https://github.com/infobloxopen/terraform-provider-nios/pull/244))
- `nios_dns_sharedrecord_txt` : Manage Shared TXT Record and retrieve existing Shared TXT Records. ([#248](https://github.com/infobloxopen/terraform-provider-nios/pull/248))

### Enhancements

- Enhanced import workflow by eliminating Update API calls during terraform plan ensuring updates occur only after apply. ([#241](https://github.com/infobloxopen/terraform-provider-nios/pull/241))

### Fixes

- DNS/DTC: Fixed inconsistent TTL behavior during modification for record PTR and DTC Pool. ([#253](https://github.com/infobloxopen/terraform-provider-nios/pull/253))
- Fixed `ExtractResourceRef` to gracefully handle malformed refs lacking /, eliminating index-out-of-bounds crashes. ([#284](https://github.com/infobloxopen/terraform-provider-nios/pull/284))
- DISCOVERY/IPAM: Added validations to ensure valid_lifetime matches dhcp-lease-time option value for `nios_ipam_ipv6network` and `nios_ipam_ipv6network_container` and added port validations for `nios_discovery_vdiscovery_task`. ([#290](https://github.com/infobloxopen/terraform-provider-nios/pull/290))

## Version 1.0.0

### Newly Supported Resources and Datasources

#### ACL

- `nios_acl_namedacl` : Manage named ACLs and retrieve existing named ACL configurations. ([#111](https://github.com/infobloxopen/terraform-provider-nios/pull/111))

#### Cloud

- `nios_cloud_aws_route53_task_group` : Manage AWS Route53 task groups and retrieve existing task group configurations. ([#112](https://github.com/infobloxopen/terraform-provider-nios/pull/112))
- `nios_cloud_aws_user` : Manage AWS users and retrieve existing AWS user data. ([#105](https://github.com/infobloxopen/terraform-provider-nios/pull/105))

#### Discovery

- `nios_discovery_credentialgroup` : Manage discovery credential groups and retrieve existing credential group data. ([#144](https://github.com/infobloxopen/terraform-provider-nios/pull/144))
- `nios_discovery_vdiscovery_task` : Manage vDiscovery tasks and retrieve existing vDiscovery task configurations. ([#169](https://github.com/infobloxopen/terraform-provider-nios/pull/169))

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
- `nios_ip_allocation` : Manage allocation and deallocation of an IP address from a network. ([#143](https://github.com/infobloxopen/terraform-provider-nios/pull/143))
- `nios_ip_association` : Manage association and disassociation of an IP address with a VM. ([#143](https://github.com/infobloxopen/terraform-provider-nios/pull/143))
- `nios_record_host` : Retrieves existing host record data. ([#143](https://github.com/infobloxopen/terraform-provider-nios/pull/143))

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

### Supported Resources and Datasources

#### DNS

- `nios_dns_view` : Manage DNS views and retrieve existing view configurations. ([#67](https://github.com/infobloxopen/terraform-provider-nios/pull/67))
- `nios_dns_zone_auth` : Manage authoritative DNS zones and retrieve existing zone data. ([#57](https://github.com/infobloxopen/terraform-provider-nios/pull/57))
- `nios_dns_zone_delegated` : Manage delegated DNS zones and retrieve existing delegation data. ([#62](https://github.com/infobloxopen/terraform-provider-nios/pull/62))
- `nios_dns_zone_forward` : Manage forwarding DNS zones and retrieve existing forward zone data. ([#33](https://github.com/infobloxopen/terraform-provider-nios/pull/33))
- `nios_dns_record_a` : Manage DNS A records and retrieve existing A record data. ([#1](https://github.com/infobloxopen/terraform-provider-nios/pull/1))
- `nios_dns_record_aaaa` : Manage DNS AAAA records and retrieve existing AAAA record data. ([#23](https://github.com/infobloxopen/terraform-provider-nios/pull/23))
- `nios_dns_record_alias` : Manage DNS ALIAS records and retrieve existing ALIAS record data. ([#46](https://github.com/infobloxopen/terraform-provider-nios/pull/46))
- `nios_dns_record_cname` : Manage DNS CNAME records and retrieve existing CNAME record data. ([#50](https://github.com/infobloxopen/terraform-provider-nios/pull/50))
- `nios_dns_record_mx` : Manage DNS MX records and retrieve existing MX record data. ([#61](https://github.com/infobloxopen/terraform-provider-nios/pull/61))
- `nios_dns_record_ns` : Manage DNS NS records and retrieve existing NS record data. ([#55](https://github.com/infobloxopen/terraform-provider-nios/pull/55))
- `nios_dns_record_ptr` : Manage DNS PTR records and retrieve existing PTR record data. ([#45](https://github.com/infobloxopen/terraform-provider-nios/pull/45))
- `nios_dns_record_srv` : Manage DNS SRV records and retrieve existing SRV record data. ([#53](https://github.com/infobloxopen/terraform-provider-nios/pull/53))
- `nios_dns_record_txt` : Manage DNS TXT records and retrieve existing TXT record data. ([#53](https://github.com/infobloxopen/terraform-provider-nios/pull/53))

#### DHCP

- `nios_dhcp_fixed_address` : Manage DHCP fixed address resources and retrieve existing fixed address data. ([#51](https://github.com/infobloxopen/terraform-provider-nios/pull/51))
- `nios_dhcp_range` : Manage DHCP range resources and retrieve existing DHCP range data. ([#63](https://github.com/infobloxopen/terraform-provider-nios/pull/63))
- `nios_dhcp_range_template` : Manage DHCP range templates and retrieve existing template data. ([#52](https://github.com/infobloxopen/terraform-provider-nios/pull/52))
- `nios_dhcp_shared_network` : Manage DHCP shared networks and retrieve existing shared network data. ([#64](https://github.com/infobloxopen/terraform-provider-nios/pull/64))

#### IPAM

- `nios_ipam_network_view` : Manage IPAM network views and retrieve existing view data. ([#68](https://github.com/infobloxopen/terraform-provider-nios/pull/68))
- `nios_ipam_network` : Manage IPAM networks and retrieve existing network data. ([#44](https://github.com/infobloxopen/terraform-provider-nios/pull/44))
- `nios_ipam_network_container` : Manage IPAM network containers and retrieve existing container data. ([#38](https://github.com/infobloxopen/terraform-provider-nios/pull/38))
- `nios_ipam_ipv6network` : Manage IPAM IPv6 networks and retrieve existing IPv6 network data. ([#69](https://github.com/infobloxopen/terraform-provider-nios/pull/69))
- `nios_ipam_ipv6network_container` : Manage IPAM IPv6 network containers and retrieve existing IPv6 container data. ([#56](https://github.com/infobloxopen/terraform-provider-nios/pull/56))

#### DTC

- `nios_dtc_lbdn` : Manage DTC LBDN resources and retrieve existing LBDN configurations. ([#27](https://github.com/infobloxopen/terraform-provider-nios/pull/27))
- `nios_dtc_pool` : Manage DTC pools and retrieve existing pool data. ([#28](https://github.com/infobloxopen/terraform-provider-nios/pull/28))
- `nios_dtc_server` : Manage DTC servers and retrieve existing server data. ([#32](https://github.com/infobloxopen/terraform-provider-nios/pull/32))
