// Create an Admin Group with Basic Fields
resource "nios_security_admin_group" "admin_group_basic_fields" {
  name = "example_admin_group"
}

// Create an Admin Group with Additional Fields
resource "nios_security_admin_group" "admin_group_additional_fields" {
  name    = "example_admin_group2"
  comment = "Example Admin Group with additional fields"
  disable = true
  extattrs = {
    Site = "location-1"
  }
}
