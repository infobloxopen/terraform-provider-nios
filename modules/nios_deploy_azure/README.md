# Deploy vNIOS on Azure

## Overview

This module provisions vNIOS on Azure. The NIOS configuration (`nios_grid_member` and `nios_grid_join` resources) should be applied after the infrastructure is deployed and NIOS grid is fully booted (~30 minutes).

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.12.1 |
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
| [azurerm_network_interface.nic3](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_interface) | resource |
| [azurerm_virtual_machine.vm](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/virtual_machine) | resource |
| [azurerm_resource_group.rg](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/data-sources/resource_group) | data source |
| [azurerm_subnet.subnet1](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/data-sources/subnet) | data source |
| [azurerm_subnet.subnet2](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/data-sources/subnet) | data source |
| [azurerm_subnet.subnet3](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/data-sources/subnet) | data source |
| [azurerm_virtual_network.vnet](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/data-sources/virtual_network) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_caching"></a> [caching](#input\_caching) | Specifies the caching requirements for the OS Disk. | `string` | `"ReadWrite"` | no |
| <a name="input_create_option_managed_disk"></a> [create\_option\_managed\_disk](#input\_create\_option\_managed\_disk) | The method to use when creating the managed disk. | `string` | `"Import"` | no |
| <a name="input_create_option_storage_os_disk_for_vm"></a> [create\_option\_storage\_os\_disk\_for\_vm](#input\_create\_option\_storage\_os\_disk\_for\_vm) | Specifies how the OS Disk should be created. | `string` | `"Attach"` | no |
| <a name="input_delete_os_disk_on_termination"></a> [delete\_os\_disk\_on\_termination](#input\_delete\_os\_disk\_on\_termination) | Should the OS Disk (either the Managed Disk / VHD Blob) be deleted when the Virtual Machine is destroyed. | `bool` | `false` | no |
| <a name="input_disk_name"></a> [disk\_name](#input\_disk\_name) | The name of the Managed Disk. | `string` | n/a | yes |
| <a name="input_disk_size"></a> [disk\_size](#input\_disk\_size) | The size of the managed disk in gigabytes. | `number` | n/a | yes |
| <a name="input_disk_url"></a> [disk\_url](#input\_disk\_url) | URI to a valid VHD file to be used for the managed disk. | `string` | n/a | yes |
| <a name="input_enable_ha"></a> [enable\_ha](#input\_enable\_ha) | Enable High Availability for the Azure VM. | `bool` | `false` | no |
| <a name="input_enable_ipv6"></a> [enable\_ipv6](#input\_enable\_ipv6) | Whether to create an additional IPv6 IP Configuration. | `bool` | `false` | no |
| <a name="input_identity_id"></a> [identity\_id](#input\_identity\_id) | Resource ID of the User-Assigned Managed Identity to attach to the VM for HA<br/>operations. Required when enable\_ha = true.<br/>Example: /subscriptions/<sub>/resourceGroups/<rg>/providers/Microsoft.ManagedIdentity/userAssignedIdentities/<name> | `string` | `null` | no |
| <a name="input_ip_configuration_name_nic1"></a> [ip\_configuration\_name\_nic1](#input\_ip\_configuration\_name\_nic1) | A name used for the IP Configuration of NIC 1. | `string` | `"internal1"` | no |
| <a name="input_ip_configuration_name_nic2"></a> [ip\_configuration\_name\_nic2](#input\_ip\_configuration\_name\_nic2) | A name used for the IP Configuration of NIC 2. | `string` | `"internal2"` | no |
| <a name="input_ip_configuration_name_nic3"></a> [ip\_configuration\_name\_nic3](#input\_ip\_configuration\_name\_nic3) | The name of the IP Configuration for NIC 3. | `string` | `"internal3"` | no |
| <a name="input_is_primary"></a> [is\_primary](#input\_is\_primary) | Indicates if this node is the primary node in a HA setup. | `bool` | `false` | no |
| <a name="input_location"></a> [location](#input\_location) | The Azure location where the resource exists. | `string` | n/a | yes |
| <a name="input_nic1_name"></a> [nic1\_name](#input\_nic1\_name) | The name of the Network Interface 1 on subnet 1. | `string` | n/a | yes |
| <a name="input_nic2_name"></a> [nic2\_name](#input\_nic2\_name) | The name of the Network Interface 2 on subnet 2. | `string` | n/a | yes |
| <a name="input_nic3_name"></a> [nic3\_name](#input\_nic3\_name) | The name of the Network Interface 3 on subnet 3. Required when enable\_ha = true. | `string` | `null` | no |
| <a name="input_os_type"></a> [os\_type](#input\_os\_type) | The operating system type of the managed disk. | `string` | `"Linux"` | no |
| <a name="input_os_type_on_storage_os_disk"></a> [os\_type\_on\_storage\_os\_disk](#input\_os\_type\_on\_storage\_os\_disk) | Specifies the Operating System on the OS Disk. | `string` | `"Linux"` | no |
| <a name="input_private_ip_address_allocation"></a> [private\_ip\_address\_allocation](#input\_private\_ip\_address\_allocation) | The allocation method used for the Private IP Address. | `string` | `"Dynamic"` | no |
| <a name="input_resource_group"></a> [resource\_group](#input\_resource\_group) | The name of the Resource Group where the Managed Disk should exist. | `string` | n/a | yes |
| <a name="input_storage_account_id"></a> [storage\_account\_id](#input\_storage\_account\_id) | Resource ID of the storage account containing the VHD. | `string` | n/a | yes |
| <a name="input_storage_account_type"></a> [storage\_account\_type](#input\_storage\_account\_type) | The type of storage to use for the managed disk. | `string` | `"Standard_LRS"` | no |
| <a name="input_subnet1_name"></a> [subnet1\_name](#input\_subnet1\_name) | Name of subnet 1 (used by NIC 1). | `string` | n/a | yes |
| <a name="input_subnet2_name"></a> [subnet2\_name](#input\_subnet2\_name) | Name of subnet 2 (used by NIC 2). | `string` | n/a | yes |
| <a name="input_subnet3_name"></a> [subnet3\_name](#input\_subnet3\_name) | Name of subnet 3 (used by NIC 3). Required when enable\_ha = true. | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to apply to all resources created by this module (managed disk, NICs, VM). | `map(string)` | `{}` | no |
| <a name="input_vm_name"></a> [vm\_name](#input\_vm\_name) | Name for the Azure Virtual Machine. | `string` | n/a | yes |
| <a name="input_vm_size"></a> [vm\_size](#input\_vm\_size) | Azure VM size (e.g. Standard\_E4s\_v5). | `string` | n/a | yes |
| <a name="input_vnet_name"></a> [vnet\_name](#input\_vnet\_name) | The name of the Virtual Network. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_instance_id"></a> [instance\_id](#output\_instance\_id) | ID of the NIOS Grid Member instance. |
| <a name="output_nic1_ip"></a> [nic1\_ip](#output\_nic1\_ip) | Private IP address of NIC1 (Subnet 1) |
| <a name="output_nic1_ipv6"></a> [nic1\_ipv6](#output\_nic1\_ipv6) | IPv6 address of NIC1 (Subnet 1). Null when enable\_ipv6 is false. |
| <a name="output_nic2_ip"></a> [nic2\_ip](#output\_nic2\_ip) | Private IP address of NIC2 (Subnet 2) |
| <a name="output_nic3_ip"></a> [nic3\_ip](#output\_nic3\_ip) | Private IP address of NIC3 (HA interface). Null when HA is disabled. |
| <a name="output_nic3_name"></a> [nic3\_name](#output\_nic3\_name) | Name of the HA NIC (NIC3). Null when HA is disabled. |
| <a name="output_subnet1_gateway"></a> [subnet1\_gateway](#output\_subnet1\_gateway) | Gateway IP for Subnet 1 (first usable IP) |
| <a name="output_subnet1_mask"></a> [subnet1\_mask](#output\_subnet1\_mask) | Subnet mask of Subnet 1 |
| <a name="output_subnet2_gateway"></a> [subnet2\_gateway](#output\_subnet2\_gateway) | Gateway IP for Subnet 2 (first usable IP) |
| <a name="output_subnet2_mask"></a> [subnet2\_mask](#output\_subnet2\_mask) | Subnet mask of Subnet 2 |
| <a name="output_subnet3_gateway"></a> [subnet3\_gateway](#output\_subnet3\_gateway) | Gateway IP for Subnet 3 (first usable IP). Null when HA is disabled. |
| <a name="output_subnet3_mask"></a> [subnet3\_mask](#output\_subnet3\_mask) | Subnet mask of Subnet 3 (HA subnet). Null when HA is disabled. |
| <a name="output_vip"></a> [vip](#output\_vip) | HA VIP (secondary IP on NIC3). Null when HA disabled or node is not primary. |
<!-- END_TF_DOCS -->

---

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
```

**Deploy the infrastructure:**
```bash
terraform apply
```

### Step 2: Wait for NIOS to Boot

NIOS takes approximately **30 minutes** to fully boot, make sure the grid is up and running before triggering the grid join.

> **Important — Manual license installation required on Azure.**
> Unlike OCI/AWS, the Azure vNIOS image does **not** support cloud-init, so the NIOS
> temporary license cannot be injected at boot. After the VM is up, you must log into
> the NIOS CLI (SSH/console) on each node and install the appropriate licenses
> manually **before** applying the `nios_grid_member` / `nios_grid_join`
> resources. The grid will not form until the required licenses are present on
> each node.

### Step 3: Join the Grid Member to the Master Grid

Once Grid is up and running, configure the grid member and join to the grid.

#### Examples

#### Example 1: Join a Member to a Master

##### Deploy Azure infrastructure for Master and Member

```hcl
module "node1" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_azure?ref=nios_v9.1.0"
  // ... (same config as Step 1)
}

module "node2" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_azure?ref=nios_v9.1.0"
  // ... (same config as Step 1)
}
```

##### After NIOS is ready (~30 min), configure grid member

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

#### Example 2: Join a Member with Dual Stack Config

To provision an IPv6 address on the NIC, set `enable_ipv6 = true`. The
member's IPv6 address is then exposed through the `nic1_ipv6` output.

##### Deploy Azure infrastructure for Master and Member

```hcl
module "node1" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_azure?ref=nios_v9.1.0"
  // ... (same config as Step 1)

  enable_ipv6 = true // enable IPv6 on the member NIC
}

module "node2" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_azure?ref=nios_v9.1.0"
  // ... (same config as Step 1)

  enable_ipv6 = true // enable IPv6 on the member NIC
}
```

##### After NIOS is ready (~30 min), configure grid member

> **Note:** On Azure, IPv6 addresses are available in the Terraform state (the
> `nic1_ipv6` output), but you must manually configure the IPv6 settings on the NIOS grid through the UI or API. Unlike AWS, Azure does not automatically configure IPv6 on the NIOS instance. This is a known limitation that will be addressed in a future NIOS release.

```hcl
provider "nios" {
  nios_host_url = "https://${module.node1.nic1_ip}"
  nios_username = "username"
  nios_password = "password"
}

resource "nios_grid_member" "member" {
  host_name        = "infoblox.member"
  config_addr_type = "BOTH"
  platform         = "VNIOS"

  vip_setting = {
    address     = module.node2.nic1_ip
    gateway     = module.node2.subnet1_gateway
    subnet_mask = module.node2.subnet1_mask
  }

  ipv6_setting = {
    virtual_ip  = module.node2.nic1_ipv6
    cidr_prefix = 64
    gateway     = "<member_ipv6_gateway_ip>"
    enabled     = true
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

#### Example 3: HA Grid Configuration

Deploy two Azure VMs for SA-HA configuration. HA on Azure requires a third NIC
(`nic3`) on a dedicated HA subnet (`subnet3`) and a User-Assigned Managed
Identity (`identity_id`) attached to each VM so the NIOS appliance can move the
VIP between nodes during failover.

```hcl
// Deploy Azure infrastructure for Node 1 (Active Node)
module "node1" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_azure?ref=nios_v9.1.0"
  // ... (same base config as Step 1)

  nic3_name    = "${var.nic3_name}-node1"
  subnet3_name = var.subnet3_name
  enable_ha    = true
  is_primary   = true
  identity_id  = var.identity_id
}

// Deploy Azure infrastructure for Node 2 (Passive Node)
module "node2" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_azure?ref=nios_v9.1.0"
  // ... (same base config as Step 1)

  nic3_name    = "${var.nic3_name}-node2"
  subnet3_name = var.subnet3_name
  enable_ha    = true
  is_primary   = false
  identity_id  = var.identity_id
}
```

#### After both nodes are up and running (~30 min) and licenses are installed, configure HA

1. Import Node1 under `nios_grid_member.ha_pair` and configure it as the HA pair.

```hcl
provider "nios" {
  nios_host_url = "https://${module.node1.nic1_ip}"
  nios_username = "<username>"
  nios_password = "<password>"
}

import {
  to = nios_grid_member.ha_pair
  id = "5e57eac828fd49fd90072931fa4eff3f"
}

resource "nios_grid_member" "ha_pair" {
  config_addr_type = "IPV4"
  host_name        = "infoblox.localdomain"
  platform         = "VNIOS"

  upgrade_group    = "Grid Master"
  master_candidate = true

  vip_setting = {
    address     = module.node1.vip
    gateway     = module.node1.subnet3_gateway
    subnet_mask = module.node1.subnet3_mask
  }

  enable_ha         = true
  ha_on_cloud       = true
  ha_cloud_platform = "AZURE"

  node_info = [
    {
      lan_ha_port_setting = {
        ha_ip_address      = module.node1.nic3_ip
        mgmt_lan           = module.node1.nic1_ip
        ha_cloud_attribute = module.node1.nic3_name
      }
    },
    {
      lan_ha_port_setting = {
        ha_ip_address      = module.node2.nic3_ip
        mgmt_lan           = module.node2.nic1_ip
        ha_cloud_attribute = module.node2.nic3_name
      }
    }
  ]
  router_id = 111

  grid_level_dns_resolver_setting = {
    resolvers = [
      "10.10.10.10"
    ]
  }
}
```

Run `terraform apply` to import and reconfigure Node1 as the HA active node.

2. Join Node2 (Passive Node) to Node1 (Active Node).

```hcl
provider "nios" {
  nios_host_url = "https://${module.node1.vip}"
  nios_username = "<username>"
  nios_password = "<password>"
}

resource "nios_grid_join" "ha_member_join" {
  member_url      = "https://${module.node2.nic1_ip}"
  member_username = "<username>"
  member_password = "<password>"
  grid_name       = "Infoblox"
  master          = module.node1.vip
  shared_secret   = "<secret>"
}
```

#### Best Practices for HA Deployment

> **Recommended Workflow:** Use a **separate Terraform workspace** for HA configuration. The NIOS HA setup is a one-time provisioning task — once the HA pair is formed and the passive node has joined the grid, the configuration is complete and does not require ongoing Terraform management.

After successfully deploying the HA pair:

1. **Verify HA formation** is complete through the NIOS UI or API.
2. **Remove the HA grid member and join resources from Terraform state** to prevent accidental modifications:
   ```bash
   terraform state rm nios_grid_member.ha_pair
   terraform state rm nios_grid_join.ha_member_join
   ```
3. Optionally, you can delete the entire Terraform state for this workspace if no further infrastructure management is needed.

This approach ensures that:
- Your HA infrastructure is provisioned correctly.
- Subsequent Terraform operations don't interfere with the running HA pair.
- The grid master configuration remains stable and is managed through NIOS directly.

### Boot Time
- NIOS takes **30 minutes** to fully boot after VM creation, make sure the grid is up and running before triggering the grid join.
- Always verify NIOS API is responding before applying `nios_grid_member` resources
