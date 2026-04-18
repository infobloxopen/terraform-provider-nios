output "instance_id" {
  description = "ID of the NIOS Grid Member instance."
  value       = google_compute_instance.grid.id
}

output "mgmt_ip" {
  description = "Internal IP of the MGMT interface (nic0)."
  value       = google_compute_instance.grid.network_interface[0].network_ip
}

output "mgmt_subnet_mask" {
  description = "Subnet Mask of the Mgmt Subnetwork"
  value       = cidrnetmask(data.google_compute_subnetwork.mgmt.ip_cidr_range)
}

output "mgmt_gateway" {
  description = "Gateway IP for the MGMT subnetwork (first usable IP)."
  value       = cidrhost(data.google_compute_subnetwork.mgmt.ip_cidr_range, 1)
  
}

output "lan1_ip" {
  description = "Internal IP of the LAN1 interface (nic1)."
  value       = google_compute_instance.grid.network_interface[1].network_ip
}

output "lan1_subnet_mask" {
  description = "Subnet mask of the LAN1 subnetwork."
  value       = cidrnetmask(data.google_compute_subnetwork.lan1.ip_cidr_range)
}

output "lan1_gateway" {
  description = "Gateway IP for the LAN1 subnetwork (first usable IP)."
  value       = cidrhost(data.google_compute_subnetwork.lan1.ip_cidr_range, 1)
}
