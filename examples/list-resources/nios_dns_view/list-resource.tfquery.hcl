// List specific DNS Views using filters
list "nios_dns_view" "list_dns_views_using_filters" {
  provider = nios
  config {
    filters = {
      name = "default"
    }
  }
}

// List specific DNS Views using Extensible Attributes
list "nios_dns_view" "list_dns_views_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List DNS Views with resource details included
list "nios_dns_view" "list_dns_views_with_resource" {
  provider         = nios
  include_resource = true
}
