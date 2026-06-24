// List specific Superhosts using filters
list "nios_ipam_superhost" "list_superhosts_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example_superhost"
    }
  }
}

// List specific Superhosts using Extensible Attributes
list "nios_ipam_superhost" "list_superhosts_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Superhosts with resource details included
list "nios_ipam_superhost" "list_superhosts_with_resource" {
  provider         = nios
  include_resource = true
}
