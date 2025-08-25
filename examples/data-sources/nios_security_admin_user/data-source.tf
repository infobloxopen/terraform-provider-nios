// Retrieve a specific Admin User by filters
data "nios_security_admin_user" "get_admin_user_using_filters" {
  filters = {
    name = "example_admin_user"
  }
}

// Retrieve specific Admin Users using Extensible Attributes
data "nios_security_admin_user" "get_admin_users_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Admin Users
data "nios_security_admin_user" "get_all_admin_users" {}
