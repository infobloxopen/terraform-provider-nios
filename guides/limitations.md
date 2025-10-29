##  Terraform limitations

Below are the limitations:

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
