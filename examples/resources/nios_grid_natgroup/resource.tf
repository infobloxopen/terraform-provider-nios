// Create NAT Group with Basic Fields
resource "nios_grid_natgroup" "natgroup_basic_fields" {
  name = "natgroup-basic"
}

// Create NAT Group with Additional Fields
resource "nios_grid_natgroup" "natgroup_with_additional_config" {
  name    = "natgroup-example"
  comment = "Example NAT Group for Grid communication"
}
