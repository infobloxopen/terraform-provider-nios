// Retrieve a NS Group by filters
data "nios_dns_nsgroup" "get_ns_group_using_filters" {
  filters = {
    name = "example_ns_group"
  }
}

// Retrieve a NS Group using Extensible Attributes
data "nios_dns_nsgroup" "get_ns_group_using_extensible_attributes" {
  extattrfilters = {
    "Site" = "location-1"
  }
}

// Retrieve all NS Groups
data "nios_dns_nsgroup" "get_all_ns_groups" {}
