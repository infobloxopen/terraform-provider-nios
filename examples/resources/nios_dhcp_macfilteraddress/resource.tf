// Create a MAC Filter (Required as Parent)
resource "nios_dhcp_filtermac" "parent_mac_filter" {
  name = "mac_filter_example"
}

// Create MAC Filter Address with Basic Fields
resource "nios_dhcp_macfilteraddress" "mac_filter_address_basic_fields" {
  mac    = "00:5A:2B:3C:4D:5E"
  filter = nios_dhcp_filtermac.parent_mac_filter.name
}

// Create MAC Filter Address with Additional Fields
resource "nios_dhcp_macfilteraddress" "mac_filter_address_additional_fields" {
  mac                   = "01:5A:2B:3C:4D:5E"
  filter                = nios_dhcp_filtermac.parent_mac_filter.name
  authentication_time   = 7200
  comment               = "This is a sample MAC filter address"
  expiration_time       = 1440
  never_expires         = false
  reserved_for_infoblox = "Updated_Reserved_For_Infoblox_Value"
  guest_custom_field1   = "CustomValue1"
  guest_custom_field2   = "CustomValue2"
  guest_custom_field3   = "CustomValue3"
  guest_custom_field4   = "CustomValue4"
  guest_email           = "abc@example.com"
  guest_first_name      = "John"
  guest_last_name       = "Doe"
  guest_middle_name     = "M"
  guest_phone           = "+1234567890"

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}
