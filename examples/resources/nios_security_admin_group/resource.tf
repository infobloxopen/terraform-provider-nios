// Create an Admin Group with Basic Fields
resource "nios_security_admin_group" "admin_group_basic_fields" {
  name = "example_admin_group"
  comment = "Example Admin Group with basic fields"
  disable = true
  extattrs = {
    Site = "location-1"
  }
}

// Create an Admin Group with Additional Fields
resource "nios_security_admin_group" "admin_group_additional_fields" {
  name    = "example_admin_group2"
  admin_set_commands = {
    set_collect_old_logs = true
    set_core_files_quota = false
  }
  access_method = ["GUI", "API", "CLI"]
  admin_set_commands = {

  }
  admin_show_commands = {

  }
}
