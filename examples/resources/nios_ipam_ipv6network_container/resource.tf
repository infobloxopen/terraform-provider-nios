// Create IPV6 Network with Basic Fields
resource "nios_ipam_ipv6network_container" "example_container" {
  network      = "10::/64"
  network_view = "default"
  comment      = "Created by Terraform"

  // Optional: Configure extensible attributes
  extattrs = {
    "Site" = "location-1"
  }
}

resource "nios_ipam_ipv6network_container" "complete_example" {
  // Required attributes
  network = "11::/64"

  // Basic configuration
  network_view = "default"
  comment      = "IPv6 network container with additional fields"


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
    "Site" = "DataCenter1"
  }
}


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
