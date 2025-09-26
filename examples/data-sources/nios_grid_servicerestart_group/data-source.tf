// Retrieve a specific GRID Service Restart Group by filters
data "nios_grid_servicerestart_group" "get_grid_servicerestart_group_using_filters" {
  filters = {
    name = "example_grid_service_restart_group"
  }
}

// Retrieve specific GRID Service Restart Group using Extensible Attributes
data "nios_grid_servicerestart_group" "get_grid_servicerestart_group_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all GRID Service Restart Groups
data "nios_grid_servicerestart_group" "get_all_grid_servicerestart_group" {}
