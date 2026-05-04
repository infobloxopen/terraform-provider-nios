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
