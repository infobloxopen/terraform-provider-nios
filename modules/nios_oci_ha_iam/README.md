# nios_oci_ha_iam

Terraform module to create the OCI IAM resources required for NIOS High
Availability (HA) Configuration.

## Overview

For vNIOS HA on OCI, both nodes must function as **OCI principals** so they
can make the API calls required during failover — primarily to move the
Virtual IP (VIP) between nodes by reassigning a secondary private IP. This
is achieved by placing the HA instances into a **Dynamic Group** and
attaching IAM policies that grant the dynamic group the required
permissions on networking and identity resources.

This module provisions:

1. A **Dynamic Resource Group** in your Identity Domain whose matching
   rule references the NIOS HA instance OCIDs you pass in.
2. A **sub-compartment-level IAM policy** granting the dynamic group the
   permissions needed to manage network interfaces, assign/unassign
   secondary IPs, and perform the other OCI operations required for
   automated HA failover.
3. A **tenancy-level IAM policy** granting the minimum identity-level
   permissions required for identity validation and authorization checks
   essential for HA workflows.

A single instance of this module can serve **multiple HA pairs** in the
same compartment — add each new pair's instance OCIDs to `instance_ocids`
and re-apply; the matching rule is updated in place.

For the full background and the complete list of policy statements, see:

- [Prerequisites for vNIOS for OCI HA](https://docs.infoblox.com/space/vniosoci/2188214440/Prerequisites+for+vNIOS+for+OCI+HA)
- [Deploying vNIOS for OCI in an HA Environment](https://docs.infoblox.com/space/vniosoci/2178056272/Deploying+vNIOS+for+OCI+in+an+HA+Environment)


<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.12.1 |
| <a name="requirement_oci"></a> [oci](#requirement\_oci) | >= 5.0.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_oci"></a> [oci](#provider\_oci) | >= 5.0.0 |

## Resources

| Name | Type |
|------|------|
| [oci_identity_domains_dynamic_resource_group.nios_ha](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/identity_domains_dynamic_resource_group) | resource |
| [oci_identity_policy.nios_ha_compartment](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/identity_policy) | resource |
| [oci_identity_policy.nios_ha_identity](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/identity_policy) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_compartment_id"></a> [compartment\_id](#input\_compartment\_id) | OCID of the compartment where NIOS instances reside. | `string` | n/a | yes |
| <a name="input_dynamic_group_name"></a> [dynamic\_group\_name](#input\_dynamic\_group\_name) | Name for the dynamic group. | `string` | `"nios-ha-dynamic-group"` | no |
| <a name="input_freeform_tags"></a> [freeform\_tags](#input\_freeform\_tags) | Freeform tags to apply to IAM resources. | `map(string)` | `{}` | no |
| <a name="input_idcs_endpoint"></a> [idcs\_endpoint](#input\_idcs\_endpoint) | IDCS endpoint URL for OCI Identity resources (e.g. https://idcs-<region>.oraclecloud.com). Required for creating dynamic groups. | `string` | n/a | yes |
| <a name="input_identity_domain_name"></a> [identity\_domain\_name](#input\_identity\_domain\_name) | Name of the Identity Domain (used in policy statements). Usually 'Default' or your custom domain name. | `string` | `"Default"` | no |
| <a name="input_instance_ocids"></a> [instance\_ocids](#input\_instance\_ocids) | List of NIOS HA instance OCIDs to include in the dynamic group. Must contain at least one valid compute instance OCID. | `list(string)` | n/a | yes |
| <a name="input_policy_name"></a> [policy\_name](#input\_policy\_name) | Name for the IAM policy. | `string` | `"nios-ha-policy"` | no |
| <a name="input_tenancy_ocid"></a> [tenancy\_ocid](#input\_tenancy\_ocid) | OCID of the tenancy. Required for creating dynamic groups. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_dynamic_group_id"></a> [dynamic\_group\_id](#output\_dynamic\_group\_id) | OCID of the created dynamic group. |
| <a name="output_dynamic_group_name"></a> [dynamic\_group\_name](#output\_dynamic\_group\_name) | Name of the created dynamic group. |
| <a name="output_identity_policy_id"></a> [identity\_policy\_id](#output\_identity\_policy\_id) | OCID of the tenancy-level identity IAM policy. |
| <a name="output_identity_policy_statements"></a> [identity\_policy\_statements](#output\_identity\_policy\_statements) | Statements in the tenancy-level policy. |
| <a name="output_matching_rule"></a> [matching\_rule](#output\_matching\_rule) | The matching rule used for the dynamic group. |
<!-- END_TF_DOCS -->

## Usage

```hcl
module "nios_iam" {
  source               = "../../../modules/nios_oci_ha_iam"
  tenancy_ocid         = var.tenancy_ocid
  compartment_id       = var.compartment_id
  idcs_endpoint        = var.idcs_endpoint
  identity_domain_name = var.identity_domain_name
  dynamic_group_name   = var.dynamic_group_name

  // Pass instance OCIDs for matching rule
  instance_ocids = [
    module.node1.instance_id,
    module.node2.instance_id,
  ]
}
```
