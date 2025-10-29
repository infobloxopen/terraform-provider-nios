# Migrating from the Legacy Infoblox Provider

This guide covers the changes introduced in the NIOS Terraform provider and outlines the steps you may need to take to upgrade your configuration.

The NIOS Terraform provider replaces the [Infoblox provider](https://registry.terraform.io/providers/infobloxopen/infoblox/latest) and is not backwards compatible. This means you will need to update your configuration to use the new provider.

## Prerequisites

- Terraform v1.5.0 or later
- Access to your Infoblox NIOS appliance
- Backup of your current Terraform state files


### Backup Your State Files

Before making changes to your state, it's a good idea to back up your state file. Any state modification commands made using the CLI will automatically create a backup. If you prefer to manually back up your state file, you can copy your `terraform.tfstate` file to a backup location.

Having a backup ensures that you have a snapshot of your infrastructure's state at a specific moment, allowing you to revert or refer to it if necessary.

### Add the New Provider

You will need to add the new NIOS provider to your configuration. Update your `terraform` block:

```hcl
terraform {
  required_providers {
    nios = {
      source  = "infobloxopen/nios"
      version = "1.0.0"
    }
  }
}

provider "nios" {
  nios_host_url = "<grid_master_ip>"
  nios_username = "<user_name>"
  nios_password = "<password>"
}
```

Run `terraform init` to download the new provider.

### Resource and Data Source Mapping

The resource and data source names have changed in the new provider. The following table shows the mapping from legacy to new names:

| Legacy Provider | New NIOS Provider |
|----------------|-------------------|
| `infoblox_network_view` | `nios_ipam_network_view` |
| `infoblox_ipv4_network_container` | `nios_ipam_network_container` |
| `infoblox_ipv6_network_container` | `nios_ipam_ipv6network_container` |
| `infoblox_ipv4_network` | `nios_ipam_network` |
| `infoblox_ipv6_network` | `nios_ipam_ipv6network` |
| `infoblox_a_record` | `nios_dns_record_a` |
| `infoblox_aaaa_record` | `nios_dns_record_aaaa` |
| `infoblox_dns_view` | `nios_dns_view` |
| `infoblox_ptr_record` | `nios_dns_record_ptr` |
| `infoblox_cname_record` | `nios_dns_record_cname` |
| `infoblox_mx_record` | `nios_dns_record_mx` |
| `infoblox_txt_record` | `nios_dns_record_txt` |
| `infoblox_srv_record` | `nios_dns_record_srv` |
| `infoblox_zone_auth` | `nios_dns_zone_auth` |
| `infoblox_zone_forward` | `nios_dns_zone_forward` |
| `infoblox_zone_delegated` | `nios_dns_zone_delegated` |
| `infoblox_dtc_lbdn` | `nios_dtc_lbdn` |
| `infoblox_dtc_pool` | `nios_dtc_pool` |
| `infoblox_dtc_server` | `nios_dtc_server` |
| `infoblox_alias_record` | `nios_dns_record_alias` |
| `infoblox_ns_record` | `nios_dns_record_ns` |
| `infoblox_ipv4_shared_network` | `nios_dhcp_shared_network` |
| `infoblox_ipv4_fixed_address` | `nios_dhcp_fixed_address` |
| `infoblox_ipv4_range` | `nios_dhcp_range` |
| `infoblox_ipv4_range_template` | `nios_dhcp_range_template` |
| `infoblox_ip_allocation` | `nios_ip_allocation` |
| `infoblox_ip_association` | `nios_ip_association` |
| `infoblox_host_record` | `nios_dns_record_host` |


### Attribute Name Changes

The new NIOS Terraform provider uses attribute names that are in parity with the NIOS WAPI field names. This means many attribute names have changed. Some common changes include:

**Legacy Provider:**
```hcl
resource "infoblox_a_record" "example" {
  fqdn         = "test.example.com"
  ip_addr      = "10.0.0.1"
  ttl          = 300
  dns_view     = "default"
  comment      = "Test record"
  ext_attrs = jsonencode({
    "Site" = "location-1"
  })

}
```

**New NIOS Provider:**
```hcl
resource "nios_dns_record_a" "example" {
  name     = "test.example.com"
  ipv4addr = "10.0.0.1"
  ttl      = 300
  view     = "default"
  comment  = "Test record"
  extattrs = {
    Site = "location-1"
  }
}
```

### Replace Resources in State

#### Get Resource IDs

First, get the IDs of all existing resources:

```bash
terraform show -json | jq -c '.values.root_module.resources[] | {"resource":.address, "id":.values.id}'
```

#### Remove Old Resource from State

Remove the old resource from state:

```bash
terraform state rm infoblox_a_record.example
```

#### Import New Resource into State

Import the new resource using the same ID:

```bash
terraform import nios_dns_record_a.example "record:a/ZG5zLmEkLl9kZWZhdWx0LmNvbS5pbmZvYmxveC50ZXN0:test.example.com/default"
```

**Recommended Approach**: If you are using Terraform v1.5.0 or later, use the import block with configuration generation:

```hcl
import {
  to = nios_dns_record_a.example
  id = "record:a/ZG5zLmEkLl9kZWZhdWx0LmNvbS5pbmZvYmxveC50ZXN0:test.example.com/default"
}
```

You can use Terraform's configuration generation feature along with import blocks:

```bash
# Generate configuration from existing resources
terraform plan -generate-config-out=generated.tf

# Review the generated configuration
cat generated.tf

# Apply the import
terraform apply
```

This approach will:
1. Generate the appropriate resource configuration automatically
2. Import the resources into your state

### Unsupported block type

Configuration written as blocks will have to be rewritten as values. For example, if you have a block like this:

```terraform
  auth_zones {
    fqdn = "info.com"
    dns_view = "default.view2"
  }
  auth_zones {
    fqdn = "test.com"
    dns_view = "default"
  }
```

you will have to rewrite it with an equal sign :
```terraform
  auth_zones = [
    nios_dns_zone_auth.parent_zone.ref,
    nios_dns_zone_auth.parent_zone2.ref
  ]
```

## ⚠️ Critical Warning

**IMPORTANT**: The new NIOS provider uses a unique identifier called "Terraform Internal ID" under extensible attributes (`extattrs_all`) that is updated during the Terraform plan phase. If you generate a plan for import and it fails, this unique identifier will be updated and you will **NOT** be able to use the legacy Infoblox provider anymore for those resources.

**Migration Complexity**: If you want to replace existing resource names and field mappings, the migration can become complicated because:
- The Terraform Internal ID structure has changed
- The extensible attributes are segregated into `extattrs` and `extattrs_all`. Where `extattrs_all` is just for internal use.
- Field names now match NIOS WAPI field names exactly

**Recommendation**: 
- Always backup your state files and test the migration process in a non-production environment first
- Use the import-only approach outlined in this guide rather than attempting to replace existing resources
- Consider starting with a fresh Terraform configuration for complex migrations
