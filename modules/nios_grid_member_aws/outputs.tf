output "instance_id" {
  description = "ID of the EC2 instance"
  value       = aws_instance.grid.id
}

output "eth0_ip" {
  description = "Mgmt IP address of the EC2 instance (ETH0)"
  value       = aws_instance.grid.private_ip
}

output "eth1_ipv4" {
  description = "LAN1 IP address of the EC2 instance (ETH1)"
  value       = aws_network_interface.eth1.private_ip
}

output "eth2_ip" {
  description = "HA IP address of the EC2 instance (ETH2)"
  value       = var.ha_enable ? aws_network_interface.eth2[0].private_ip : null
}

output "eth2_secondary_ip_for_ha" {
  description = "Secondary private IP address for HA configuration on ETH2"
  value = var.ha_enable && length(aws_network_interface.eth2) > 0 ? try(
    [for ip in tolist(aws_network_interface.eth2[0].private_ips) : ip if ip != aws_network_interface.eth2[0].private_ip][0],
    null
  ) : null
}

output "eth1_ipv6" {
  description = "Private IPv6 address of ETH1"
  value       = length(aws_network_interface.eth1.ipv6_addresses) > 0 ? tolist(aws_network_interface.eth1.ipv6_addresses)[0] : null
}

output "eth1_subnet_mask" {
  description = "Subnet mask of the LAN1 Subnet"
  value       = cidrnetmask(data.aws_subnet.lan1.cidr_block)
}

output "eth1_gateway" {
  description = "Gateway IP for the LAN1 Subnet"
  value       = cidrhost(data.aws_subnet.lan1.cidr_block, 1)
}

output "eth2_eni" {
  description = "ENI ID of the HA interface (ETH2)"
  value = var.ha_enable ? (
    length(aws_network_interface.eth2) > 0 ?
    aws_network_interface.eth2[0].id
    : null
  ) : null
}
