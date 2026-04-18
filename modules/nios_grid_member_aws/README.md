# NIOS Grid Member AWS Module

## Overview

This module provisions the AWS infrastructure (EC2 instance, network interfaces, etc.) for NIOS Grid. The NIOS configuration (`nios_grid_member` and `nios_grid_join` resources) should be applied after the infrastructure is deployed and NIOS grid is fully booted (~15-25 minutes).

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 6.38.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | >= 6.38.0 |

## Resources

| Name | Type |
|------|------|
| [aws_instance.grid](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance) | resource |
| [aws_network_interface.eth1](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/network_interface) | resource |
| [aws_network_interface.eth2](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/network_interface) | resource |
| [aws_subnet.lan1](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/subnet) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_ami_id"></a> [ami\_id](#input\_ami\_id) | AMI ID for NIOS instance. | `string` | `null` | no |
| <a name="input_availability_zone"></a> [availability\_zone](#input\_availability\_zone) | AWS availability zone. | `string` | `"us-west-1a"` | no |
| <a name="input_default_admin_password"></a> [default\_admin\_password](#input\_default\_admin\_password) | Default admin password for NIOS. | `string` | n/a | yes |
| <a name="input_delete_on_termination"></a> [delete\_on\_termination](#input\_delete\_on\_termination) | Whether to delete the volume on instance termination. | `bool` | `true` | no |
| <a name="input_enable_ipv6"></a> [enable\_ipv6](#input\_enable\_ipv6) | Enable IPv6 configuration. | `bool` | `false` | no |
| <a name="input_ha_enable"></a> [ha\_enable](#input\_ha\_enable) | Enable HA configuration. | `bool` | `false` | no |
| <a name="input_iam_instance_profile"></a> [iam\_instance\_profile](#input\_iam\_instance\_profile) | IAM instance profile to attach to the instance. | `string` | `null` | no |
| <a name="input_instance_type"></a> [instance\_type](#input\_instance\_type) | EC2 instance type. | `string` | `"r6i.large"` | no |
| <a name="input_key_name"></a> [key\_name](#input\_key\_name) | Name of the SSH key pair. | `string` | `null` | no |
| <a name="input_lan1_subnet_id"></a> [lan1\_subnet\_id](#input\_lan1\_subnet\_id) | ID of the LAN1 subnet (ETH1). | `string` | n/a | yes |
| <a name="input_mgmt_subnet_id"></a> [mgmt\_subnet\_id](#input\_mgmt\_subnet\_id) | ID of the management subnet (ETH0). | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Prefix for instance name. | `string` | `"nios-aws-instance"` | no |
| <a name="input_nios_license"></a> [nios\_license](#input\_nios\_license) | NIOS temporary license string. | `string` | `"nios IB-V825 enterprise dns dhcp cloud"` | no |
| <a name="input_private_ips_count_eth2"></a> [private\_ips\_count\_eth2](#input\_private\_ips\_count\_eth2) | Number of IPs to assign to ETH2 (HA interface). Set 1 for secondary IP (VIP) for HA, 0 for no secondary IP. | `number` | `0` | no |
| <a name="input_remote_console_enabled"></a> [remote\_console\_enabled](#input\_remote\_console\_enabled) | Enable remote console access. | `bool` | `true` | no |
| <a name="input_security_group_id"></a> [security\_group\_id](#input\_security\_group\_id) | ID of the existing AWS security group. | `string` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | Tags to apply to AWS resources. | `map(string)` | <pre>{<br/>  "Name": "nios-aws-instance",<br/>  "dontStop": "true",<br/>  "dontTerminate": "true"<br/>}</pre> | no |
| <a name="input_volume_size"></a> [volume\_size](#input\_volume\_size) | Size of the root volume in GB. | `number` | `500` | no |
| <a name="input_volume_type"></a> [volume\_type](#input\_volume\_type) | Type of the root volume. | `string` | `"gp3"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_eth0_ip"></a> [eth0\_ip](#output\_eth0\_ip) | Mgmt IP address of the EC2 instance (ETH0) |
| <a name="output_eth1_gateway"></a> [eth1\_gateway](#output\_eth1\_gateway) | Gateway IP for the LAN1 Subnet |
| <a name="output_eth1_ipv4"></a> [eth1\_ipv4](#output\_eth1\_ipv4) | LAN1 IP address of the EC2 instance (ETH1) |
| <a name="output_eth1_ipv6"></a> [eth1\_ipv6](#output\_eth1\_ipv6) | Private IPv6 address of ETH1 |
| <a name="output_eth1_subnet_mask"></a> [eth1\_subnet\_mask](#output\_eth1\_subnet\_mask) | Subnet mask of the LAN1 Subnet |
| <a name="output_eth2_eni"></a> [eth2\_eni](#output\_eth2\_eni) | ENI ID of the HA interface (ETH2) |
| <a name="output_eth2_ip"></a> [eth2\_ip](#output\_eth2\_ip) | HA IP address of the EC2 instance (ETH2) |
| <a name="output_eth2_secondary_ip_for_ha"></a> [eth2\_secondary\_ip\_for\_ha](#output\_eth2\_secondary\_ip\_for\_ha) | Secondary private IP address for HA configuration on ETH2 |
| <a name="output_instance_id"></a> [instance\_id](#output\_instance\_id) | ID of the EC2 instance |
<!-- END_TF_DOCS -->

---

## Architecture

### Standalone Mode (`ha_enable = false`)
- 1 EC2 instance with NIOS AMI
- ETH0: Mgmt interface (auto-created by AWS)
- ETH1: LAN1 Grid communication interface

### HA Mode (`ha_enable = true`)
- 1 EC2 instance with NIOS AMI (node in HA pair)
- ETH0: Mgmt interface
- ETH1: LAN1 interface
- ETH2: HA interface with secondary IP for VIP

## Usage

### Step 1: Deploy AWS Infrastructure 

```hcl
provider "aws" {
  region     = "aws_region"
  access_key = "aws_access_key"
  secret_key = "aws_secret_key"
}

module "node1" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_grid_member_aws"

  security_group_id = var.security_group_id
  mgmt_subnet_id    = var.mgmt_subnet_id
  lan1_subnet_id    = var.lan1_subnet_id

  ami_id          = var.ami_id

  instance_type     = var.instance_type
  key_name          = var.key_name
  availability_zone = var.availability_zone

  volume_size           = var.volume_size
  volume_type           = var.volume_type
  delete_on_termination = var.delete_on_termination

  enable_ipv6      = var.enable_ipv6

  name = var.name
  tags        = var.tags

  nios_license           = var.nios_license
  remote_console_enabled = var.remote_console_enabled
  default_admin_password = var.default_admin_password
}
```

**Deploy the infrastructure:**
```bash
terraform apply
```

### Step 2: Wait for NIOS to Boot

NIOS takes approximately **15-20 minutes** to fully boot.

### Step 3: Join the Grid Member to the Master Grid / Configure HA

Once Grid is up and running, configure the grid member and join to the grid.


#### Examples

#### Example 1: Join a Member to a Master

#### Deploy AWS infrastructure for Master and Member

```
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
  nios_host_url = "https://${module.node1.eth1_ipv4}"
  nios_username = "username"
  nios_password = "password"
}

resource "nios_grid_member" "member" {
  host_name        = "infoblox.member"
  config_addr_type = "BOTH"
  platform         = "VNIOS"

  vip_setting = {
    address     = module.node2.eth1_ipv4
    gateway     = module.node2.eth1_gateway
    subnet_mask = module.node2.eth1_subnet_mask
  }

  ipv6_setting = {
    virtual_ip  = module.node2.eth1_ipv6
    cidr_prefix = 64
    gateway     = "<member_ipv6_gateway_ip>"
    enabled     = true
  }
}

// Join member to existing grid master
resource "nios_grid_join" "member_join" {
  member_url       = "https://${module.node2.eth1_ipv4}"
  member_username = "UserName"
  member_password = "Password"
  grid_name       = "Infoblox"
  master          = module.node1.eth1_ipv4
  shared_secret   = "secret"
  depends_on = [nios_grid_member.member]
}
```

### Example 2: HA Grid Configuration

Deploy 2 AWS EC2 instances for SA-HA Config with the required IAM Permissions.

```hcl
# Deploy AWS infrastructure for Node 1 (Active Node)
module "node1" {
  // ... (same config as Step 1)
  ha_enable = true
  private_ips_count_eth2 = 1
  iam_instance_profile = var.iam_instance_profile
}

# Deploy AWS infrastructure for Node 2 (Passive Node)
module "node2" {
  // ... (same config as Step 1)
  ha_enable = true
  iam_instance_profile = var.iam_instance_profile
}
```
#### After both the grids are up and running (~20 min), configure HA 

1. Import Node1 under nios_grid_member.ha_pair

```hcl 
resource "nios_grid_member" "ha_pair"{}
```

```hcl 
terraform import nios_grid_member.ha_pair <uuid>
```

2. Modify the resource to set ha_on_cloud to true and provide the cloud attributes.

```
provider "nios" {
  nios_host_url = "https://${module.node1.eth1_ipv4}"
  nios_username = "admin"
  nios_password = "Infoblox@123"
}

resource "nios_grid_member" "ha_pair" {
  host_name         = "infoblox.ha-member"
  config_addr_type  = "IPV4"
  platform          = "VNIOS"
  enable_ha         = true
  router_id         = 100
  ha_on_cloud       = true
  ha_cloud_platform = "AWS"

  vip_setting = {
    address     = module.node1.eth2_secondary_ip_for_ha
    gateway     = module.node1.eth1_gateway
    subnet_mask = module.node1.eth1_subnet_mask
  }

  node_info = [
    {
      # Node 1 configuration
      lan_ha_port_setting = {
        ha_ip_address      = module.node1.eth2_ip
        mgmt_lan           = module.node1.eth1_ipv4
        ha_cloud_attribute = module.node1.eth2_eni 
      }
    },
    {
      # Node 2 configuration
      lan_ha_port_setting = {
        ha_ip_address      = module.node2.eth2_ip
        mgmt_lan           = module.node2.eth1_ipv4
        ha_cloud_attribute = module.node2.eth2_eni  
      }
    }
  ]

  // To configure grid level dns resolver settings, use the grid_level_dns_resolver_setting attribute 
  grid_level_dns_resolver_setting = {
    resolvers = [
      "10.10.10.10"
  ] }
}

3. Join Node2 (Passive Node) to Node1 (Active Node).

resource "nios_grid_join" "ha_member_join" {
  member_url      = "https://${module.node2.eth1_ipv4}"
  member_username = "admin"
  member_password = "password"
  grid_name       = "Grid Name"
  master          = module.node1.eth2_secondary_ip_for_ha
  shared_secret   = "your-shared-secret"
  depends_on      = [nios_grid_member.ha_pair]
}
```

## Outputs Usage

The module outputs can be used directly in NIOS provider resources:

| Output | NIOS Resource Usage |
|--------|---------------------|
| `eth1_ipv4` | `vip_setting.address`, `member_ip` in grid_join |
| `eth1_gateway` | `vip_setting.gateway` |
| `eth1_subnet_mask` | `vip_setting.subnet_mask` |
| `eth1_ipv6` | `ipv6_setting.virtual_ip` |
| `eth2_ip` | HA `node_info.lan_ha_port_setting.ha_ip_address` |
| `eth2_secondary_ip_for_ha` | HA `vip_setting.address` |

---

### Boot Time
- NIOS takes **15-20 minutes** to fully boot after EC2 instance creation
- Always verify NIOS API is responding before applying `nios_grid_member` resources

### HA Requirements
- Set `ha_enable = true` to create ETH2 interface
- Provide `iam_instance_profile` with permissions for HA operations
