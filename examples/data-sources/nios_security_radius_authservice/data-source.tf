// Retrieve a specific Radius Authservice by filters
data "nios_security_radius_authservice" "get_radius_authservice_using_filters" {
  filters = {
    name = "radius_authservice1"
  }
}

// Retrieve all Radius Authservices
data "nios_security_radius_authservice" "get_all_radius_authservices" {}
