data "oci_core_vnic" "lan1_vnic" {
  vnic_id = oci_core_vnic_attachment.lan1_vnic_attachment.vnic_id
}

data "oci_core_subnet" "lan1_subnet" {
  subnet_id = var.lan1_subnet_id
}

output "instance_id" {
  description = "OCID of the NIOS compute instance."
  value       = oci_core_instance.nios_instance.id
}

output "lan1_vnic_id" {
  description = "OCID of the secondary (LAN1) VNIC."
  value       = oci_core_vnic_attachment.lan1_vnic_attachment.vnic_id
}

output "lan1_private_ip" {
  description = "Private IP address of the secondary (LAN1) VNIC."
  value       = data.oci_core_vnic.lan1_vnic.private_ip_address
}

output "lan1_subnet_mask" {
  description = "Subnet mask of the LAN1 subnet (e.g. 255.255.255.0)."
  value       = cidrnetmask(data.oci_core_subnet.lan1_subnet.cidr_block)
}

output "lan1_gateway" {
  description = "Gateway IP of the LAN1 subnet (OCI virtual router IP)."
  value       = data.oci_core_subnet.lan1_subnet.virtual_router_ip
}

data "oci_core_subnet" "mgmt_subnet" {
  subnet_id = var.mgmt_subnet_id
}

output "mgmt_ip" {
  description = "Private IP of the primary (MGMT) VNIC."
  value       = oci_core_instance.nios_instance.private_ip
}

output "mgmt_subnet_mask" {
  description = "Subnet mask of the MGMT subnet."
  value       = cidrnetmask(data.oci_core_subnet.mgmt_subnet.cidr_block)
}

output "mgmt_gateway" {
  description = "Gateway IP of the MGMT subnet."
  value       = data.oci_core_subnet.mgmt_subnet.virtual_router_ip
}

data "oci_core_vnic" "ha_vnic" {
  count   = var.enable_ha ? 1 : 0
  vnic_id = oci_core_vnic_attachment.ha_vnic_attachment[0].vnic_id
}

data "oci_core_subnet" "ha_subnet" {
  count     = var.enable_ha ? 1 : 0
  subnet_id = var.ha_subnet_id
}

output "ha_ip" {
  description = "Private IP of the HA VNIC."
  value       = var.enable_ha ? data.oci_core_vnic.ha_vnic[0].private_ip_address : null
}

output "ha_subnet_mask" {
  description = "Subnet mask of the HA subnet."
  value       = var.enable_ha ? cidrnetmask(data.oci_core_subnet.ha_subnet[0].cidr_block) : null
}

output "ha_gateway" {
  description = "Gateway IP of the HA subnet."
  value       = var.enable_ha ? data.oci_core_subnet.ha_subnet[0].virtual_router_ip : null
}

output "ha_vnic_id" {
  description = "OCID of the HA VNIC."
  value       = var.enable_ha ? oci_core_vnic_attachment.ha_vnic_attachment[0].vnic_id : null
}

output "vip" {
  description = "Virtual IP (VIP) for HA - floats between primary/secondary."
  value       = var.enable_ha && var.is_primary ? oci_core_private_ip.ha_vip[0].ip_address : null
}
