// Create an Upgrade Group with Basic fields
resource "nios_grid_upgradegroup" "upgradegroup_with_basic_fields" {
  name = "upgradegroup-basic"
}

// Create an Upgrade Group with Additional fields
resource "nios_grid_upgradegroup" "upgradegroup_with_additional_fields" {
  name    = "upgradegroup-additional"
  comment = "This is a sample comment"
  members = [
    {
      member = "infoblox.10.38.23.8"
    },
  ]
  distribution_dependent_group = "distribution-dependent-group-1"
  distribution_policy          = "SEQUENTIALLY"
  upgrade_dependent_group      = "upgrade-dependent-group-1"
  upgrade_policy               = "SIMULTANEOUSLY"
  distribution_time            = 172834354
  upgrade_time                 = 172834355
}
