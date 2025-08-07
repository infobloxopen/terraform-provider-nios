resource "nios_security_admin_user" "admin_user" {
  name = "example_admin_user"
  password = "Example-password1!1223a"
  admin_groups = ["admin-group"]
}