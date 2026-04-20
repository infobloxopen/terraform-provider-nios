// Retrieve information about existing Azure Resource Group
data "azurerm_resource_group" "rg" {
  name = var.resource_group
}

// Retrieve information about existing Azure Virtual Network
data "azurerm_virtual_network" "vnet" {
  name                = var.vnet_name
  resource_group_name = data.azurerm_resource_group.rg.name
}

// Retrieve information about existing Subnet 1
data "azurerm_subnet" "subnet1" {
  name                 = var.subnet1_name
  virtual_network_name = data.azurerm_virtual_network.vnet.name
  resource_group_name  = data.azurerm_resource_group.rg.name
}

// Retrieve information about existing Subnet 2
data "azurerm_subnet" "subnet2" {
  name                 = var.subnet2_name
  virtual_network_name = data.azurerm_virtual_network.vnet.name
  resource_group_name  = data.azurerm_resource_group.rg.name
}

// Managed Disk imported from VHD
resource "azurerm_managed_disk" "disk" {
  name                 = var.disk_name
  location             = var.location
  resource_group_name  = data.azurerm_resource_group.rg.name
  storage_account_type = var.storage_account_type
  create_option        = var.create_option_managed_disk
  storage_account_id   = var.storage_account_id
  source_uri           = var.disk_url
  os_type              = var.os_type
  disk_size_gb         = var.disk_size
}

// Network Interface for NIC 1 (primary interface)
resource "azurerm_network_interface" "nic1" {
  name                = var.nic1_name
  location            = var.location
  resource_group_name = data.azurerm_resource_group.rg.name

  ip_configuration {
    name                          = var.ip_configuration_name_nic1
    subnet_id                     = data.azurerm_subnet.subnet1.id
    private_ip_address_allocation = var.private_ip_address_allocation
    primary                       = true
  }
}

// Manage a Network Interface for NIC 2 (secondary interface)
resource "azurerm_network_interface" "nic2" {
  name                = var.nic2_name
  location            = var.location
  resource_group_name = data.azurerm_resource_group.rg.name

  ip_configuration {
    name                          = var.ip_configuration_name_nic2
    subnet_id                     = data.azurerm_subnet.subnet2.id
    private_ip_address_allocation = var.private_ip_address_allocation
    primary                       = true
  }
}

// Manage a Virtual Machine for NIOS Grid Member
resource "azurerm_virtual_machine" "vm" {
  name                = var.vm_name
  location            = var.location
  resource_group_name = data.azurerm_resource_group.rg.name
  vm_size             = var.vm_size

  network_interface_ids = [
    azurerm_network_interface.nic1.id,
    azurerm_network_interface.nic2.id
  ]

  primary_network_interface_id = azurerm_network_interface.nic1.id

  delete_os_disk_on_termination = var.delete_os_disk_on_termination

  storage_os_disk {
    name            = var.disk_name
    managed_disk_id = azurerm_managed_disk.disk.id
    create_option   = var.create_option_storage_os_disk_for_vm
    caching         = var.caching
    os_type         = var.os_type_on_storage_os_disk
  }
}
