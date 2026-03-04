// Retrieve a specific Threat Protection Grid Rule by filters
data "nios_threatprotection_grid_rule" "get_threatprotection_grid_rule_using_filters" {
  filters = {
    template = ""
  }
}

// Retrieve all Threat Protection Grid Rules
data "nios_threatprotection_grid_rule" "get_all_threatprotection_grid_rule" {}
