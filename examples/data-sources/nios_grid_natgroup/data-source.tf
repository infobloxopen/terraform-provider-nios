// Retrieve a specific NAT Group by filters
data "nios_grid_natgroup" "get_natgroup_using_filters" {
  filters = {
    name = "test-natgroup"
  }
}

// Retrieve all NAT Groups
data "nios_grid_natgroup" "get_all_natgroups" {}
