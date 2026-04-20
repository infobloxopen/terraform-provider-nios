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
