// List specific Network Views using filters
list "nios_ipam_network_view" "list_network_views_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example_network_view"
    }
  }
}

// List specific Network Views using Extensible Attributes
list "nios_ipam_network_view" "list_network_views_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Network Views with resource details included
list "nios_ipam_network_view" "list_network_views_with_resource" {
  provider         = nios
  include_resource = true
}
