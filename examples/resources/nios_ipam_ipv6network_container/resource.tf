terraform {
  required_providers {
    nios = {
      source  = "infobloxopen/nios"
      version = "1.0.0"
    }
  }
}

provider "nios" {
  nios_host_url = "https://172.28.82.213"
  nios_username = "admin"
  nios_password = "Infoblox@123"
}

// Create IPV6 Network Container with Basic Fields
resource "nios_ipam_ipv6network_container" "example_container" {
  network      = "11::/64"
  subscribe_settings = {
    enabled_attributes = [
      "SECURITY_GROUP",
    ]
  }
}

// Create IPV6 Network Container with Additional Fields
resource "nios_ipam_ipv6network_container" "complete_example" {
  // Required attributes
  network = "11::/64"

  // Basic configuration
  network_view = "default"
  comment      = "IPv6 network container with additional fields"

  options = [
    {
      name         = "dhcp6.fqdn",
      num          = 39,
      value        = "test_options.com",
      vendor_class = "DHCPv6"
    }
  ]
  use_options = true
  // DDNS settings
  enable_ddns                = true
  use_enable_ddns            = true
  ddns_domainname            = "example.com"
  ddns_generate_hostname     = true
  ddns_ttl                   = 3600
  use_ddns_domainname        = true
  use_ddns_generate_hostname = true
  use_ddns_ttl               = true

  // Extensible attributes
  extattrs = {
    Site = "location-1"
  }
}


// Create IPV6 Network Container with Function Call
resource "nios_ipam_ipv6network_container" "example_func_call" {
  func_call = {
    attribute_name  = "network"
    object_function = "next_available_network"
    result_field    = "networks"
    object          = "ipv6networkcontainer"
    object_parameters = {
      network      = "10::/64"
      network_view = "default"
    }
    parameters = {
      cidr = 72
    }
  }
  comment = "Network container created with function call"
  depends_on = [
    nios_ipam_ipv6network_container.example_container
  ]
}
