// Create an Admin Role with Basic Fields
resource "nios_security_admin_role" "admin_role_basic_fields" {
  name = "example_admin_role"
}

// Create an Admin Role with additional fields
resource "nios_security_admin_role" "admin_role_additional_fields" {
  name    = "example_admin_role2"
  comment = "Example Admin Role with additional fields"
  disable = true
  extattrs = {
    Site = "location-1"
  }
}
