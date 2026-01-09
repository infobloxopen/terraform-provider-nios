// Manage DHCP Fingerprint with Basic Fields
resource "nios_dhcp_fingerprint" "dhcp_fingerprint_basic" {
  device_class = "Windows OS"
  name         = "example_fingerprint_1"
  option_sequence = [
    "1,2,3,4,5,6,7,8,9,10,11,99,100,199,205",
    "1,3,6,15,28,51,58,59",
    "22,23,24,25,26,27,43,44,45,46,47,48,49,55"
  ]
}

// Manage DHCP Fingerprint with Additional Fields
resource "nios_dhcp_fingerprint" "dhcp_fingerprint_with_additional_fields" {
  device_class = "Windows OS"
  name         = "example_fingerprint_2"
  option_sequence = [
    "1,2,3,4,5,6,7,8,9,10,11,99,100,199,206",
    "22,23,24,25,26,27,43,44,45,46,47,48,49,51"
  ]

  // Additional Fields
  comment = "DHCP Fingerprint managed by Terraform"
  type    = "CUSTOM"
  disable = false
  vendor_id = [
    "vendor-id-1",
    "vendor-id-2"
  ]

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}
