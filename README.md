# Terraform Provider for Infoblox NIOS

The Terraform Provider for Infoblox NIOS allows you to manage your Infoblox NIOS resources such as DNS records, networks, fixed addresses, and more using Terraform. This provider uses the [infoblox-nios-go-client](https://github.com/infobloxopen/infoblox-nios-go-client) for all API calls to interact with the Infoblox NIOS WAPI.


## Table of Contents

- [Requirements](#requirements)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
    - [Setting Up Terraform Internal ID](#setting-up-terraform-internal-id)
    - [Installation](#installation)
- [Usage Examples](#usage-examples)
- [Available Resources and DataSources](#available-resources-and-datasources)
- [Host Record Management](#host-record-management)
- [Importing Existing Resources](#importing-existing-resources)
- [Documentation](#documentation)
- [Debugging](#logging-and-debugging)
- [Support](#support)

## Requirements

- [Go](https://golang.org/doc/install) >= 1.25.1
- [Terraform](https://www.terraform.io/downloads.html) >= 1.12.1
- [Infoblox NIOS](https://www.infoblox.com/products/nios/) (version 9.1.0 or higher)

## Version Compatibility Matrix

The table below shows the compatibility between different versions of the Terraform Provider NIOS and the required versions of NIOS, WAPI, Terraform, and Go.

| Provider Version | Go Version | Terraform Version | NIOS Version | WAPI Version |
|-----------------|------------|-------------------|--------------|--------------|
| 2.0.0 | >= 1.25.1 | >= 1.12.1 | 9.1.0 or higher | v2.14 |
| 1.1.0 , 1.0.0 , 0.0.1 | >= 1.18 | >= 1.8.0 | 9.0.6 or higher | v2.13.6 |

**Important Notes:**
- **Version 2.0.0+** introduces UUID based resource identification for NIOS 9.1.0. Resources must be imported using UUID instead of ref.
- **Version 1.1.0 and earlier** uses reference based resource identification. Resources must be imported using ref.
- **Recommended**: When upgrading from provider v1.x.x to v2.x.x, execute `terraform refresh` to automatically update your state file with UUID.
- For migration from the legacy terraform provider, refer to the [Migration Guide](MIGRATION.md).

### Known Issues

The following resources are affected by known WAPI limitations in NIOS 9.1.0 and are not functional for create, read, update, delete, or import operations.

**DTC Records**

- `nios_dtc_record_a`
- `nios_dtc_record_aaaa`
- `nios_dtc_record_cname`
- `nios_dtc_record_naptr`
- `nios_dtc_record_srv`

**IPAM Objects**

- `nios_ipam_ipv6networktemplate`

These limitations are due to underlying WAPI behavior in NIOS 9.1.0 and will be addressed in future releases.

## Getting Started

### Prerequisites

#### Setting Up Terraform Internal ID

- A resource can manage its drift state by using the extensible attribute `Terraform Internal ID` when its Reference ID is changed by any manual intervention.
- To use the Terraform Provider for Infoblox NIOS, you must either define the following extensible attributes in NIOS or 
  install the Cloud Network Automation license in the NIOS Grid, which adds the extensible attributes by default:
  * `Tenant ID`: String Type 
  * `CMP Type`: String Type 
  * `Cloud API Owned`: List Type (Values: True, False)
- To use the NIOS Terraform Plugin, you must either define the extensible attribute `Terraform Internal ID`
  in NIOS or use `super user` to execute the below cmd. It will create the read only extensible attribute `Terraform Internal ID`.

  ```shell
  curl -k -u <SUPERUSER>:<PASSWORD> -H "Content-Type: application/json" -X POST https://<NIOS_GRID_IP>/wapi/<WAPI_VERSION>/extensibleattributedef -d '{"name": "Terraform Internal ID", "flags": "CR", "type": "STRING", "comment": "Internal ID for Terraform Resource"}'
  ``` 

  For more details refer to the prerequisites in [Terraform Internal ID](guides/tf_internal_id_management.md) page.

### Installation

For detailed installation instructions, please refer to the [Quickstart Guide](guides/quickstart.md).

## Usage Examples

Detailed examples for each resource and data source are available in the `examples` directory of the repository. Each resource and data source has its own directory with sample configurations.

For example:
- Resources examples: [`examples/resources/nios_*`](examples/resources/)
- Data sources examples: [`examples/data-sources/nios_*`](examples/data-sources/)
- Modules examples: [`examples/modules/nios_*`](examples/modules/)

Please refer to these examples for detailed usage patterns and configurations. 

## Available Resources and DataSources

The object groups available in this provider are categorized as follows:
  - [DHCP](guides/resources_datasources.md#dhcp)
  - [DNS](guides/resources_datasources.md#dns)
  - [DTC](guides/resources_datasources.md#dtc)
  - [RPZ](guides/resources_datasources.md#rpz)
  - [IPAM](guides/resources_datasources.md#ipam)
  - [CLOUD](guides/resources_datasources.md#cloud)
  - [GRID](guides/resources_datasources.md#grid)
  - [SECURITY](guides/resources_datasources.md#security)
  - [MICROSOFT](guides/resources_datasources.md#microsoft)
  - [PARENTAL CONTROL](guides/resources_datasources.md#parental-control)
  - [RIR](guides/resources_datasources.md#rir)
  - [MISC](guides/resources_datasources.md#miscellaneous)
  - [SMARTFOLDER](guides/resources_datasources.md#smartfolder)
  - [ACL](guides/resources_datasources.md#acl)
  - [DISCOVERY](guides/resources_datasources.md#discovery)
  - [NOTIFICATION](guides/resources_datasources.md#notification)

For a detailed list of available resources and data sources, refer to the [Resources and Data Sources](guides/resources_datasources.md) page.

## Host Record Management

- The `ip_allocation` resource allocates a new IP address from an existing NIOS network and manages the corresponding DNS-related settings. It creates a Host Record in NIOS with either an IPv4 address, an IPv6 address, or both. The IP can be allocated statically (by specifying the address) or dynamically (as the next available address from a network). Once allocated, the address is marked as used in NIOS.

- The `ip_association` resource manages DHCP-related settings of the Host Record created via ip_allocation. It updates the record with VM-specific details such as the MAC address for IPv4 and the DUID for IPv6, enabling full integration with cloud or virtualized environments.

Detailed documentation for these resources can be found in [Host Record Documentation](guides/host_record_management.md) page.

## Importing Existing Resources

Resources can be imported using UUID:

For detailed information, refer to the [Importing Existing Resources](guides/importing_resources.md) page.

## Documentation

For detailed documentation, refer to the [Documentation](guides/documentation_details.md) page.

## Logging and Debugging

For detailed information, refer to the Logging and Debugging page in the docs: [Debugging](guides/logging_debugging.md)

## Support

If you have any questions or issues, you can reach out to us using the following channels:

- Github Issues:
  - Submit your issues or requests for enhancements on the [Github Issues Page](https://github.com/infobloxopen/terraform-provider-nios/issues)
- Infoblox Support:
  - For any questions or issues, please contact [Infoblox Support](https://info.infoblox.com/contact-form/).
