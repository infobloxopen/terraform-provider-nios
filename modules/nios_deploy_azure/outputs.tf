output "instance_id" {
  description = "ID of the NIOS Grid Member instance."
  value       = azurerm_virtual_machine.vm.id
}
output "nic1_ip" {
  description = "Private IP address of NIC1 (Subnet 1)"
  value       = azurerm_network_interface.nic1.private_ip_address
}

output "nic2_ip" {
  description = "Private IP address of NIC2 (Subnet 2)"
  value       = azurerm_network_interface.nic2.private_ip_address
}

output "subnet1_mask" {
  description = "Subnet mask of Subnet 1"
  value       = cidrnetmask(data.azurerm_subnet.subnet1.address_prefixes[0])
}

output "subnet1_gateway" {
  description = "Gateway IP for Subnet 1 (first usable IP)"
  value       = cidrhost(data.azurerm_subnet.subnet1.address_prefixes[0], 1)
}

output "subnet2_mask" {
  description = "Subnet mask of Subnet 2"
  value       = cidrnetmask(data.azurerm_subnet.subnet2.address_prefixes[0])
}

output "subnet2_gateway" {
  description = "Gateway IP for Subnet 2 (first usable IP)"
  value       = cidrhost(data.azurerm_subnet.subnet2.address_prefixes[0], 1)
}

output "nic3_ip" {
  description = "Private IP address of NIC3 (HA interface). Null when HA is disabled."
  value       = var.enable_ha ? azurerm_network_interface.nic3[0].private_ip_address : null
}

// VIP — secondary IP on NIC3, only present on the primary HA node.
output "vip" {
  description = "HA VIP (secondary IP on NIC3). Null when HA disabled or node is not primary."
  value = (var.enable_ha && var.is_primary
    ? one([
      for cfg in azurerm_network_interface.nic3[0].ip_configuration :
      cfg.private_ip_address if cfg.name == "${var.nic3_name}-vip"
    ])
    : null
  )
}

output "subnet3_mask" {
  description = "Subnet mask of Subnet 3 (HA subnet). Null when HA is disabled."
  value       = var.enable_ha ? cidrnetmask(data.azurerm_subnet.subnet3[0].address_prefixes[0]) : null
}

output "subnet3_gateway" {
  description = "Gateway IP for Subnet 3 (first usable IP). Null when HA is disabled."
  value       = var.enable_ha ? cidrhost(data.azurerm_subnet.subnet3[0].address_prefixes[0], 1) : null
}

output "nic3_name" {
  description = "Name of the HA NIC (NIC3). Null when HA is disabled."
  value       = var.enable_ha ? azurerm_network_interface.nic3[0].name : null
}

output "nic1_ipv6" {
  description = "IPv6 address of NIC1 (Subnet 1). Null when enable_ipv6 is false."
  value = var.enable_ipv6 ? one([
    for cfg in azurerm_network_interface.nic1.ip_configuration :
    cfg.private_ip_address if cfg.name == "${var.ip_configuration_name_nic1}-v6"
  ]) : null
}
