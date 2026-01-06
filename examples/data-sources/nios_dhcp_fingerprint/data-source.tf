// Retrieve a specific DHCP Fingerprint by filters
data "nios_dhcp_fingerprint" "get_dhcp_fingerprint_using_filters" {
  filters = {
    device_class = "Windows OS"
    name         = "example_fingerprint_1"
  }
}
// Retrieve specific DHCP Fingerprints using Extensible Attributes
data "nios_dhcp_fingerprint" "get_dhcp_fingerprints_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DHCP Fingerprints
data "nios_dhcp_fingerprint" "get_all_dhcp_fingerprints" {}
