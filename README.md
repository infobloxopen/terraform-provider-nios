# Terraform Provider for Infoblox NIOS

The Terraform Provider for Infoblox NIOS allows you to manage your Infoblox NIOS resources such as DNS records, networks, fixed addresses, and more using Terraform. This provider uses the [infoblox-nios-go-client](https://github.com/infobloxopen/infoblox-nios-go-client) for all API calls to interact with the Infoblox NIOS WAPI.


## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
  - [Prerequisites](#prerequisites)
    - [Terraform Internal ID](#terraform-internal-id)
  - [Terraform RC Configuration for local usage](#terraform-rc-configuration-for-local-usage)
  - [Using Pre-built Binaries from Github Releases](#using-pre-built-binaries-from-github-releases)
  - [Build the Provider from Source](#build-the-provider-from-source)
- [Example Provider Configuration](#example-provider-configuration)
  - [Provider Arguments](#provider-arguments)
- [Usage Examples](#usage-examples)
- [Available Resources and DataSources](#available-resources-and-datasources)
- [Host Record Management](#host-record-management)
- [Importing Existing Resources](#importing-existing-resources)
- [Documentation](#documentation)
- [Debugging](#debugging)
  - [Terraform Logging](#terraform-logging)
  - [Provider-Specific Debugging](#provider-specific-debugging)
- [Terraform Limitations / Anomalies and Known Issues](#terraform-limitations--anomalies-and-known-issues)
- [Support](#support)

## Requirements

- [Go](https://golang.org/doc/install) >= 1.18 (to build the provider plugin) (recommended version is 1.24.4 or later)
- [Terraform](https://www.terraform.io/downloads.html) >= 1.8.0
- [Infoblox NIOS](https://www.infoblox.com/products/nios/) (version 9.0.6 or higher)

## Installation

### Prerequisites

#### Terraform Internal ID

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

  For more details refer to the prerequisites in [Terraform Internal ID](guides/tf_internal_id) page.

### Using Pre-built Binaries from Github Releases

1. Download the latest release from the [releases page](https://github.com/infobloxopen/terraform-provider-nios/releases).
2. Extract the binary and move it to the Terraform plugins directory (`~/.terraform.d/plugins/`) . Use the following command to create the necessary directory structure:
```bash
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/infobloxopen/nios/0.0.1/<OS_ARCH>(linux_amd64, darwin_amd64, windows_amd64)
mv terraform-provider-nios ~/.terraform.d/plugins/registry.terraform.io/infobloxopen/nios/0.0.1/<OS_ARCH>
```
3. Additional Step for macOS Users:
   On Apple devices, you must authorize the binary to run by executing the following command once:
```bash
xattr -d com.apple.quarantine ~/.terraform.d/plugins/registry.terraform.io/infobloxopen/nios/0.0.1/<OS_ARCH>/terraform-provider-nios
```

### Build the Provider from Source

Instead of using pre-built binaries, you can build the provider from source. This is useful for development and testing purposes and to build the latest changes pushed to the repository.

1. Clone the repository:
```bash
git clone https://github.com/infobloxopen/terraform-provider-nios.git
```

2. Change to the repository directory:
```bash
cd <path-to-provider>/terraform-provider-nios
```

3. Ensure you have the necessary dependencies installed. You can use `go mod tidy` to ensure all dependencies are fetched:
```bash
go mod tidy
go mod vendor
```

3. Build and install the provider:
```bash
make build
make install
```

OR instead of `make install`, you can manually move the built binary to the Terraform plugins directory:

```bash
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/infobloxopen/nios/0.0.1/<OS_ARCH>(linux_amd64, darwin_amd64, windows_amd64)
mv terraform-provider-nios ~/.terraform.d/plugins/registry.terraform.io/infobloxopen/nios/0.0.1/<OS_ARCH>
```

4. Additional Step for macOS Users:
   On Apple devices, you must authorize the binary to run by executing the following command once:
```bash
xattr -d com.apple.quarantine ~/.terraform.d/plugins/registry.terraform.io/infobloxopen/nios/0.0.1/<OS_ARCH>/terraform-provider-nios
```



This configuration allows Terraform to use your local provider instead of the one from the Terraform registry, which is particularly useful during development and testing.
 

## Example Provider Configuration

```hcl
terraform {
  required_providers {
    nios = {
      source  = "infobloxopen/nios"
      version = "0.0.1"
    }
  }
}

provider "nios" {
  nios_host_url = "<NIOS_HOST_URL>"
  nios_username = "<NIOS_USERNAME>"
  nios_password = "<NIOS_PASSWORD>"
}
```

### Provider Arguments

- `nios_host_url` - (Required) The full URL of your NIOS Grid Manager (e.g., "https://gridmaster.example.com").
- `nios_username` - (Required) The username to access the NIOS Grid Manager.
- `nios_password` - (Required) The password to access the NIOS Grid Manager.

## Usage Examples

Detailed examples for each resource and data source are available in the `examples` directory of the repository. Each resource and data source has its own directory with sample configurations.

For example:
- Resources examples: [`examples/resources/nios_*`](examples/resources/)
- Data sources examples: [`examples/data-sources/nios_*`](examples/data-sources/)

Please refer to these examples for detailed usage patterns and configurations. 

## Available Resources and DataSources

The object groups available in this provider are categorized as follows:
  - [DHCP](guides/resources_datasources.md#dhcp)
  - [DNS](guides/resources_datasources.md#dns)
  - [DTC](guides/resources_datasources.md#dtc)
  - [IPAM](guides/resources_datasources.md#ipam)
  - [CLOUD](guides/resources_datasources.md#cloud)
  - [SECURITY](guides/resources_datasources.md#security)
  - [MISC](guides/resources_datasources.md#misc)
  - [SMARTFOLDER](guides/resources_datasources.md#smartfolder)
  - [ACL](guides/resources_datasources.md#acl)
  - [GRID](guides/resources_datasources.md#grid)
  - [DISCOVERY](guides/resources_datasources.md#discovery)
  - [NOTIFICATION](guides/resources_datasources.md#notification)

For a detailed list of available resources and data sources, refer to the [Resources and Data Sources](guides/resources_datasources.md) page.

## Host Record Management

- The `ip_allocation` resource allocates a new IP address from an existing NIOS network and manages the corresponding DNS-related settings. It creates a Host Record in NIOS with either an IPv4 address, an IPv6 address, or both. The IP can be allocated statically (by specifying the address) or dynamically (as the next available address from a network). Once allocated, the address is marked as used in NIOS.

- The `ip_association` resource manages DHCP-related settings of the Host Record created via ip_allocation. It updates the record with VM-specific details such as the MAC address for IPv4 and the DUID for IPv6, enabling full integration with cloud or virtualized environments.

**Note:**

- Do not destroy the `ip_association` resource directly; destroying the `ip_allocation` will automatically remove the associated record.
- Each allocation supports at most one IPv4 and one IPv6 address. Multiple addresses of the same family are not supported.

Detailed documentation for these resources can be found in [Host Record Documentation](guides/host_record.md) page.

## Importing Existing Resources

Resources can be imported using their reference ID:

For detailed information, refer to the [Importing Existing Resources](guides/importing_resources.md) page.

## Documentation

For detailed documentation, refer to the [Documentation](guides/documentation.md) page.

## Logging and Debugging

For detailed information, refer to the Logging and Debugging page in the docs: [Debugging](guides/logging_debugging.md)

##  Terraform Limitations / Anomalies and Known Issues

For detailed information about limitations, refer to the [Terraform Limitations / Anomalies and Known Issues](guides/limitations.md) page.

For details information about known issues, refer to the [Terraform Known Issues](guides/known_issues.md) page.

### Terraform RC Configuration for local usage

As the Provider isn't available on registry , to develop the provider locally , you need to set up the `.terraformrc` file in your home directory to point to the `terraform-provider-nios` repository.

```bash
provider_installation {
  dev_overrides {
    "infobloxopen/nios" = "/Users/<user-name>/<path-to-provider>/terraform-provider-nios"
  }
  filesystem_mirror {
    path    = "/Users/<user-name>/.terraform.d/plugins/"
    include = ["infobloxopen/nios"]
  }
  direct {
    exclude = ["infobloxopen/nios"]
  }
}
```
Using this configuration allows Terraform to use the local provider instead of the one from the Terraform registry, which is particularly useful during development and testing.

## Support

If you have any questions or issues, you can reach out to us using the following channels:

- Github Issues:
  - Submit your issues or requests for enhancements on the [Github Issues Page](https://github.com/infobloxopen/terraform-provider-nios/issues)
- Infoblox Support:
  - For any questions or issues, please contact [Infoblox Support](https://info.infoblox.com/contact-form/).
