provider "azurerm" {
  features {}

  subscription_id = var.subscription_id
  client_id       = var.client_id
  client_secret   = var.client_secret
  tenant_id       = var.tenant_id
}

module "node1" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_azure?ref=nios_v9.1.0"

  resource_group = var.resource_group
  location       = var.location

  vnet_name    = var.vnet_name
  subnet1_name = var.subnet1_name
  subnet2_name = var.subnet2_name

  disk_name          = var.disk_name
  disk_size          = var.disk_size
  disk_url           = var.disk_url
  storage_account_id = var.storage_account_id

  nic1_name = var.nic1_name
  nic2_name = var.nic2_name

  vm_name = var.vm_name
  vm_size = var.vm_size
}

// HA Configuration
// To create the primary node for an HA pair, set enable_ha = true and is_primary = true.
module "node2" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_azure?ref=nios_v9.1.0"

  resource_group = var.resource_group
  location       = var.location

  vnet_name    = var.vnet_name
  subnet1_name = var.subnet1_name
  subnet2_name = var.subnet2_name

  disk_name          = "${var.disk_name}-node2"
  disk_size          = var.disk_size
  disk_url           = var.disk_url
  storage_account_id = var.storage_account_id

  nic1_name = "${var.nic1_name}-node2"
  nic2_name = "${var.nic2_name}-node2"

  vm_name = "${var.vm_name}-node2"
  vm_size = var.vm_size

  nic3_name    = "${var.nic3_name}-node2"
  subnet3_name = var.subnet3_name
  enable_ha    = true
  is_primary   = true
  identity_id  = var.identity_id

  tags = var.tags
}

// To create the secondary node for an HA pair, set enable_ha = true and is_primary = false.
module "node3" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_azure?ref=nios_v9.1.0"

  resource_group = var.resource_group
  location       = var.location

  vnet_name    = var.vnet_name
  subnet1_name = var.subnet1_name
  subnet2_name = var.subnet2_name

  disk_name          = "${var.disk_name}-node3"
  disk_size          = var.disk_size
  disk_url           = var.disk_url
  storage_account_id = var.storage_account_id

  nic1_name = "${var.nic1_name}-node3"
  nic2_name = "${var.nic2_name}-node3"

  vm_name = "${var.vm_name}-node3"
  vm_size = var.vm_size

  nic3_name    = "${var.nic3_name}-node3"
  subnet3_name = var.subnet3_name
  enable_ha    = true
  is_primary   = false
  identity_id  = var.identity_id

  tags = var.tags
}

// Dual-stack Configuration 
// Set enable_ipv6 = true to create an additional IPv6 configuration on each NIC.
module "node4" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_azure?ref=nios_v9.1.0"

  resource_group = var.resource_group
  location       = var.location

  vnet_name    = var.vnet_name
  subnet1_name = var.subnet1_name
  subnet2_name = var.subnet2_name

  disk_name          = "${var.disk_name}-node4"
  disk_size          = var.disk_size
  disk_url           = var.disk_url
  storage_account_id = var.storage_account_id

  nic1_name = "${var.nic1_name}-node4"
  nic2_name = "${var.nic2_name}-node4"

  vm_name = "${var.vm_name}-node4"
  vm_size = var.vm_size

  enable_ipv6 = true

  tags = var.tags
}
