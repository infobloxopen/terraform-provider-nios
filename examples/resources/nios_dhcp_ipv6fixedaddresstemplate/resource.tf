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
    }
  ]
  use_options        = true
  valid_lifetime     = 5000
  use_valid_lifetime = true
  extattrs = {
    "Site" = "location-1"
  }
}
