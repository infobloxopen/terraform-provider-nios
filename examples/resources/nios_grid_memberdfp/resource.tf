// Create an Member dfp attribute definition with Basic Fields
resource "nios_grid_memberdfp" "memberdfp__with_basic_fields" {
  host_name       = "infoblox.member1"
  is_dfp_override = true
}

//Create an Member dfp attribute definition with additional fields
resource "nios_grid_memberdfp" "memberdfp__with_basic_fields" {
  host_name         = "infoblox.member1"
  is_dfp_override   = true
  dfp_forward_first = true
}
