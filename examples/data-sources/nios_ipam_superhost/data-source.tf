// Retrieve a specific IPAM Super Host by filters
data "nios_ipam_superhost" "get_ipam_super_host_using_filters" {
  filters = {
    name = "example_super_host"
  }
}
// Retrieve specific IPAM Super Hosts using Extensible Attributes
data "nios_ipam_superhost" "get_ipam_super_host_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all IPAM Super Hosts
data "nios_ipam_superhost" "get_all_ipam_superhosts" {}