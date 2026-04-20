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
