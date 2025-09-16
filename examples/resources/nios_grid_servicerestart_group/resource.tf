//Create a Ruleset with Basic Fields
resource "nios_misc_ruleset" "ruleset_basic" {
  name = "example_grid_service_restart_group"
  service = "DNS"
}

//Create a Ruleset with Additional Fields
resource "nios_misc_ruleset" "ruleset_with_additional_fields" {
  name     = "example_grid_service_restart_group_additional"
  service = "DHCP"
  comment  = "Comment for GRID Service Restart Group"
  disabled = false

  extattrs = {
      Site = "location-1"
  }

  members = ["infoblox.member"]
  mode = "SEQUENTIAL"

  recurring_schedule = {
      
      }
}
