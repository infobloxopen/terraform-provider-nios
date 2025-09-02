// Retrieve a specific Extensible Attribute definition by filters
data "nios_grid_extensibleattributedef" "get_extensibleattributedef_with_filter" {
  filters = {
    name = "example_ea_1"
  }
}
// Retrieve all Extensible Attribute definitions
data "nios_grid_extensibleattributedef" "get_all_extensibleattributedef" {}
