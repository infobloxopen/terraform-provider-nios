// List specific NS Groups using filters
list "nios_dns_nsgroup" "list_ns_groups_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example_ns_group"
    }
  }
}

// List specific NS Groups using Extensible Attributes
list "nios_dns_nsgroup" "list_ns_groups_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List NS Groups with resource details included
list "nios_dns_nsgroup" "list_ns_groups_with_resource" {
  provider         = nios
  include_resource = true
}
