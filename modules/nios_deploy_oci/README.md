# Deploy vNIOS on OCI

## Overview

This module provisions vNIOS on Oracle Cloud Infrastructure.
It uploads a NIOS QCOW2 image to Object Storage, imports it as a custom compute image, and launches
a vNIOS instance with MGMT (primary) and LAN1 (secondary) VNICs. An optional reporting block volume
can also be attached.

The NIOS configuration (`nios_grid_member` and `nios_grid_join` resources) should be applied **after**
the infrastructure is deployed and NIOS is fully booted (~30 minutes).

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.12.1 |
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
| [oci_core_private_ip.ha_vip](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/core_private_ip) | resource |
| [oci_core_vnic_attachment.ha_vnic_attachment](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/core_vnic_attachment) | resource |
| [oci_core_vnic_attachment.lan1_vnic_attachment](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/core_vnic_attachment) | resource |
| [oci_core_volume.reporting_volume](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/core_volume) | resource |
| [oci_core_volume_attachment.reporting_volume_attachment](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/core_volume_attachment) | resource |
| [oci_objectstorage_bucket.nios_bucket](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/objectstorage_bucket) | resource |
| [oci_objectstorage_object.nios_qcow2](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/objectstorage_object) | resource |
| [oci_core_subnet.ha_subnet](https://registry.terraform.io/providers/oracle/oci/latest/docs/data-sources/core_subnet) | data source |
| [oci_core_subnet.lan1_subnet](https://registry.terraform.io/providers/oracle/oci/latest/docs/data-sources/core_subnet) | data source |
| [oci_core_subnet.mgmt_subnet](https://registry.terraform.io/providers/oracle/oci/latest/docs/data-sources/core_subnet) | data source |
| [oci_core_vnic.ha_vnic](https://registry.terraform.io/providers/oracle/oci/latest/docs/data-sources/core_vnic) | data source |
| [oci_core_vnic.lan1_vnic](https://registry.terraform.io/providers/oracle/oci/latest/docs/data-sources/core_vnic) | data source |
| [oci_core_vnic.mgmt_vnic](https://registry.terraform.io/providers/oracle/oci/latest/docs/data-sources/core_vnic) | data source |
| [oci_core_vnic_attachments.mgmt_vnic_attachments](https://registry.terraform.io/providers/oracle/oci/latest/docs/data-sources/core_vnic_attachments) | data source |
| [oci_objectstorage_namespace.ns](https://registry.terraform.io/providers/oracle/oci/latest/docs/data-sources/objectstorage_namespace) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_availability_domain"></a> [availability\_domain](#input\_availability\_domain) | Full availability domain name (e.g. Uocm:US-ASHBURN-AD-1). | `string` | n/a | yes |
| <a name="input_bucket_name"></a> [bucket\_name](#input\_bucket\_name) | Name of the Object Storage bucket for the NIOS QCOW2 image. Required when create\_image = true; must not be set when create\_image = false. | `string` | `null` | no |
| <a name="input_compartment_id"></a> [compartment\_id](#input\_compartment\_id) | OCID of the compartment in which all resources will be created. | `string` | n/a | yes |
| <a name="input_create_bucket"></a> [create\_bucket](#input\_create\_bucket) | Set to true to create a new bucket; false to reuse an existing one. Only used when create\_image = true. | `bool` | `true` | no |
| <a name="input_create_image"></a> [create\_image](#input\_create\_image) | If true, the module uploads the NIOS QCOW2 to Object Storage and imports it as a custom image. If false (default), the module uses an existing image via var.image\_id. | `bool` | `false` | no |
| <a name="input_default_admin_password"></a> [default\_admin\_password](#input\_default\_admin\_password) | Default admin password for NIOS. | `string` | n/a | yes |
| <a name="input_enable_ha"></a> [enable\_ha](#input\_enable\_ha) | Enable High Availability configuration (adds HA VNIC). | `bool` | `false` | no |
| <a name="input_enable_reporting_volume"></a> [enable\_reporting\_volume](#input\_enable\_reporting\_volume) | Create and attach a reporting block volume to the Grid Member. | `bool` | `false` | no |
| <a name="input_freeform_tags"></a> [freeform\_tags](#input\_freeform\_tags) | A map of key/value freeform tags to assign to the instance. | `map(string)` | <pre>{<br/>  "dontstop": "yes",<br/>  "dontterminate": "yes",<br/>  "product": "nios"<br/>}</pre> | no |
| <a name="input_ha_assign_public_ip"></a> [ha\_assign\_public\_ip](#input\_ha\_assign\_public\_ip) | Assign a public IP to the HA VNIC. | `bool` | `false` | no |
| <a name="input_ha_subnet_id"></a> [ha\_subnet\_id](#input\_ha\_subnet\_id) | OCID of the subnet for the HA interface. Required when enable\_ha = true. | `string` | `null` | no |
| <a name="input_ha_vnic_name"></a> [ha\_vnic\_name](#input\_ha\_vnic\_name) | Display name for the HA VNIC. | `string` | `"nios-ha-vnic"` | no |
| <a name="input_image_id"></a> [image\_id](#input\_image\_id) | OCID of an existing NIOS custom image to use for instance creation. Required when create\_image = false; must not be set when create\_image = true. | `string` | `null` | no |
| <a name="input_image_name"></a> [image\_name](#input\_image\_name) | Display name for the custom OCI image imported from the QCOW2. Required when create\_image = true; must not be set when create\_image = false. | `string` | `null` | no |
| <a name="input_instance_memory_in_gbs"></a> [instance\_memory\_in\_gbs](#input\_instance\_memory\_in\_gbs) | Memory in GB — used only for IB-V5005. | `number` | `32` | no |
| <a name="input_instance_name"></a> [instance\_name](#input\_instance\_name) | Display name for the OCI instance. | `string` | `"nios"` | no |
| <a name="input_instance_ocpus"></a> [instance\_ocpus](#input\_instance\_ocpus) | OCPUs — used only for IB-V5005. | `number` | `4` | no |
| <a name="input_is_primary"></a> [is\_primary](#input\_is\_primary) | True for the primary node in an HA pair. It has the VIP assigned to its HA VNIC. Set to false for the secondary node. | `bool` | `false` | no |
| <a name="input_lan1_assign_public_ip"></a> [lan1\_assign\_public\_ip](#input\_lan1\_assign\_public\_ip) | Assign a public IP to the LAN1 VNIC. | `bool` | `false` | no |
| <a name="input_lan1_subnet_id"></a> [lan1\_subnet\_id](#input\_lan1\_subnet\_id) | OCID of the subnet for the LAN1 interface. | `string` | n/a | yes |
| <a name="input_lan1_vnic_name"></a> [lan1\_vnic\_name](#input\_lan1\_vnic\_name) | Display name for the secondary (LAN1) VNIC. | `string` | `"nios-lan1-vnic"` | no |
| <a name="input_legacy_shape"></a> [legacy\_shape](#input\_legacy\_shape) | Fixed shape for NIOS < 9.x.x (e.g. VM.Standard2.2). | `string` | `"VM.Standard2.2"` | no |
| <a name="input_mgmt_assign_public_ip"></a> [mgmt\_assign\_public\_ip](#input\_mgmt\_assign\_public\_ip) | Assign a public IP to the MGMT VNIC. | `bool` | `false` | no |
| <a name="input_mgmt_subnet_id"></a> [mgmt\_subnet\_id](#input\_mgmt\_subnet\_id) | OCID of the subnet for the MGMT interface. | `string` | n/a | yes |
| <a name="input_mgmt_vnic_name"></a> [mgmt\_vnic\_name](#input\_mgmt\_vnic\_name) | Display name for the primary (MGMT) VNIC. | `string` | `"nios-mgmt-vnic"` | no |
| <a name="input_nios_license"></a> [nios\_license](#input\_nios\_license) | NIOS temporary license string. | `string` | `"nios IB-V825 enterprise dns dhcp cloud"` | no |
| <a name="input_nios_model"></a> [nios\_model](#input\_nios\_model) | NIOS appliance model — sets OCPUs and memory for Flex shape.<br/>One of: IB-V926, IB-V1516, IB-V1526, IB-V2326, IB-V4126, IB-V5005. | `string` | `"IB-V926"` | no |
| <a name="input_nios_object_name"></a> [nios\_object\_name](#input\_nios\_object\_name) | Object name to store the QCOW2 as in the bucket. Required when create\_image = true; must not be set when create\_image = false. | `string` | `null` | no |
| <a name="input_nios_qcow2_local_path"></a> [nios\_qcow2\_local\_path](#input\_nios\_qcow2\_local\_path) | Absolute local path to the NIOS QCOW2 image file. Required when create\_image = true; must not be set when create\_image = false. | `string` | `null` | no |
| <a name="input_nios_version_gte_9xx"></a> [nios\_version\_gte\_9xx](#input\_nios\_version\_gte\_9xx) | true → VM.Standard3.Flex (NIOS >= 9.x.x). false → legacy\_shape. | `bool` | `true` | no |
| <a name="input_remote_console_enabled"></a> [remote\_console\_enabled](#input\_remote\_console\_enabled) | Enable remote console access. | `bool` | `true` | no |
| <a name="input_reporting_volume_name"></a> [reporting\_volume\_name](#input\_reporting\_volume\_name) | Display name for the reporting block volume. | `string` | `"nios-reporting-volume"` | no |
| <a name="input_reporting_volume_size_gb"></a> [reporting\_volume\_size\_gb](#input\_reporting\_volume\_size\_gb) | Size in GB for the reporting volume. Minimum 250 GB recommended. | `number` | `250` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_ha_gateway"></a> [ha\_gateway](#output\_ha\_gateway) | Gateway IP of the HA subnet. |
| <a name="output_ha_ip"></a> [ha\_ip](#output\_ha\_ip) | Private IP of the HA VNIC. |
| <a name="output_ha_subnet_mask"></a> [ha\_subnet\_mask](#output\_ha\_subnet\_mask) | Subnet mask of the HA subnet. |
| <a name="output_ha_vnic_id"></a> [ha\_vnic\_id](#output\_ha\_vnic\_id) | OCID of the HA VNIC. |
| <a name="output_instance_id"></a> [instance\_id](#output\_instance\_id) | OCID of the NIOS compute instance. |
| <a name="output_lan1_gateway"></a> [lan1\_gateway](#output\_lan1\_gateway) | Gateway IP of the LAN1 subnet (OCI virtual router IP). |
| <a name="output_lan1_private_ip"></a> [lan1\_private\_ip](#output\_lan1\_private\_ip) | Private IP address of the secondary (LAN1) VNIC. |
| <a name="output_lan1_subnet_mask"></a> [lan1\_subnet\_mask](#output\_lan1\_subnet\_mask) | Subnet mask of the LAN1 subnet (e.g. 255.255.255.0). |
| <a name="output_lan1_vnic_id"></a> [lan1\_vnic\_id](#output\_lan1\_vnic\_id) | OCID of the secondary (LAN1) VNIC. |
| <a name="output_mgmt_gateway"></a> [mgmt\_gateway](#output\_mgmt\_gateway) | Gateway IP of the MGMT subnet. |
| <a name="output_mgmt_ip"></a> [mgmt\_ip](#output\_mgmt\_ip) | Private IP of the primary (MGMT) VNIC. |
| <a name="output_mgmt_subnet_mask"></a> [mgmt\_subnet\_mask](#output\_mgmt\_subnet\_mask) | Subnet mask of the MGMT subnet. |
| <a name="output_vip"></a> [vip](#output\_vip) | Virtual IP (VIP) for HA - floats between primary/secondary. |
<!-- END_TF_DOCS -->

---

## Deployment Order

| Step | Resource | Notes |
|---|---|---|
| 1 | Object Storage Bucket | Created only when `create_image = true` and `create_bucket = true` |
| 2 | NIOS QCOW2 Upload | Created only when `create_image = true`; use OCI CLI for files > 5 GB |
| 3 | Custom Image import | Created only when `create_image = true`; takes 30–60 min; timeout set to 60 min |
| 4 | Compute Instance | Primary VNIC = MGMT (eth0). Uses `image_id` when `create_image = false`, else the imported image |
| 5 | Secondary VNIC (LAN1) | Attached after instance reaches running state |
| 6 | HA VNIC + VIP | Created only when `enable_ha = true`; VIP only on primary (`is_primary = true`) |
| 7 | Reporting Block Volume | Skipped when `enable_reporting_volume = false` |

## Usage

### Step 1: Deploy OCI Infrastructure

The module supports two mutually exclusive image-source modes, controlled by `create_image`:

| Mode | `create_image` | Required inputs | Forbidden inputs |
|---|---|---|---|
| **Use existing image** (default) | `false` | `image_id` | `bucket_name`, `nios_qcow2_local_path`, `nios_object_name`, `image_name` |
| **Create image from QCOW2** | `true` | `bucket_name`, `nios_qcow2_local_path`, `nios_object_name`, `image_name` | `image_id` |

When `create_image = true`, the module uploads the QCOW2 to Object Storage and imports it as a custom image. Set `create_bucket = false` to reuse a pre-existing bucket (only valid when`create_image = true`).

#### Option A — Use an existing custom image (default, `create_image = false`)

```hcl
provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

module "node1" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_oci?ref=nios_v9.1.0"

  default_admin_password = var.default_admin_password
  remote_console_enabled = var.remote_console_enabled
  nios_license           = var.nios_license

  # Compartment
  compartment_id = var.compartment_id

  # Image (existing custom image)
  image_id = var.image_id

  # Instance
  instance_name          = var.instance_name
  availability_domain    = var.availability_domain
  nios_model             = var.nios_model
  nios_version_gte_9xx   = var.nios_version_gte_9xx
  legacy_shape           = var.legacy_shape
  instance_ocpus         = var.instance_ocpus
  instance_memory_in_gbs = var.instance_memory_in_gbs

  # Networking
  mgmt_subnet_id        = var.mgmt_subnet_id
  mgmt_assign_public_ip = var.mgmt_assign_public_ip
  lan1_subnet_id        = var.lan1_subnet_id
  lan1_assign_public_ip = var.lan1_assign_public_ip

  # Reporting volume (optional)
  enable_reporting_volume  = var.enable_reporting_volume
  reporting_volume_name    = var.reporting_volume_name
  reporting_volume_size_gb = var.reporting_volume_size_gb
}
```

#### Option B — Upload QCOW2 and import as custom image (`create_image = true`)

```hcl
module "node1" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_oci?ref=nios_v9.1.0"

  default_admin_password = var.default_admin_password
  remote_console_enabled = var.remote_console_enabled
  nios_license           = var.nios_license

  # Compartment
  compartment_id = var.compartment_id

  # Image (upload QCOW2)
  create_image          = true
  create_bucket         = true # set false to reuse an existing bucket
  bucket_name           = var.bucket_name
  nios_qcow2_local_path = var.nios_qcow2_local_path
  nios_object_name      = var.nios_object_name
  image_name            = var.image_name

  instance_name          = var.instance_name
  availability_domain    = var.availability_domain
  nios_model             = var.nios_model
  mgmt_subnet_id         = var.mgmt_subnet_id
  lan1_subnet_id         = var.lan1_subnet_id
}
```

**Deploy the infrastructure:**
```bash
terraform apply
```

### Step 2: Wait for NIOS to Boot

NIOS takes approximately **30 minutes** to fully boot after the instance starts.

### Step 3: Join the Grid Member to the Master Grid

Once Grid is up and running, configure the grid member and join to the grid.

#### Examples

#### Example 1: Join a Member to a Master

#### Deploy OCI infrastructure for Master and Member

```hcl
module "node1" {
  // ... (same config as Step 1)
}

module "node2" {
  // ... (same config as Step 1)
}

// After NIOS is ready (~30 mins), configure grid member
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

#### Example 2: HA Grid Configuration

Deploy two OCI instances for SA-HA Config

```hcl
// Deploy OCI instance for Node 1 (Active Node)
module "node1" {
  // ... (same config as Step 1)
  enable_ha    = true
  ha_subnet_id = var.ha_subnet_id
  is_primary   = true
}

// Deploy OCI instance for Node 2 (Passive Node)
module "node2" {
  // ... (same config as Step 1)
  enable_ha    = true
  ha_subnet_id = var.ha_subnet_id
  is_primary   = false
}

// Create the OCI IAM resources required for NIOS HA VIP failover
module "nios_ha_iam" {
  source               = "../../../modules/nios_oci_ha_iam"
  tenancy_ocid         = var.tenancy_ocid
  compartment_id       = var.compartment_id
  idcs_endpoint        = var.idcs_endpoint
  identity_domain_name = var.identity_domain_name
  dynamic_group_name   = var.dynamic_group_name

  // Pass instance OCIDs
  instance_ocids = [
    module.node1.instance_id,
    module.node2.instance_id,
  ]
  depends_on = [module.node1, module.node2]
}
```

#### Why the `nios_oci_ha_iam` module is needed

For vNIOS HA on OCI, both nodes must function as **OCI principals** so they
can make the API calls required during failover — primarily to move the
Virtual IP (VIP) between nodes by reassigning a secondary private IP. This
is achieved by placing the HA instances into a **Dynamic Group** and
attaching IAM policies that grant the dynamic group the required
permissions on networking and identity resources.

The [`nios_oci_ha_iam`](../nios_oci_ha_iam) module automates this setup:

1. A **Dynamic Resource Group** in your Identity Domain whose matching
   rule references the NIOS HA instance OCIDs you pass in.

2. A **sub-compartment-level IAM policy** granting the dynamic group permissions to manage network interfaces, assign/unassign secondary IPs, and perform other OCI operations required for automated HA failover.

3. A **tenancy-level IAM policy** minimum required tenancy‑level permissions that must be granted to the NIOS dynamic group to support identity validation and authorization checks essential for HA workflows.

A single instance of this module can serve **multiple HA pairs** in the
same compartment — add each new pair's instance OCIDs to `instance_ocids`
and re-apply; the matching rule is updated in place.

For the full background and the complete list of policy statements, see:

- [Prerequisites for vNIOS for OCI HA](https://docs.infoblox.com/space/vniosoci/2188214440/Prerequisites+for+vNIOS+for+OCI+HA)
- [Deploying vNIOS for OCI in an HA Environment](https://docs.infoblox.com/space/vniosoci/2178056272/Deploying+vNIOS+for+OCI+in+an+HA+Environment)

> **If you choose not to use this module**, you must perform the
> equivalent setup manually in the OCI Console (or via your own
> Terraform/CLI) as described in the *Prerequisites for vNIOS for OCI HA*
> document above: create a dynamic group whose matching rule references
> the OCID of every NIOS HA instance, then attach the sub-compartment and
> tenancy-level IAM policies listed there.

#### After both nodes are up and running (~30 min), configure HA

> **Important — MGMT interface usage.** On a freshly booted NIOS instance
> the MGMT interface (`eth0`) is not yet enabled, so the **first** API call
> (the import in Step 1 below) must go through the **LAN1** interface — that
> is why the first `provider "nios"` block targets `module.node1.lan1_private_ip`.
> As part of that same step we enable `mgmt_port_setting` on the grid member
> and configure the grid-level DNS resolver  **All subsequent API calls** (Step 2
> onwards) are made over the MGMT interface, so those `provider "nios"`
> blocks target `module.node1.mgmt_ip`.

1. Import Node1 under nios_grid_member.ha_pair

```
provider "nios" {
  nios_host_url = "https://${module.node1.lan1_private_ip}"
  nios_username = "admin"
  nios_password = "infoblox"
}

import {
  to = nios_grid_member.example_ha
  id = "1a1915890950470093f7d3484b5d44a7"
}

resource "nios_grid_member" "ha_pair" {
  config_addr_type = "IPV4"
  host_name        = "infoblox.localdomain"
  platform         = "VNIOS"

  upgrade_group    = "Grid Master"
  master_candidate = true
  mgmt_port_setting = {
    enabled                 = true
    security_access_enabled = false
    vpn_enabled             = false
  }

  vip_setting = {
    address     = module.node1.lan1_private_ip
    gateway     = module.node1.lan1_gateway
    subnet_mask = module.node1.lan1_subnet_mask
  }

  node_info = [
    {
      mgmt_network_setting = {
        address     = module.node1.mgmt_ip
        gateway     = module.node1.mgmt_gateway
        subnet_mask = module.node1.mgmt_subnet_mask
      }
    },
  ]

  grid_level_dns_resolver_setting = {
    resolvers = [
      "10.103.3.10"
  ] }
}
```

Run Terraform Apply for the resource to be imported and
modified to enable mgmt settings

2. Modify the resource to set ha_on_cloud to true and provide the cloud attributes.

```
provider "nios" {
  nios_host_url = "https://${module.node1.mgmt_ip}"
  nios_username = "admin"
  nios_password = "infoblox"
}

resource "nios_grid_member" "ha_pair" {
  config_addr_type = "IPV4"
  host_name        = "infoblox.localdomain"
  platform         = "VNIOS"

  upgrade_group    = "Grid Master"
  master_candidate = true
  mgmt_port_setting = {
    enabled                 = true
    security_access_enabled = false
    vpn_enabled             = false
  }
    vip_setting = {
    address          = module.node1.vip
    gateway          = module.node1.ha_gateway
    subnet_mask      = module.node1.ha_subnet_mask
    lan1_gateway     = module.node1.lan1_gateway
    lan1_subnet_mask = module.node1.lan1_subnet_mask
  }

  enable_ha         = true
  ha_on_cloud       = true
  ha_cloud_platform = "OCI"

  node_info = [
    {
      lan_ha_port_setting = {
        ha_ip_address      = module.node1.ha_ip
        mgmt_lan           = module.node1.lan1_private_ip
        ha_cloud_attribute = module.node1.ha_vnic_id
      }
      mgmt_network_setting = {
        address     = module.node1.mgmt_ip
        gateway     = module.node1.mgmt_gateway
        subnet_mask = module.node1.mgmt_subnet_mask
      }
    },
    {
      lan_ha_port_setting = {
        ha_ip_address      = module.node2.ha_ip
        mgmt_lan           = module.node2.lan1_private_ip
        ha_cloud_attribute = module.node2.ha_vnic_id
      }
      mgmt_network_setting = {
        address     = module.node2.mgmt_ip
        gateway     = module.node2.mgmt_gateway
        subnet_mask = module.node2.mgmt_subnet_mask
      }
    }
  ]
  router_id = 111

  grid_level_dns_resolver_setting = {
    resolvers = [
      "10.103.3.10"
  ] }
}

```

Run terraform apply to update the resource 

3. Join Node2 (Passive Node) to Node1 (Active Node).

```
resource "nios_grid_join" "ha_member_join" {
  member_url      = "https://${module.node2.lan1_private_ip}"
  member_username = "admin"
  member_password = "infoblox"
  grid_name       = "Infoblox"
  master          = module.node1.vip
  shared_secret   = "test"
}
```


### Boot Time
- NIOS takes around **30 minutes** to fully boot after instance creation, make sure the grid is up and running before triggering grid join.
- Always verify the NIOS API is responding before applying `nios_grid_member` resources
