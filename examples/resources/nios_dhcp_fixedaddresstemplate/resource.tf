// Create a DHCP Fixed Address Template with Basic Fields
resource "nios_dhcp_fixedaddresstemplate" "basic" {
  name = "example_fixed_address_template_1"
}

// Create a DHCP Fixed Address Template with Additional Fields
resource "nios_dhcp_fixedaddresstemplate" "additional_fields" {
  name = "example_fixed_address_template_2"

  // Additional Fields
  comment             = "Fixed Address Template Created by Terraform"
  ddns_domainname     = "example.com"
  use_ddns_domainname = true

  options = [
    {
      "name" : "time-offset",
      "num" : 2,
      "value" : "50",
    },
    {
      "name" : "subnet-mask",
      "value" : "1.1.1.1",
    },
  ]
  use_options = true

  enable_ddns     = true
  use_enable_ddns = true

  extattrs = {
    "Site" = "location-1"
  }
}
