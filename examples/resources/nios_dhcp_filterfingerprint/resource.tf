// Create a DHCP Fingerprints (Required as Parent)
resource "nios_dhcp_fingerprint" "dhcp_fingerprint_parent" {
  name         = "example_fingerprint_1"
  device_class = "Windows OS"
  vendor_id    = ["MSFT"]
}

resource "nios_dhcp_fingerprint" "dhcp_fingerprint_parent_2" {
  name         = "example_fingerprint_2"
  device_class = "Windows OS"
  vendor_id    = ["MSFT"]

}
// Create a DHCP Fingerprint Filter with Basic Fields
resource "nios_dhcp_filterfingerprint" "dhcp_filterfingerprint_basic" {
  name = "example_filterfingerprint_1"
  fingerprint = [
    nios_dhcp_fingerprint.dhcp_fingerprint_parent.name,
    nios_dhcp_fingerprint.dhcp_fingerprint_parent_2.name
  ]
}

// Manage dhcp Filterfingerprint with Additional Fields
resource "nios_dhcp_filterfingerprint" "dhcp_filterfingerprint_with_additional_fields" {
  name = "example_filterfingerprint_2"
  fingerprint = [
    nios_dhcp_fingerprint.dhcp_fingerprint_parent.name,
    nios_dhcp_fingerprint.dhcp_fingerprint_parent_2.name
  ]
  comment = "Fingerprint Filter created by Terraform."
  //Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}
