// Retrieve a specific grid Gmcgroup by filters
data "nios_grid_gmcgroup" "get_grid_gmcgroup_using_filters" {
  filters = {
    name = "example_gmcgroup_2"
  }
}

// Retrieve all grid Gmcgroup
data "nios_grid_gmcgroup" "get_all_grid_gmcgroup" {}
