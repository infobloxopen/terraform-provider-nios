output "instance_id" {
  description = "ID of the NIOS Grid Member instance."
  value       = google_compute_instance.grid.id
}

output "instance_name" {
  description = "Name of the NIOS Grid Member instance."
  value       = google_compute_instance.grid.name
}
output "mgmt_ip" {
  description = "Internal IP of the MGMT interface (nic0)."
  value       = google_compute_instance.grid.network_interface[0].network_ip
}

output "lan1_ip" {
  description = "Internal IP of the LAN interface (nic1)."
  value       = google_compute_instance.grid.network_interface[1].network_ip
}

output "ha_ip" {
  description = "Internal IP of the HA interface (nic2)."
  value       = var.enable_ha ? google_compute_instance.grid.network_interface[2].network_ip : null
}

output "mgmt_subnet_mask" {
  description = "Subnet Mask of the Mgmt Subnetwork"
  value       = cidrnetmask(data.google_compute_subnetwork.mgmt.ip_cidr_range)
}

output "mgmt_gateway" {
  description = "Gateway IP for the MGMT subnetwork."
  value       = cidrhost(data.google_compute_subnetwork.mgmt.ip_cidr_range, 1)
}

output "lan1_subnet_mask" {
  description = "Subnet mask of the LAN1 subnetwork."
  value       = cidrnetmask(data.google_compute_subnetwork.lan.ip_cidr_range)
}

output "lan1_gateway" {
  description = "Gateway IP for the LAN1 subnetwork."
  value       = cidrhost(data.google_compute_subnetwork.lan.ip_cidr_range, 1)
}

output "ha_subnet_mask" {
  description = "Subnet mask of the HA subnetwork."
  value       = var.enable_ha ? cidrnetmask(data.google_compute_subnetwork.ha[0].ip_cidr_range) : null
}

output "ha_gateway" {
  description = "Gateway IP for the HA subnetwork."
  value       = var.enable_ha ? cidrhost(data.google_compute_subnetwork.ha[0].ip_cidr_range, 1) : null
}

output "vip" {
  description = "Alias IP (floating VIP) on nic2 for HA."
  value = try(
    split("/", google_compute_instance.grid.network_interface[2].alias_ip_range[0].ip_cidr_range)[0],
    null
  )
}

output "mgmt_ipv6_address" {
  description = "IPv6 address of the MGMT interface (nic0)."
  value       = var.enable_ipv6 ? google_compute_instance.grid.network_interface[0].ipv6_address : null
}

output "lan1_ipv6_address" {
  description = "IPv6 address of the LAN1 interface (nic1)."
  value       = var.enable_ipv6 ? google_compute_instance.grid.network_interface[1].ipv6_address : null
}
