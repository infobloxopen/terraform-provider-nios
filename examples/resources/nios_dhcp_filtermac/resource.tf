#create a filtermac resource
resource "nios_dhcp_filtermac" "mac_filter" {
  name = "mac_filter_example"
}

#Create another filtermac with different name
resource "nios_dhcp_filtermac" "mac_filter_update" {
  name                           = "mac_filter_example_2"
  lease_time                     = 7200
  enforce_expiration_times       = true
  default_mac_address_expiration = 1440
  never_expires                  = false
  reserved_for_infoblox          = "Updated_Reserved_For_Infoblox_Value"

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}

# Create another filtermac resource with dhcp option dhcp_lease_time
resource "nios_dhcp_filtermac" "mac_filter_with_option" {
  name = "mac_filter_with_option"
  options = [
    {
      name  = "dhcp-lease-time"
      num   = 51
      value = "3600"
    }
  ]
}
