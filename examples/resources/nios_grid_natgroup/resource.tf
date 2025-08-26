// Create NAT Group with basic fields
resource "nios_grid_natgroup" "basic_natgroup" {
  name = "test-natgroup"
}

// Create NAT Group with additional fields
resource "nios_grid_natgroup" "natgroup_with_additional_config" {
  name    = "test-natgroup"
  comment = "Test NAT Group for Grid communication"
}
