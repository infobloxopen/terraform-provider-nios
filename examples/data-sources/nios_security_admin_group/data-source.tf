// Retrieve a specific Admin Group by filters
data "nios_security_admin_group" "get_admin_group_using_filters" {
  filters = {
    name = "example_admin_group"
  }
}

// Retrieve specific Admin Groups using Extensible Attributes
data "nios_security_admin_group" "get_admin_groups_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Admin Groups
data "nios_security_admin_group" "get_all_admin_groups" {}
