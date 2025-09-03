// Create NAT Group with basic fields
resource "nios_grid_natgroup" "natgroup_basic_fields" {
  name = "natgroup-basic"
}

// Create NAT Group with additional fields
resource "nios_grid_natgroup" "natgroup_with_additional_config" {
  name    = "natgroup-example"
  comment = "Test NAT Group for Grid communication"
}
