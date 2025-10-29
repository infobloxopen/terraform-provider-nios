---
page_title: "Managing DDI services with the NIOS Terraform Provider"
subcategory: "Guides"
description: |-
  This guide provides step-by-step instructions for using the NIOS Terraform Provider to manage DDI resources.
---

# Managing DDI services with the NIOS Terraform Provider

This guide provides step-by-step instructions for using the NIOS Terraform Provider to manage DDI resources.

## Configuring the Provider

Before getting started, ensure you have completed the [prerequisites](../README.md#prerequisites).

The provider needs to be configured with a `NIOSHostURL`, `NIOSUsername` and `NIOSPassword`.

Create a directory for the Terraform configuration and create a file named `main.tf` with the following content:

````terraform
terraform {
  required_providers {
    nios = {
      source  = "infobloxopen/nios"
      version = ">= 0.0.1"
    }
  }
  required_version = ">= 1.8.0"
}

provider "nios" {
  nios_host_url = "<NIOS_HOST_URL>"
  nios_username = "<NIOS_USERNAME>"
  nios_password = "<NIOS_PASSWORD>"
}
````

> ⚠️ **Warning**: Hard-coded credentials are not recommended in any configuration file. It is recommended to use environment variables.

You can also use the following environment variables to configure the provider: NIOS_HOST_URL, NIOS_USERNAME and NIOS_PASSWORD.

Initialize the provider by running the following command. This will download the provider and initialize the working directory.


```shell
terraform init
```

## Configuring Resources

This section demonstrates how to create and manage resources using the NIOS Terraform provider.

### DNS Resources

### Authoritative Zone

In this example, you will use the following resources to create an authoritative zone.

- [nios_dns_zone_auth](https://registry.terraform.io/providers/infobloxopen/nios/latest/docs/resources/nios_dns_zone_auth)

Add the following to the `main.tf` file:

````terraform
// Create a DNS zone for the domain
resource "nios_dns_zone_auth" "create_zone_auth" {
  fqdn = "exampledomain.com"
  extattrs = {
    Site = "location-1"
  }
}
````

Here the `view` attribute has not been set, so the default view will be used.

### DNS Records
Further, you will create an A record and a CNAME record within the zone.

You will use the following resources to create these
- [nios_dns_record_a](https://registry.terraform.io/providers/infobloxopen/nios/latest/docs/resources/nios_dns_record_a)
- [nios_dns_record_cname](https://registry.terraform.io/providers/infobloxopen/nios/latest/docs/resources/nios_dns_record_cname)

Add the following code to your main.tf:

````terraform
// Create an A record
resource "nios_dns_record_a" "create_record_a" {
  name     = "a_record.${nios_dns_zone_auth.create_zone_auth.fqdn}"
  ipv4addr = "10.0.0.10"
  view     = "default"
  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}

// Create a CNAME record
resource "nios_dns_record_cname" "create_record_cname" {
  name      = "cname_record.${nios_dns_zone_auth.create_zone_auth.fqdn}"
  canonical = "example-canonical-name.${nios_dns_zone_auth.create_zone_auth.fqdn}"
  depends_on = [nios_dns_zone_auth.create_zone_auth]
}
````

## IPAM Resources
In this example, you will use the following resources to create a Network View and a Network

- [nios_ipam_network_view](https://registry.terraform.io/providers/infobloxopen/nios/latest/docs/resources/nios_ipam_network_view)
- [nios_ipam_network](https://registry.terraform.io/providers/infobloxopen/nios/latest/docs/resources/nios_ipam_network)

Add the following to your main.tf:

````terraform
//Create a Network View
resource "nios_ipam_network_view" "create_network_view" {
  name = "example_network_view"
}

// Create an IPV4 Network
resource "nios_ipam_network" "create_network" {
  network      = "15.0.0.0/24"
  network_view = "example_network_view"
  comment      = "Created by Terraform"
  extattrs = {
    "Site" = "location-1"
  }
}
````

## DHCP Resources
In this example, you will use the following resources to create a Fixed Address within the Network created above. 

- [nios_dhcp_fixed_address](https://registry.terraform.io/providers/infobloxopen/nios/latest/docs/resources/nios_dhcp_fixed_address)

Add the following to your main.tf:

````terraform
//Create a fixed address within the above network
resource "nios_dhcp_fixed_address" "create_fixed_address" {
  ipv4addr     = "15.0.0.10"
  match_client = "MAC_ADDRESS"
  mac          = "00:1a:2b:3c:4d:5e"
  depends_on   = [nios_ipam_network.create_network]
}
````

You can now run `terraform plan` to see what resources will be created.

```shell
terraform plan
```

## Applying the Configuration

To create the resources, run the following command:

```shell
terraform apply
```

## Destroying the Configuration

To destroy all the resources, run the following command:

```shell
terraform destroy
```

## Configuring Datasources

Datasources allow you to retrieve existing NIOS objects. Here's a simple example:

````terraform
// Get an existing DNS zone
data "nios_dns_zone_auth" "get_auth_zone" {
  fqdn = "exampledomain.com"
  view = "default"
}

// Output the zone information
output "zone_info" {
  value = {
    zone_name = data.nios_dns_zone_auth.get_auth_zone.fqdn
    zone_view = data.nios_dns_zone_auth.get_auth_zone.view
  }
}
````

## Next steps

You can also use the NIOS Terraform Provider to manage other resources. For more information, see the [NIOS Terraform Provider documentation](https://registry.terraform.io/providers/infobloxopen/nios/latest/docs).