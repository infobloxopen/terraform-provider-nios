// Create a DHCP IPv6 Fixed Address Template with Basic Fields
resource "nios_dhcp_ipv6fixedaddresstemplate" "basic" {
  name = "example_ipv6_fixed_address_template_1"
}

// Create a DHCP IPv6 Fixed Address Template with Additional Fields
resource "nios_dhcp_ipv6fixedaddresstemplate" "additional_fields" {
  name = "example_ipv6_fixed_address_template_2"
  // Additional Fields
  comment         = "IPv6 Fixed Address Template Created by Terraform"
  domain_name     = "example.com"
  use_domain_name = true
  options = [
    {
      name  = "dhcp-lease-time"
      num   = "51"
      value = "5000"
    },
    {
      name  = "domain-name-servers"
      num   = 6
      value = "2001:4860:4860::8888,2001:4860:4860::8844"
    }
  ]
  use_options        = true
  valid_lifetime     = 5000
  use_valid_lifetime = true
  extattrs = {
    "Site" = "location-1"
  }
}

terraform {
	  required_providers {
	    nios = {
	      source  = "infobloxopen/nios"
	      version = "1.1.0"
	    }
	  }
	}
	
	provider "nios" {
	  nios_host_url = "https://172.28.82.33"
	  nios_username = "admin"
	  nios_password = "Infoblox@123"
}