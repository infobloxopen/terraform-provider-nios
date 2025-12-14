// Retrieve a specific DHCP Fingerprint Filter by filters
data "nios_dhcp_filterfingerprint" "get_dhcp_filterfingerprint_using_filters" {
  filters = {
    name = "example_filterfingerprint_1"
  }
}
// Retrieve specific DHCP Fingerprint Filters using Extensible Attributes
data "nios_dhcp_filterfingerprint" "get_dhcp_filterfingerprins_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DHCP Fingerprint Filters
data "nios_dhcp_filterfingerprint" "get_all_dhcp_filterfingerprints" {}
