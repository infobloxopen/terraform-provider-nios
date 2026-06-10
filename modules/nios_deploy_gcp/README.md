# Deploy vNIOS on GCP

## Overview

This module provisions vNIOS on GCP. Use one module call per instance — Grid Master, IB member, CP member, Reporting, or Discovery — they all share the same resource structure. The NIOS configuration (`nios_grid_member` and `nios_grid_join` resources) should be applied after the infrastructure is deployed and NIOS grid is fully booted (~30 minutes).

### NIOS Model -> Machine Type Mapping

The module automatically maps NIOS models to GCP machine types:

| NIOS Model | GCP Machine Type |
|------------|-----------------|
| IB-V825 / TE-V810 / CP-V800 | n2-standard-2 |
| IB-V1425 / TE-V1410 / CP-V1400 | n2-standard-4 |
| IB-V2225 / TE-V2210 / CP-V2200 | n2-standard-8 |
| IB-V4025 / TE-V4010 / CP-V4000 | n2-standard-16 |

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.12.1 |
| <a name="requirement_google"></a> [google](#requirement\_google) | >= 5.0.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_google"></a> [google](#provider\_google) | >= 5.0.0 |

## Resources

| Name | Type |
|------|------|
| [google_compute_instance.grid](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_instance) | resource |
| [google_compute_subnetwork.ha](https://registry.terraform.io/providers/hashicorp/google/latest/docs/data-sources/compute_subnetwork) | data source |
| [google_compute_subnetwork.lan1](https://registry.terraform.io/providers/hashicorp/google/latest/docs/data-sources/compute_subnetwork) | data source |
| [google_compute_subnetwork.mgmt](https://registry.terraform.io/providers/hashicorp/google/latest/docs/data-sources/compute_subnetwork) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_boot_disk_size"></a> [boot\_disk\_size](#input\_boot\_disk\_size) | The size of the boot disk in GB. | `number` | `250` | no |
| <a name="input_boot_disk_type"></a> [boot\_disk\_type](#input\_boot\_disk\_type) | The type of the boot disk. | `string` | `"pd-standard"` | no |
| <a name="input_default_admin_password"></a> [default\_admin\_password](#input\_default\_admin\_password) | The default admin password for the NIOS instance. | `string` | n/a | yes |
| <a name="input_enable_ha"></a> [enable\_ha](#input\_enable\_ha) | Whether to enable high availability (HA) for the NIOS instance. | `bool` | `false` | no |
| <a name="input_enable_ipv6"></a> [enable\_ipv6](#input\_enable\_ipv6) | Enable IPv6 (dual-stack) on network interfaces | `bool` | `false` | no |
| <a name="input_ha_subnet_name"></a> [ha\_subnet\_name](#input\_ha\_subnet\_name) | The name of the subnetwork to attach to the high availability network interface (nic2). | `string` | `null` | no |
| <a name="input_image_name"></a> [image\_name](#input\_image\_name) | The image from which to initialize this disk. | `string` | n/a | yes |
| <a name="input_is_primary"></a> [is\_primary](#input\_is\_primary) | Whether this is the primary node in an HA pair. If true, an alias IP (VIP) is allocated on the HA interface. | `bool` | `false` | no |
| <a name="input_labels"></a> [labels](#input\_labels) | A map of key/value labels to assign to the instance. | `map(string)` | <pre>{<br/>  "dontstop": "yes",<br/>  "dontterminate": "yes",<br/>  "product": "nios"<br/>}</pre> | no |
| <a name="input_lan1_subnet_name"></a> [lan1\_subnet\_name](#input\_lan1\_subnet\_name) | The name of the subnetwork to attach to the secondary network interface (nic1). | `string` | n/a | yes |
| <a name="input_machine_type"></a> [machine\_type](#input\_machine\_type) | The machine type to use for the instance. Used if nios\_model is not mapped. | `string` | `"n2-standard-4"` | no |
| <a name="input_mgmt_subnet_name"></a> [mgmt\_subnet\_name](#input\_mgmt\_subnet\_name) | The name of the subnetwork to attach to the primary network interface (nic0). | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | The name of the compute instance. | `string` | `"nios-gcp-instance"` | no |
| <a name="input_nios_license"></a> [nios\_license](#input\_nios\_license) | The NIOS license string applied during instance initialization. | `string` | `"nios IB-V1425 enterprise dns dhcp cloud"` | no |
| <a name="input_nios_model"></a> [nios\_model](#input\_nios\_model) | The NIOS appliance model used to determine the machine type. | `string` | `"IB-V1425"` | no |
| <a name="input_project"></a> [project](#input\_project) | The default project to manage resources in. | `string` | n/a | yes |
| <a name="input_region"></a> [region](#input\_region) | The region in which to manage resources. | `string` | `"us-west1"` | no |
| <a name="input_remote_console_enabled"></a> [remote\_console\_enabled](#input\_remote\_console\_enabled) | Whether to enable remote console access. | `bool` | `true` | no |
| <a name="input_service_account_email"></a> [service\_account\_email](#input\_service\_account\_email) | The service account e-mail address. | `string` | `null` | no |
| <a name="input_service_account_scopes"></a> [service\_account\_scopes](#input\_service\_account\_scopes) | A list of service scopes to assign to the service account. | `list(string)` | <pre>[<br/>  "https://www.googleapis.com/auth/cloud-platform"<br/>]</pre> | no |
| <a name="input_zone"></a> [zone](#input\_zone) | The zone in which the compute instance will be created. | `string` | `"us-west1-b"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_ha_gateway"></a> [ha\_gateway](#output\_ha\_gateway) | Gateway IP for the HA subnetwork. |
| <a name="output_ha_ip"></a> [ha\_ip](#output\_ha\_ip) | Internal IP of the HA interface (nic2). |
| <a name="output_ha_subnet_mask"></a> [ha\_subnet\_mask](#output\_ha\_subnet\_mask) | Subnet mask of the HA subnetwork. |
| <a name="output_instance_id"></a> [instance\_id](#output\_instance\_id) | ID of the NIOS Grid Member instance. |
| <a name="output_instance_name"></a> [instance\_name](#output\_instance\_name) | Name of the NIOS Grid Member instance. |
| <a name="output_lan1_gateway"></a> [lan1\_gateway](#output\_lan1\_gateway) | Gateway IP for the LAN1 subnetwork. |
| <a name="output_lan1_ip"></a> [lan1\_ip](#output\_lan1\_ip) | Internal IP of the LAN1 interface (nic1). |
| <a name="output_lan1_ipv6_address"></a> [lan1\_ipv6\_address](#output\_lan1\_ipv6\_address) | IPv6 address of the LAN1 interface (nic1). |
| <a name="output_lan1_subnet_mask"></a> [lan1\_subnet\_mask](#output\_lan1\_subnet\_mask) | Subnet mask of the LAN1 subnetwork. |
| <a name="output_mgmt_gateway"></a> [mgmt\_gateway](#output\_mgmt\_gateway) | Gateway IP for the MGMT subnetwork. |
| <a name="output_mgmt_ip"></a> [mgmt\_ip](#output\_mgmt\_ip) | Internal IP of the MGMT interface (nic0). |
| <a name="output_mgmt_ipv6_address"></a> [mgmt\_ipv6\_address](#output\_mgmt\_ipv6\_address) | IPv6 address of the MGMT interface (nic0). |
| <a name="output_mgmt_subnet_mask"></a> [mgmt\_subnet\_mask](#output\_mgmt\_subnet\_mask) | Subnet Mask of the Mgmt Subnetwork |
| <a name="output_vip"></a> [vip](#output\_vip) | Alias IP (floating VIP) on nic2 for HA. |
<!-- END_TF_DOCS -->

---

## Usage

### Step 1: Deploy GCP Infrastructure

```hcl
provider "google" {
  project = var.project
  region  = var.region
  zone    = var.zone
  credentials = file("path/to/service-account-key-file.json")
}

module "node1" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_gcp"
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_gcp"

  project = var.project
  region     = var.region
  zone       = var.zone

  image_name        = var.image_name
  name              = var.name
  nios_model        = var.nios_model
  mgmt_subnet_name  = var.mgmt_subnet_name
  lan1_subnet_name  = var.lan1_subnet_name

  boot_disk_type = var.boot_disk_type
  boot_disk_size = var.boot_disk_size

  nios_license           = var.nios_license
  default_admin_password = var.default_admin_password

  service_account_email  = var.service_account_email
  service_account_scopes = var.service_account_scopes

  labels = var.labels
}
```

**Deploy the infrastructure:**
```bash
terraform apply
```

### Step 2: Wait for NIOS to Boot

NIOS takes approximately around **30 minutes** to fully boot.

### Step 3: Join the Grid Member to the Master Grid

Once Grid is up and running, configure the grid member and join to the grid.

### Example 1: Join a Member to a Master

#### Deploy GCP infrastructure for Master and Member

```hcl
module "node1" {
  // ...(same config as Step 1)
}

module "node2" {
  // ... (same config as Step 1)
}
```

#### After both the grids are up and running (~30 min), configure grid member

```hcl
provider "nios" {
  nios_host_url = "https://${module.node1.lan1_ip}"
  nios_username = "username"
  nios_password = "password"
}

resource "nios_grid_member" "member" {
  host_name        = "infoblox.member"
  config_addr_type = "IPV4"
  platform         = "VNIOS"

  vip_setting = {
    address     = module.node2.lan1_ip
    gateway     = module.node2.lan1_gateway
    subnet_mask = module.node2.lan1_subnet_mask
  }
}

// Join member to existing grid master
resource "nios_grid_join" "member_join" {
  member_url      = "https://${module.node2.lan1_ip}"
  member_username = "Username"
  member_password = "Password"
  grid_name       = "Infoblox"
  master          = module.node1.lan1_ip
  shared_secret   = "<secret>"
  depends_on      = [nios_grid_member.member]
}
```

### Example 2: HA Grid Configuration

#### Deploy two GCP instances for SA-HA Config

> **Note:** HA configuration only supports IPv4. Dual-stack (IPv4 and IPv6) is not supported with HA on GCP.

> **Important - Service Account IAM Configuration:** HA formation fails when using predefined Google roles (e.g., `Compute Instance Admin (v1)`, `Compute Network Admin`, `Service Account User`) attached to the service account. To avoid this issue, create a **single custom IAM role** that combines all required permissions for both Terraform VM provisioning and HA operations. Attach only this custom role to the service account. This is a known limitation related to how NIOS validates IAM permissions with predefined Google roles.

```hcl
// Deploy GCP infrastructure for Node 1 (Active Node)
module "node1" {
  // ... (same config as Step 1)
  enable_ha         = true
  is_primary = true
  ha_subnet_name = "example-ha-subnet"
}

// Deploy GCP infrastructure for Node 2 (Passive Node)
module "node2" {
  // ... (same config as Step 1)
  enable_ha         = true
  ha_subnet_name = "example-ha-subnet"
}
```
#### After both the grids are up and running (~30 min), configure HA

1. Import Node1 under nios_grid_member.ha_pair

```hcl 
provider "nios" {
  nios_host_url = "https://${module.node1.lan1_ip}"
  nios_username = "username"
  nios_password = "password"
}

import {
  to = nios_grid_member.ha_pair
  id = "5c08e1293cf34363878d4cae5bd37636"
}
```

2. Modify the resource to set ha_on_cloud to true and provide the cloud attributes.

```
resource "nios_grid_member" "ha_pair" {
  host_name         = "infoblox.localdomain"
  config_addr_type  = "IPV4"
  platform          = "VNIOS"
  upgrade_group     = "Grid Master"
  master_candidate  = true

  enable_ha         = true
  router_id         = 100
  ha_on_cloud       = true
  ha_cloud_platform = "GCP"

  vip_setting = {
    address         = module.node1.vip
    gateway         = module.node1.ha_gateway
    subnet_mask     = module.node1.ha_subnet_mask
    lan1_gateway     = module.node1.lan1_gateway
    lan1_subnet_mask = module.node1.lan1_subnet_mask
    dscp            = 0
    primary         = true
    use_dscp        = false
  }

  node_info = [
    {
      lan_ha_port_setting = {
        ha_ip_address      = module.node1.ha_ip
        mgmt_lan           = module.node1.lan1_ip
        ha_cloud_attribute = module.node1.instance_name
      }
    },
    {
      lan_ha_port_setting = {
        ha_ip_address      = module.node2.ha_ip
        mgmt_lan           = module.node2.lan1_ip
        ha_cloud_attribute = module.node2.instance_name
      }
    }
  ]

  // To configure grid level dns resolver settings, use the 
  // grid_level_dns_resolver_setting attribute 
  
  grid_level_dns_resolver_setting = {
    resolvers = [
      "10.10.10.10"
  ] }
}
```

3. Join Node2 (Passive Node) to Node1 (Active Node).

```
provider "nios" {
  nios_host_url = "https://${module.node1.vip}"
  nios_username = "username"
  nios_password = "password"
}

resource "nios_grid_join" "ha_member_join" {
  member_url      = "https://${module.node2.lan1_ip}"
  member_username = "username"
  member_password = "password"
  grid_name       = "Infoblox"
  master          = module.node1.vip
  shared_secret   = "shared-secret"
}
```

#### Best Practices for HA Deployment

> **Recommended Workflow:** Use a **separate Terraform workspace** for HA configuration. The NIOS HA setup is a one-time provisioning task — once the HA pair is formed and the passive node has joined the grid, the configuration is complete and does not require ongoing Terraform management. Re-applying Terraform after HA formation may attempt to modify NIOS-managed HA attributes.

After successfully deploying the HA pair:

1. **Verify HA formation** is complete through the NIOS UI or API
2. **Remove the grid master from Terraform state** to prevent accidental modifications:
   ```bash
   terraform state rm nios_grid_member.ha_pair
   terraform state rm nios_grid_join.ha_member_join
   ```
3. Optionally, you can delete the entire Terraform state for this workspace if no further infrastructure management is needed.

This approach ensures that:
- Your HA infrastructure is provisioned correctly
- Subsequent Terraform operations don't interfere with the running HA pair
- The grid master configuration remains stable and is managed through NIOS directly

### Example 3: Join a Member to a Master with Dual Stack Config

#### Deploy GCP infrastructure for Master and Member

```hcl
module "node1" {
  // ...(same config as Step 1)
  enable_ipv6 = true
}

module "node2" {
  // ... (same config as Step 1)
  enable_ipv6 = true
}
```

#### After both the grids are up and running (~30 min), configure grid member

> **Note:** On GCP, IPv6 addresses are available in the Terraform state (`mgmt_ipv6_address` and `lan1_ipv6_address` outputs), but you must manually configure the IPv6 settings on the NIOS grid through the UI or API. Unlike AWS, GCP does not automatically configure IPv6 on the NIOS instance. This is a known limitation that will be addressed in a future NIOS release.

```hcl
provider "nios" {
  nios_host_url = "https://${module.node1.lan1_ip}"
  nios_username = "username"
  nios_password = "password"
}

resource "nios_grid_member" "member" {
  host_name        = "infoblox.member"
  config_addr_type = "BOTH"
  platform         = "VNIOS"
  vip_setting = {
    address     = module.node2.lan1_ip
    gateway     = module.node2.lan1_gateway
    subnet_mask = module.node2.lan1_subnet_mask
  }
  ipv6_setting = {
    virtual_ip  = module.node2.lan1_ipv6_address
    cidr_prefix = 64
    gateway     = "<ipv6_lan1_gateway>"
    enabled     = true
  }
}

// Join member to existing grid master
resource "nios_grid_join" "member_join" {
  member_url      = "https://${module.node2.lan1_ip}"
  member_username = "<username>"
  member_password = "<password>"
  grid_name       = "Infoblox"
  master          = module.node1.lan1_ip
  shared_secret   = "<secret>"
  depends_on      = [nios_grid_member.member]
}
```

### Boot Time
- NIOS takes **30 minutes** to fully boot after instance creation, make sure the grid is up and running before triggering the grid join.
- Always verify NIOS API is responding before applying `nios_grid_member` resources
