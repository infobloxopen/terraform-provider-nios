// Retrieve a specific Network View by filters
data "nios_ipam_network_view" "get_network_views_using_filters" {
  filters = {
    name = "example_network_view"
  }
}

// Retrieve specific Network Views using Extensible Attributes
data "nios_ipam_network_view" "get_network_views_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Network Views
data "nios_ipam_network_view" "get_all_network_views" {}
