// Retrieve a specific Admin Role by filters
data "nios_security_admin_role" "get_admin_role_using_filters" {
  filters = {
    name = "example_admin_role"
  }
}

// Retrieve specific Admin Roles using Extensible Attributes
data "nios_security_admin_role" "get_admin_roles_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Admin Roles
data "nios_security_admin_role" "get_all_admin_roles" {}
