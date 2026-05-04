// List specific Fixed Addresses using filters
list "nios_dhcp_fixed_address" "list_fixed_addresses_using_filters" {
  provider = nios
  config {
    filters = {
      ipv4addr = "10.0.0.1"
    }
  }
}

// List specific Fixed Addresses using Extensible Attributes
list "nios_dhcp_fixed_address" "list_fixed_addresses_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Fixed Addresses with resource details included
list "nios_dhcp_fixed_address" "list_fixed_addresses_with_resource" {
  provider         = nios
  include_resource = true
}
