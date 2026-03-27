// Create GMC Group with Basic Fields
resource "nios_grid_gmcgroup" "gmcgroup_with_basic_fields" {
  name = "example_gmcgroup_1"
}

// Create GMC Group with Additional Fields
resource "nios_grid_gmcgroup" "gmcgroup_with_additional_fields" {
  name                 = "example_gmcgroup_2"
  gmc_promotion_policy = "SEQUENTIALLY"
  comment              = "Example comment"
  members = [
    {
      "member" : "infoblox.member1"
    },
  ]
}
