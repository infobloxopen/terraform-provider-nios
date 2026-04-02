// Retrieve a specific GRID Member by filters
data "nios_grid_member" "get_grid_member_using_filters" {
  filters = {
    host_name = "infoblox.localdomain"
  }
}

// Retrieve specific GRID Members using Extensible Attributes
data "nios_grid_member" "get_grid_member_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all GRID Members
data "nios_grid_member" "get_all_grid_members" {}