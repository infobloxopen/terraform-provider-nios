# NIOS Grid Member Azure Module

## Overview

This module provisions the Azure infrastructure (Virtual Machine, NICs, Managed Disk, etc.) for NIOS Grid. The NIOS configuration (`nios_grid_member` and `nios_grid_join` resources) should be applied after the infrastructure is deployed and NIOS grid is fully booted (~15-25 minutes).

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | >= 4.0.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_azurerm"></a> [azurerm](#provider\_azurerm) | >= 4.0.0 |

## Resources

| Name | Type |
|------|------|
| [azurerm_managed_disk.disk](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/managed_disk) | resource |
| [azurerm_network_interface.nic1](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_interface) | resource |
| [azurerm_network_interface.nic2](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_interface) | resource |
| [azurerm_virtual_machine.vm](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/virtual_machine) | resource |
| [azurerm_resource_group.rg](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/data-sources/resource_group) | data source |
| [azurerm_subnet.subnet1](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/data-sources/subnet) | data source |
| [azurerm_subnet.subnet2](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/data-sources/subnet) | data source |
| [azurerm_virtual_network.vnet](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/data-sources/virtual_network) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_disk_name"></a> [disk\_name](#input\_disk\_name) | Name for the managed OS disk. | `string` | n/a | yes |
| <a name="input_disk_size"></a> [disk\_size](#input\_disk\_size) | Size of the OS disk in GB. | `number` | n/a | yes |
| <a name="input_disk_url"></a> [disk\_url](#input\_disk\_url) | Source VHD URL for importing the NIOS managed disk image. | `string` | n/a | yes |
| <a name="input_location"></a> [location](#input\_location) | Azure region for the deployment. | `string` | n/a | yes |
| <a name="input_nic1_name"></a> [nic1\_name](#input\_nic1\_name) | Name for NIC 1 (primary interface on subnet 1). | `string` | n/a | yes |
| <a name="input_nic2_name"></a> [nic2\_name](#input\_nic2\_name) | Name for NIC 2 (secondary interface on subnet 2). | `string` | n/a | yes |
| <a name="input_resource_group"></a> [resource\_group](#input\_resource\_group) | Name of the existing Azure resource group. | `string` | n/a | yes |
| <a name="input_storage_account_id"></a> [storage\_account\_id](#input\_storage\_account\_id) | Resource ID of the storage account containing the VHD. | `string` | n/a | yes |
| <a name="input_subnet1_name"></a> [subnet1\_name](#input\_subnet1\_name) | Name of subnet 1 (used by NIC 1). | `string` | n/a | yes |
| <a name="input_subnet2_name"></a> [subnet2\_name](#input\_subnet2\_name) | Name of subnet 2 (used by NIC 2). | `string` | n/a | yes |
| <a name="input_vm_name"></a> [vm\_name](#input\_vm\_name) | Name for the Azure Virtual Machine. | `string` | n/a | yes |
| <a name="input_vm_size"></a> [vm\_size](#input\_vm\_size) | Azure VM size (e.g. Standard\_E4s\_v5). | `string` | n/a | yes |
| <a name="input_vnet_name"></a> [vnet\_name](#input\_vnet\_name) | Name of the existing Azure Virtual Network. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_instance_id"></a> [instance\_id](#output\_instance\_id) | ID of the NIOS Grid Member instance. |
| <a name="output_nic1_ip"></a> [nic1\_ip](#output\_nic1\_ip) | Private IP address of NIC1 (Subnet 1) |
| <a name="output_nic2_ip"></a> [nic2\_ip](#output\_nic2\_ip) | Private IP address of NIC2 (Subnet 2) |
| <a name="output_subnet1_gateway"></a> [subnet1\_gateway](#output\_subnet1\_gateway) | Gateway IP for Subnet 1 (first usable IP) |
| <a name="output_subnet1_mask"></a> [subnet1\_mask](#output\_subnet1\_mask) | Subnet mask of Subnet 1 |
| <a name="output_subnet2_gateway"></a> [subnet2\_gateway](#output\_subnet2\_gateway) | Gateway IP for Subnet 2 (first usable IP) |
| <a name="output_subnet2_mask"></a> [subnet2\_mask](#output\_subnet2\_mask) | Subnet mask of Subnet 2 |
<!-- END_TF_DOCS -->

---

## Architecture

### Standalone Mode
- 1 Azure VM with NIOS VHD image (imported as Managed Disk)
- NIC1: Primary interface on Subnet 1
- NIC2: Secondary interface on Subnet 2

## Prerequisites

Before using this module, ensure the following Azure resources exist:
- **Resource Group** — the target resource group for all resources
- **Virtual Network** — with at least two subnets
- **Storage Account** — containing the NIOS VHD image blob
- **Service Principal** — with appropriate permissions (subscription\_id, client\_id, client\_secret, tenant\_id)

## Usage

### Step 1: Deploy Azure Infrastructure

```hcl
provider "azurerm" {
  features {}

  subscription_id = var.subscription_id
  client_id       = var.client_id
  client_secret   = var.client_secret
  tenant_id       = var.tenant_id
}

module "node1" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_grid_member_azure"

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
```

**Deploy the infrastructure:**
```bash
terraform apply
```

### Step 2: Wait for NIOS to Boot

NIOS takes approximately **15 to 20 minutes** to fully boot.

### Step 3: Join the Grid Member to the Master Grid

Once Grid is up and running, configure the grid member and join to the grid.

#### Example: Join a Member to a Master

##### Deploy Azure infrastructure for Master and Member

```hcl
module "node1" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_grid_member_azure"
  // ... (same config as Step 1)
}

module "node2" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_grid_member_azure"
  // ... (same config as Step 1)
}
```

##### After NIOS is ready (~20 min), configure grid member

```hcl
provider "nios" {
  nios_host_url = "https://${module.node1.nic1_ip}"
  nios_username = "username"
  nios_password = "password"
}

resource "nios_grid_member" "member" {
  host_name        = "infoblox.member"
  config_addr_type = "IPV4"
  platform         = "VNIOS"

  vip_setting = {
    address     = module.node2.nic1_ip
    gateway     = module.node2.subnet1_gateway
    subnet_mask = module.node2.subnet1_mask
  }
}

// Join member to existing grid master
resource "nios_grid_join" "member_join" {
  member_url      = "https://${module.node2.nic1_ip}"
  member_username = "Username"
  member_password = "Password"
  grid_name       = "Infoblox"
  master          = module.node1.nic1_ip
  shared_secret   = "<secret>"
  depends_on      = [nios_grid_member.member]
}
```

## Outputs Usage

The module outputs can be used directly in NIOS provider resources:

| Output | NIOS Resource Usage |
|--------|---------------------|
| `nic2_ip` | `vip_setting.address`, `member_url` in grid_join |
| `subnet2_gateway` | `vip_setting.gateway` |
| `subnet2_mask` | `vip_setting.subnet_mask` |
| `nic1_ip` | Management access |

---

### Boot Time
- NIOS takes **15-20 minutes** to fully boot after VM creation
- Always verify NIOS API is responding before applying `nios_grid_member` resources
