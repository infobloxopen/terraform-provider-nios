//Create an Upgradegroup with Basic fields

resource "nios_grid_upgradegroup" "upgradegroup_with_basic_fields" {
  name = "upgradegroup-basic"
}

//Create an Upgradegroup with Additional fields
resource "nios_grid_upgradegroup" "upgradegroup_with_additional_fields" {
  name    = "upgradegroup-additional"
  comment = "This is a sample comment"
  members = [
    {
      member = "infoblox.172_28_82_185"
    },
    # {
    #   member = "infoblox.172.28.82.186"
    # }
  ]
  distribution_dependent_group = "distribution-dependent-group-1"
  distribution_policy          = "SEQUENTIALLY"
  upgrade_dependent_group      = "upgrade-dependent-group-1"
  upgrade_policy               = "SIMULTANEOUSLY"
  distribution_time            = 172834354
  upgrade_time                 = 172834355
}
