// Retrieve a specific extensibleattributedef by name
data "nios_grid_extensibleattributedef" "test" {
  filters = {
    name = "tf_test_ea_basic"
  }
}

// Retrieve all extensibleattributedef
data "nios_grid_extensibleattributedef" "test_additional_fields" {
}
