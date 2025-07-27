// Retrieve a specific Network View by filters
data "nios_ipam_network_view" "get_network_view_using_filters" {
  filters = {
    name = "my_network_view"
  }
}

// Retrieve specific Network Views using Extensible Attributes
data "nios_ipam_network_view" "get_network_view_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Network Views
data "nios_ipam_network_view" "get_all_network_views" {}
