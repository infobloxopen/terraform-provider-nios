//Retrieve a specific Upgradegroup by filters
data "nios_grid_upgradegroup" "get_upgradegroup_using_filters" {
  filters = {
    name = "upgradegroup-additional"
  }
}

//Retrieve all Upgradegroups
data "nios_grid_upgradegroup" "get_all_upgradegroups" {}
