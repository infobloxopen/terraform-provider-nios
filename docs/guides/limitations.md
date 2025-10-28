##  Terraform limitations

Below are the limitations:

- Cannot specify reverse mapping notation for Zone FQDNs (Auth, Forward, Delegated)
- When setting use_ttl=false or removing the field, the provider fails to unset TTL properly. Users must remove both ttl and use_ttl fields to successfully unset TTL.
- Function call next_available_network for IPv4/IPv6 networks fails during subsequent terraform apply operations with "overlap an existing network" error.
- Data Source provides no value for extattrs_all since this is for internal use only. Users should only work with the extattrs field.
- Range templates have cloud_api_compatible set to true by default, as Terraform's internal ID structure requires cloud compatibility. Setting this to false causes errors when adding the Terraform Internal ID extensible attribute.
- IPv6 PTR records lack function call support.
- Forward Zones lack TSIG support
- Cloud platform configuration must be nullified before modification or removal.

Below are the limitations which cannot be fixed:

- Extensible Attribute definitions(extensibleattributedef):
    - Unable to set G flag in EA.
- AWS Route 53 Task Groups(aws_route53_task_group):
    - It is possible to set the ARN value from Terraform for the nios_cloud_aws_route53_task_group object even when syn_child_account is set to false.
- IPv6 Network(ipv6network):
    - Few fields can't be present when template field is present 
- Response Policy Zones(zone_rp):
    - Need to support rpz_type = "FEED" while creating a nios_dns_zone_rp resource.
- Authoritative Zone(zone_auth):
    - Discrepancy with state of dnssec_key_params under Zone auth with each terraform apply.
