variable "resource_group" {
  description = "Name of the existing Azure resource group."
  type        = string
}

variable "location" {
  description = "Azure region for the deployment."
  type        = string
}

variable "disk_name" {
  description = "Name for the managed OS disk."
  type        = string
}

variable "disk_size" {
  description = "Size of the OS disk in GB."
  type        = number
}

variable "disk_url" {
  description = "Source VHD URL for importing the NIOS managed disk image."
  type        = string
}

variable "storage_account_id" {
  description = "Resource ID of the storage account containing the VHD."
  type        = string
}

variable "nic1_name" {
  description = "Name for NIC 1 (primary interface on subnet 1)."
  type        = string
}

variable "nic2_name" {
  description = "Name for NIC 2 (secondary interface on subnet 2)."
  type        = string
}

variable "vnet_name" {
  description = "Name of the existing Azure Virtual Network."
  type        = string
}

variable "subnet1_name" {
  description = "Name of subnet 1 (used by NIC 1)."
  type        = string
}

variable "subnet2_name" {
  description = "Name of subnet 2 (used by NIC 2)."
  type        = string
}

variable "vm_name" {
  description = "Name for the Azure Virtual Machine."
  type        = string
}

variable "vm_size" {
  description = "Azure VM size (e.g. Standard_E4s_v5)."
  type        = string
}
