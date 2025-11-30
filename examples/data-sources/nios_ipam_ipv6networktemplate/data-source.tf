// Retrieve a specific IPAM IPv6 Network Template using filters
data "nios_ipam_ipv6networktemplate" "get_ipam_ipv6network_template_using_filters" {
  filters = {
    name = "example_ipv6_network_template"
  }
}

// Retrieve specific IPAM IPv6 Network Templates using Extensible Attributes
data "nios_ipam_ipv6networktemplate" "get_ipam_ipv6network_template_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all IPAM IPv6 Network Templates
data "nios_ipam_ipv6networktemplate" "get_all_ipam_ipv6network_templates" {}
