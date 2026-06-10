variable "resource_group" {
  description = "The name of the Resource Group where the Managed Disk should exist."
  type        = string
}

variable "location" {
  description = "The Azure location where the resource exists."
  type        = string
}

variable "disk_name" {
  description = "The name of the Managed Disk."
  type        = string
}

variable "disk_size" {
  description = "The size of the managed disk in gigabytes."
  type        = number
}

variable "disk_url" {
  description = "URI to a valid VHD file to be used for the managed disk."
  type        = string
}

variable "storage_account_id" {
  description = "Resource ID of the storage account containing the VHD."
  type        = string
}

variable "storage_account_type" {
  description = "The type of storage to use for the managed disk."
  type        = string
  default     = "Standard_LRS"
}

variable "create_option_managed_disk" {
  description = "The method to use when creating the managed disk."
  type        = string
  default     = "Import"
}

variable "os_type" {
  description = "The operating system type of the managed disk."
  type        = string
  default     = "Linux"
}

variable "nic1_name" {
  description = "The name of the Network Interface 1 on subnet 1."
  type        = string
}

variable "nic2_name" {
  description = "The name of the Network Interface 2 on subnet 2."
  type        = string
}

variable "vnet_name" {
  description = "The name of the Virtual Network."
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

variable "private_ip_address_allocation" {
  description = "The allocation method used for the Private IP Address."
  type        = string
  default     = "Dynamic"
}

variable "ip_configuration_name_nic1" {
  description = "A name used for the IP Configuration of NIC 1."
  type        = string
  default     = "internal1"
}

variable "ip_configuration_name_nic2" {
  description = "A name used for the IP Configuration of NIC 2."
  type        = string
  default     = "internal2"
}

variable "create_option_storage_os_disk_for_vm" {
  description = "Specifies how the OS Disk should be created."
  type        = string
  default     = "Attach"
}

variable "caching" {
  description = "Specifies the caching requirements for the OS Disk."
  type        = string
  default     = "ReadWrite"
}

variable "os_type_on_storage_os_disk" {
  description = "Specifies the Operating System on the OS Disk."
  type        = string
  default     = "Linux"
}

variable "delete_os_disk_on_termination" {
  description = "Should the OS Disk (either the Managed Disk / VHD Blob) be deleted when the Virtual Machine is destroyed."
  type        = bool
  default     = false
}

variable "subnet3_name" {
  description = "Name of subnet 3 (used by NIC 3). Required when enable_ha = true."
  type        = string
  default     = null

  validation {
    condition     = !var.enable_ha || (var.subnet3_name != null && length(var.subnet3_name) > 0)
    error_message = "subnet3_name must be set when enable_ha = true."
  }
}

variable "ip_configuration_name_nic3" {
  description = "The name of the IP Configuration for NIC 3."
  type        = string
  default     = "internal3"
}

variable "nic3_name" {
  description = "The name of the Network Interface 3 on subnet 3. Required when enable_ha = true."
  type        = string
  default     = null

  validation {
    condition     = !var.enable_ha || (var.nic3_name != null && length(var.nic3_name) > 0)
    error_message = "nic3_name must be set when enable_ha = true."
  }
}

variable "enable_ha" {
  description = "Enable High Availability for the Azure VM."
  type        = bool
  default     = false
}

variable "is_primary" {
  description = "Indicates if this node is the primary node in a HA setup."
  type        = bool
  default     = false
}

variable "identity_id" {
  description = <<-EOT
    Resource ID of the User-Assigned Managed Identity to attach to the VM for HA
    operations.Required when enable_ha = true.
    Example: /subscriptions/<sub>/resourceGroups/<rg>/providers/Microsoft.ManagedIdentity/userAssignedIdentities/<name>
  EOT
  type        = string
  default     = null

  validation {
    condition     = !var.enable_ha || (var.identity_id != null && length(var.identity_id) > 0)
    error_message = "identity_id must be set when enable_ha = true."
  }
}

variable "tags" {
  description = "A map of tags to apply to all resources created by this module (managed disk, NICs, VM)."
  type        = map(string)
  default     = {}
}

variable "enable_ipv6" {
  description = "Whether to create an additional IPv6 IP Configuration."
  type        = bool
  default     = false
}
