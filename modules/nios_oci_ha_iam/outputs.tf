output "dynamic_group_id" {
  description = "OCID of the created dynamic group."
  value       = oci_identity_domains_dynamic_resource_group.nios_ha.id
}

output "dynamic_group_name" {
  description = "Name of the created dynamic group."
  value       = oci_identity_domains_dynamic_resource_group.nios_ha.display_name
}

output "identity_policy_id" {
  description = "OCID of the tenancy-level identity IAM policy."
  value       = oci_identity_policy.nios_ha_identity.id
}

output "identity_policy_statements" {
  description = "Statements in the tenancy-level policy."
  value       = oci_identity_policy.nios_ha_identity.statements
}

output "matching_rule" {
  description = "The matching rule used for the dynamic group."
  value       = local.matching_rule
}
