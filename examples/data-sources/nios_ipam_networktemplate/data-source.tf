// Retrieve a specific IPAM Network Template using filters
data "nios_ipam_networktemplate" "get_ipam_networktemplate_using_filters" {
  filters = {
    name = "example_network_template"
  }
}

// Retrieve specific IPAM Network Templates using Extensible Attributes
data "nios_ipam_networktemplate" "get_ipam_networktemplates_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all IPAM network Templates
data "nios_ipam_networktemplate" "get_all_ipam_networktemplates" {}
