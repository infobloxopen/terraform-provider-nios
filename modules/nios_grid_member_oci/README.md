# NIOS Grid Member OCI Module

## Overview

This module provisions the OCI infrastructure for a NIOS Grid Member on Oracle Cloud Infrastructure.
It uploads a NIOS QCOW2 image to Object Storage, imports it as a custom compute image, and launches
a vNIOS instance with MGMT (primary) and LAN1 (secondary) VNICs. An optional reporting block volume
can also be attached.

The NIOS configuration (`nios_grid_member` and `nios_grid_join` resources) should be applied **after**
the infrastructure is deployed and NIOS is fully booted (~15–25 minutes).

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.8.0 |
| <a name="requirement_oci"></a> [oci](#requirement\_oci) | >= 5.0.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_oci"></a> [oci](#provider\_oci) | >= 5.0.0 |

## Resources

| Name | Type |
|------|------|
| [oci_core_image.nios_image](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/core_image) | resource |
| [oci_core_instance.nios_instance](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/core_instance) | resource |
| [oci_core_vnic_attachment.lan1_vnic_attachment](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/core_vnic_attachment) | resource |
| [oci_core_volume.reporting_volume](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/core_volume) | resource |
| [oci_core_volume_attachment.reporting_volume_attachment](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/core_volume_attachment) | resource |
| [oci_objectstorage_bucket.nios_bucket](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/objectstorage_bucket) | resource |
| [oci_objectstorage_object.nios_qcow2](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/objectstorage_object) | resource |
| [oci_core_subnet.lan1_subnet](https://registry.terraform.io/providers/oracle/oci/latest/docs/data-sources/core_subnet) | data source |
| [oci_core_vnic.lan1_vnic](https://registry.terraform.io/providers/oracle/oci/latest/docs/data-sources/core_vnic) | data source |
| [oci_objectstorage_namespace.ns](https://registry.terraform.io/providers/oracle/oci/latest/docs/data-sources/objectstorage_namespace) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_availability_domain"></a> [availability\_domain](#input\_availability\_domain) | Full availability domain name (e.g. Uocm:US-ASHBURN-AD-1). | `string` | n/a | yes |
| <a name="input_bucket_name"></a> [bucket\_name](#input\_bucket\_name) | Name of the Object Storage bucket for the NIOS QCOW2 image. | `string` | n/a | yes |
| <a name="input_cloud_init_content"></a> [cloud\_init\_content](#input\_cloud\_init\_content) | Inline cloud-init YAML. Takes precedence over cloud\_init\_script\_path. | `string` | `""` | no |
| <a name="input_cloud_init_script_path"></a> [cloud\_init\_script\_path](#input\_cloud\_init\_script\_path) | Path to a cloud-init YAML file. Leave empty to skip. | `string` | `""` | no |
| <a name="input_compartment_id"></a> [compartment\_id](#input\_compartment\_id) | OCID of the compartment in which all resources will be created. | `string` | n/a | yes |
| <a name="input_create_bucket"></a> [create\_bucket](#input\_create\_bucket) | Set to true to create a new bucket; false to reuse an existing one. | `bool` | `true` | no |
| <a name="input_default_admin_password"></a> [default\_admin\_password](#input\_default\_admin\_password) | Default admin password for NIOS. | `string` | n/a | yes |
| <a name="input_enable_reporting_volume"></a> [enable\_reporting\_volume](#input\_enable\_reporting\_volume) | Create and attach a reporting block volume to the Grid Member. | `bool` | `false` | no |
| <a name="input_fingerprint"></a> [fingerprint](#input\_fingerprint) | Fingerprint of the API signing key. | `string` | n/a | yes |
| <a name="input_image_name"></a> [image\_name](#input\_image\_name) | Display name for the custom OCI image imported from the QCOW2. | `string` | `"nios-custom-image"` | no |
| <a name="input_instance_memory_in_gbs"></a> [instance\_memory\_in\_gbs](#input\_instance\_memory\_in\_gbs) | Memory in GB — used only for IB-V5005. | `number` | `32` | no |
| <a name="input_instance_name"></a> [instance\_name](#input\_instance\_name) | Base display name. Instances will be named <name>-gm and <name>-member. | `string` | `"nios"` | no |
| <a name="input_instance_ocpus"></a> [instance\_ocpus](#input\_instance\_ocpus) | OCPUs — used only for IB-V5005. | `number` | `4` | no |
| <a name="input_lan1_assign_public_ip"></a> [lan1\_assign\_public\_ip](#input\_lan1\_assign\_public\_ip) | Assign a public IP to the LAN1 VNIC. | `bool` | `false` | no |
| <a name="input_lan1_subnet_id"></a> [lan1\_subnet\_id](#input\_lan1\_subnet\_id) | OCID of the subnet for the LAN1 interface. | `string` | n/a | yes |
| <a name="input_lan1_vnic_name"></a> [lan1\_vnic\_name](#input\_lan1\_vnic\_name) | Display name for the secondary (LAN1) VNIC. | `string` | `"nios-lan1-vnic"` | no |
| <a name="input_legacy_shape"></a> [legacy\_shape](#input\_legacy\_shape) | Fixed shape for NIOS < 9.0.2 (e.g. VM.Standard2.2). | `string` | `"VM.Standard2.2"` | no |
| <a name="input_mgmt_assign_public_ip"></a> [mgmt\_assign\_public\_ip](#input\_mgmt\_assign\_public\_ip) | Assign a public IP to the MGMT VNIC. | `bool` | `false` | no |
| <a name="input_mgmt_subnet_id"></a> [mgmt\_subnet\_id](#input\_mgmt\_subnet\_id) | OCID of the subnet for the MGMT interface. | `string` | n/a | yes |
| <a name="input_mgmt_vnic_name"></a> [mgmt\_vnic\_name](#input\_mgmt\_vnic\_name) | Display name for the primary (MGMT) VNIC. | `string` | `"nios-mgmt-vnic"` | no |
| <a name="input_nios_license"></a> [nios\_license](#input\_nios\_license) | NIOS temporary license string. | `string` | `"nios IB-V825 enterprise dns dhcp cloud"` | no |
| <a name="input_nios_model"></a> [nios\_model](#input\_nios\_model) | NIOS appliance model — sets OCPUs and memory for Flex shape.<br/>One of: IB-V926, IB-V1516, IB-V1526, IB-V2326, IB-V4126, IB-V5005. | `string` | `"IB-V926"` | no |
| <a name="input_nios_object_name"></a> [nios\_object\_name](#input\_nios\_object\_name) | Object name to store the QCOW2 as in the bucket. | `string` | n/a | yes |
| <a name="input_nios_qcow2_local_path"></a> [nios\_qcow2\_local\_path](#input\_nios\_qcow2\_local\_path) | Absolute local path to the NIOS QCOW2 image file. | `string` | n/a | yes |
| <a name="input_nios_version_gte_902"></a> [nios\_version\_gte\_902](#input\_nios\_version\_gte\_902) | true → VM.Standard3.Flex (NIOS >= 9.0.2). false → legacy\_shape. | `bool` | `true` | no |
| <a name="input_private_key_path"></a> [private\_key\_path](#input\_private\_key\_path) | Absolute local path to the OCI API private key (PEM file). | `string` | n/a | yes |
| <a name="input_region"></a> [region](#input\_region) | OCI region identifier (e.g. us-ashburn-1). | `string` | n/a | yes |
| <a name="input_remote_console_enabled"></a> [remote\_console\_enabled](#input\_remote\_console\_enabled) | Enable remote console access. | `bool` | `true` | no |
| <a name="input_reporting_volume_name"></a> [reporting\_volume\_name](#input\_reporting\_volume\_name) | Display name for the reporting block volume. | `string` | `"nios-reporting-volume"` | no |
| <a name="input_reporting_volume_size_gb"></a> [reporting\_volume\_size\_gb](#input\_reporting\_volume\_size\_gb) | Size in GB for the reporting volume. Minimum 250 GB recommended. | `number` | `250` | no |
| <a name="input_tenancy_ocid"></a> [tenancy\_ocid](#input\_tenancy\_ocid) | OCID of your OCI tenancy. | `string` | n/a | yes |
| <a name="input_user_ocid"></a> [user\_ocid](#input\_user\_ocid) | OCID of the OCI IAM user used for API authentication. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_instance_id"></a> [instance\_id](#output\_instance\_id) | OCID of the NIOS compute instance. |
| <a name="output_lan1_gateway"></a> [lan1\_gateway](#output\_lan1\_gateway) | Gateway IP of the LAN1 subnet (OCI virtual router IP). |
| <a name="output_lan1_private_ip"></a> [lan1\_private\_ip](#output\_lan1\_private\_ip) | Private IP address of the secondary (LAN1) VNIC. |
| <a name="output_lan1_subnet_mask"></a> [lan1\_subnet\_mask](#output\_lan1\_subnet\_mask) | Subnet mask of the LAN1 subnet (e.g. 255.255.255.0). |
| <a name="output_lan1_vnic_id"></a> [lan1\_vnic\_id](#output\_lan1\_vnic\_id) | OCID of the secondary (LAN1) VNIC. |
<!-- END_TF_DOCS -->

---

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│  OCI Compartment                                            │
│                                                             │
│  ┌──────────────────────────────────────────────────────┐  │
│  │  Object Storage Bucket                               │  │
│  │  └── NIOS QCOW2 Object                               │  │
│  └──────────────────────────────────────────────────────┘  │
│               │ import (30–60 min)                          │
│  ┌────────────▼─────────────────────────────────────────┐  │
│  │  Custom Image  (PARAVIRTUALIZED launch mode)         │  │
│  └──────────────────────────────────────────────────────┘  │
│               │ boot                                        │
│  ┌────────────▼─────────────────────────────────────────┐  │
│  │  Compute Instance  (VM.Standard3.Flex / legacy)      │  │
│  │  ├── Primary VNIC    (MGMT  / eth0 equivalent)       │  │
│  │  ├── Secondary VNIC  (LAN1  / eth1 equivalent)       │  │
│  │  └── Block Volume    (Reporting — optional)          │  │
│  └──────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
```

## NIOS Model → Shape Configuration

| NIOS Model | OCPUs | Memory (GB) | Notes |
|---|---|---|---|
| IB-V926  | 4  | 32  | NIOS >= 9.0.2 |
| IB-V1516 | 6  | 64  | NIOS >= 9.0.2 |
| IB-V1526 | 8  | 64  | NIOS >= 9.0.2 |
| IB-V2326 | 10 | 192 | NIOS >= 9.0.2 |
| IB-V4126 | 16 | 384 | NIOS >= 9.0.2 |
| IB-V5005 | user-defined | user-defined | NIOS >= 9.0.6 (reporting appliance) |

> For **IB-V5005** set `instance_ocpus` and `instance_memory_in_gbs` in your `tfvars`.  
> For **NIOS < 9.0.2** set `nios_version_gte_902 = false` and provide `legacy_shape`.

## Deployment Order

| Step | Resource | Notes |
|---|---|---|
| 1 | Object Storage Bucket | Skipped when `create_bucket = false` |
| 2 | NIOS QCOW2 Upload | Streamed to bucket; use OCI CLI for files > 5 GB |
| 3 | Custom Image import | Takes 30–60 min; timeout set to 60 min |
| 4 | Compute Instance | Primary VNIC = MGMT (eth0) |
| 5 | Secondary VNIC (LAN1) | Attached after instance reaches running state |
| 6 | Reporting Block Volume | Skipped when `enable_reporting_volume = false` |

## Usage

### Step 1: Deploy OCI Infrastructure

```hcl
module "nios_grid_member" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_grid_member_oci"

  # OCI authentication
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region

  # Compartment
  compartment_id = var.compartment_id

  # Image
  create_bucket         = true
  bucket_name           = "nios-images"
  nios_object_name      = "nios-9.0.2.qcow2"
  nios_qcow2_local_path = "/path/to/nios-9.0.2.qcow2"
  image_name            = "nios-9.0.2"

  # Instance
  instance_name        = "nios-grid-member-1"
  availability_domain  = "Uocm:PHX-AD-1"
  nios_model           = "IB-V1526"
  nios_version_gte_902 = true

  # Networking
  mgmt_subnet_id        = var.mgmt_subnet_id
  lan1_subnet_id        = var.lan1_subnet_id
  mgmt_assign_public_ip = true
  lan1_assign_public_ip = false

  # Reporting volume (optional)
  enable_reporting_volume  = true
  reporting_volume_size_gb = 250
}
```

**Deploy the infrastructure:**
```bash
terraform apply
```

### Step 2: Wait for NIOS to Boot

NIOS takes approximately **15–25 minutes** to fully boot after the instance starts.

### Step 3: Join the Grid Member to the Master Grid

Once Grid is up and running, configure the grid member and join to the grid.

#### Examples

#### Example 1: Join a Member to a Master

#### Deploy AWS infrastructure for Master and Member

```hcl
module "node1" {
  // ... (same config as Step 1)
  ha_enable = false
}

module "node2" {
  // ... (same config as Step 1)
  ha_enable = false
}

// After NIOS is ready (~20 min), configure grid member
provider "nios" {
  nios_host_url = "https://${module.nios_grid_member.mgmt_private_ip}"
  nios_username = "admin"
  nios_password = var.nios_password
}

resource "nios_grid_member" "member" {
  host_name        = "infoblox.member"
  config_addr_type = "IPV4"
  platform         = "VNIOS"

  vip_setting = {
    address     = module.nios_grid_member.lan1_private_ip
    gateway     = "<lan1_gateway_ip>"
    subnet_mask = "<lan1_subnet_mask>"
  }
}

resource "nios_grid_join" "member_join" {
  member_url      = "https://${module.nios_grid_member.lan1_private_ip}"
  member_username = "admin"
  member_password = var.nios_password
  grid_name       = "Infoblox"
  master          = "<master_ip>"
  shared_secret   = var.shared_secret
  depends_on      = [nios_grid_member.member]
}
```

## Outputs Usage

| Output | NIOS Resource Usage |
|---|---|
| `mgmt_private_ip` | Provider `nios_host_url`, SSH access |
| `mgmt_public_ip` | Public access (when `mgmt_assign_public_ip = true`) |
| `lan1_private_ip` | `vip_setting.address`, `master` in `nios_grid_join` |
| `lan1_vnic_id` | OCI-level VNIC operations |
| `nios_image_id` | Reference for additional instances using the same image |
| `reporting_volume_id` | OCI-level volume management |

## Large Image Upload (> 5 GB)

For large QCOW2 files, use the OCI CLI for better performance:

```bash
oci os object bulk-upload \
  --namespace <namespace> \
  --bucket-name <bucket> \
  --src-dir /path/to/dir \
  --include "*.qcow2"
```

Then remove the upload resource from Terraform state to prevent re-upload:

```bash
terraform state rm oci_objectstorage_object.nios_qcow2
```

---

### Boot Time
- NIOS takes **15–25 minutes** to fully boot after instance creation
- Always verify the NIOS API is responding before applying `nios_grid_member` resources

### Image Import
- Custom image import can take **30–60 minutes**
- The `create` timeout is set to `60m` to accommodate this
- Launch mode is set to `PARAVIRTUALIZED` as required by Infoblox

### Cloud-Init
- Inline content (`cloud_init_content`) takes precedence over file path (`cloud_init_script_path`)
- Changes to cloud-init after initial deployment are ignored (`lifecycle.ignore_changes = [metadata]`)

### Reporting Volume
- Minimum **250 GB** recommended by Infoblox
- Encrypted with Oracle-managed keys (default OCI behavior)
- Attached as Paravirtualized, Read/Write