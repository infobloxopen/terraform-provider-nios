locals {
  // Build matching rule from the explicitly provided instance OCIDs.
  instance_rules = [for id in var.instance_ocids : "instance.id = '${id}'"]

  matching_rule = "Any {${join(", ", local.instance_rules)}}"
}

// Manage a Dynamic Resource Group in the specified Identity Domain
resource "oci_identity_domains_dynamic_resource_group" "nios_ha" {
  idcs_endpoint = var.idcs_endpoint
  display_name  = var.dynamic_group_name
  description   = "NIOS HA instances for VIP failover"
  matching_rule = local.matching_rule
  schemas       = ["urn:ietf:params:scim:schemas:oracle:idcs:DynamicResourceGroup"]

  lifecycle {
    ignore_changes = [schemas]
  }
}


// Manages a policy at the sub compartment level granting necessary permissions for NIOS HA
resource "oci_identity_policy" "nios_ha_compartment" {
  compartment_id = var.compartment_id
  name           = var.policy_name
  description    = "Policy allowing NIOS HA instances to manage VIP failover"

  statements = [
    "Allow dynamic-group '${var.identity_domain_name}'/'${var.dynamic_group_name}' to read vnics in compartment id ${var.compartment_id}",
    "Allow dynamic-group '${var.identity_domain_name}'/'${var.dynamic_group_name}' to read vnic-attachments in compartment id ${var.compartment_id}",
    "Allow dynamic-group '${var.identity_domain_name}'/'${var.dynamic_group_name}' to read instances in compartment id ${var.compartment_id}",
    "Allow dynamic-group '${var.identity_domain_name}'/'${var.dynamic_group_name}' to manage private-ips in compartment id ${var.compartment_id}",
    "Allow dynamic-group '${var.identity_domain_name}'/'${var.dynamic_group_name}' to read domains in compartment id ${var.compartment_id}",
    "Allow dynamic-group '${var.identity_domain_name}'/'${var.dynamic_group_name}' to read dynamic-groups in compartment id ${var.compartment_id}",
    "Allow dynamic-group '${var.identity_domain_name}'/'${var.dynamic_group_name}' to read policies in compartment id ${var.compartment_id}",
    "Allow dynamic-group '${var.identity_domain_name}'/'${var.dynamic_group_name}' to inspect drgs in compartment id ${var.compartment_id}",
    "Allow dynamic-group '${var.identity_domain_name}'/'${var.dynamic_group_name}' to inspect dynamic-groups in compartment id ${var.compartment_id}",
    "Allow dynamic-group '${var.identity_domain_name}'/'${var.dynamic_group_name}' to use virtual-network-family in compartment id ${var.compartment_id}",
    "Allow dynamic-group '${var.identity_domain_name}'/'${var.dynamic_group_name}' to read virtual-network-family in compartment id ${var.compartment_id}",
  ]
  freeform_tags = var.freeform_tags
  depends_on    = [oci_identity_domains_dynamic_resource_group.nios_ha]
}

// Manages Policy at the tenancy level
resource "oci_identity_policy" "nios_ha_identity" {
  compartment_id = var.tenancy_ocid
  name           = "${var.policy_name}-identity"
  description    = "Policy allowing NIOS HA instances to read identity resources for dynamic group verification"

  statements = [
    "Allow dynamic-group '${var.identity_domain_name}'/'${var.dynamic_group_name}' to read domains in tenancy",
    "Allow dynamic-group '${var.identity_domain_name}'/'${var.dynamic_group_name}' to read dynamic-groups in tenancy",
    "Allow dynamic-group '${var.identity_domain_name}'/'${var.dynamic_group_name}' to read policies in tenancy",
  ]

  freeform_tags = var.freeform_tags
  depends_on    = [oci_identity_domains_dynamic_resource_group.nios_ha]
}
