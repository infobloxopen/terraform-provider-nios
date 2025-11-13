// Retrieve a specific IPv6 Fixed Address Template by filters
data "nios_dhcp_ipv6fixedaddresstemplate" "get_ipv6_fixed_address_template_using_filters" {
  filters = {
    name = "example_ipv6_fixed_address_template"
  }
}

// Retrieve specific IPv6 Fixed Address Templates using Extensible Attributes
data "nios_dhcp_ipv6fixedaddresstemplate" "get_ipv6_fixed_address_template_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all IPv6 Fixed Address Templates
data "nios_dhcp_ipv6fixedaddresstemplate" "get_all_ipv6_fixed_address_templates" {}
