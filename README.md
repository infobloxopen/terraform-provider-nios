# Terraform Provider for Infoblox NIOS

The Terraform Provider for Infoblox NIOS allows you to manage your Infoblox NIOS resources such as DNS records, networks, fixed addresses, and more using Terraform.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
- [Go](https://golang.org/doc/install) >= 1.18 (to build the provider plugin)
- [Infoblox NIOS](https://www.infoblox.com/products/nios/) >= 9.0.6

## Installation

### Terraform RC Configuration for local usage 

As the Provider isn't available on registry , to use a locally built version of the provider for development purposes:

1. Modify the `.terraformrc` file in your home directory:

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

### Manual Installation

#### Using Pre-built Binaries from Github Releases

1. Download the latest release from the [releases page](https://github.com/infobloxopen/terraform-provider-nios/releases).
2. Extract the binary and move it to the Terraform plugins directory (`~/.terraform.d/plugins/`) . Use the following command to create the necessary directory structure:
```bash
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/infobloxopen/nios/1.0.0/<OS_ARCH>(linux_amd64, darwin_amd64, windows_amd64)
```

#### Build the Provider from Source

1. Clone the repository:
```bash
git clone https://github.com/infobloxopen/terraform-provider-nios.git
```

2. Change to the repository directory:
```bash
cd terraform-provider-nios
```

3. Build and install the provider:
```bash
make build
make install
```

OR instead of `make install`, you can manually move the built binary to the Terraform plugins directory:

```bash
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/infobloxopen/nios/1.0.0/linux_amd64
mv terraform-provider-nios_v1.0.0 ~/.terraform.d/plugins/registry.terraform.io/infobloxopen/nios/1.0.0/linux_amd64/terraform-provider-nios_v1.0.0
```


This configuration allows Terraform to use your local provider instead of the one from the Terraform registry, which is particularly useful during development and testing.

## Example Provider Configuration

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
- Resources examples: `examples/resources/nios_*`
- Data sources examples: `examples/data-sources/nios_*`

Please refer to these examples for detailed usage patterns and configurations.

## Available Resources

- `nios_dns_record_a` - Manages A records
- `nios_dns_record_aaaa` - Manages AAAA records
- `nios_dns_record_cname` - Manages CNAME records
- `nios_dns_record_mx` - Manages MX records
- `nios_dns_record_txt` - Manages TXT records
- `nios_dns_record_srv` - Manages SRV records
- `nios_dns_record_ptr` - Manages PTR records
- `nios_dns_view` - Manages DNS views
- `nios_dns_zone` - Manages DNS zones
- `nios_dhcp_fixedaddress` - Manages fixed addresses
- `nios_dhcp_range` - Manages DHCP ranges
- `nios_dhcp_network` - Manages DHCP networks
- `nios_dhcp_network_container` - Manages DHCP network containers

## Available Data Sources

- `nios_dns_record_a` - Fetches A records
- `nios_dns_record_aaaa` - Fetches AAAA records
- `nios_dns_record_cname` - Fetches CNAME records
- `nios_dns_view` - Fetches DNS views
- `nios_dns_zone` - Fetches DNS zones
- `nios_dhcp_fixedaddress` - Fetches fixed addresses
- `nios_dhcp_range` - Fetches DHCP ranges
- `nios_dhcp_network` - Fetches DHCP networks

## Importing Existing Resources

Resources can be imported using their reference ID:

```bash
terraform import nios_dns_record_a.example record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmV4YW1wbGUsc2FtcGxlLDE5Mi4xNjguMS4xMA:example.mydomain.com/default
```

Alternatively, you can use Terraform's import blocks (available in Terraform 1.5.0 and later) to declaratively import resources:

```hcl
import {
  to = nios_dns_record_a.example
  id = "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmV4YW1wbGUsc2FtcGxlLDE5Mi4xNjguMS4xMA:example.mydomain.com/default"
}

resource "nios_dns_record_a" "example" {
  # Configuration will be imported from the ID
  # After import, update the configuration as needed
}
```

After running `terraform plan` and `terraform apply`, the resource will be imported and you can then update the configuration as needed.

## Building the Provider

1. Clone the repository:

```bash
git clone https://github.com/infobloxopen/terraform-provider-nios.git
```

2. Build the provider:

```bash
cd terraform-provider-nios
go build
```

## Contributing

Contributions are welcome! Please read the [contribution guidelines](CONTRIBUTING.md) before submitting a pull request.

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MPL 2.0 License - see the [LICENSE](LICENSE) file for details.

## Support

For issues, feature requests, or questions, please [open an issue](https://github.com/infobloxopen/terraform-provider-nios/issues/new) on GitHub.

For commercial support, please contact [Infoblox Support](https://www.infoblox.com/support/).
