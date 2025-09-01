//Retrieve a specific Upgrade Group by filters
data "nios_grid_upgradegroup" "get_upgradegroup_using_filters" {
  filters = {
    name = "upgradegroup-additional"
  }
}

//Retrieve all Upgrade Groups
data "nios_grid_upgradegroup" "get_all_upgradegroups" {}
