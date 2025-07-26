// Retrieve a specific DNS view by filters
data "nios_dns_view" "get_view_using_filters" {
  filters = {
    name = "example_custom_view"
  }
}

// Retrieve specific DNS view using Extensible Attributes
data "nios_dns_view" "get_view_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DNS views
data "nios_dns_view" "get_all_views" {}
