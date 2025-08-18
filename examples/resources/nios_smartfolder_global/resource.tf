
// Create Smartfolder Global with Basic Fields (Required only)
resource "nios_smartfolder_global" "basic" {
  name = "example-global-smartfolder"
}

// Create Smartfolder Global with Additional Fields
resource "nios_smartfolder_global" "complete" {
  name    = "example-global-smartfolder-2"
  comment = "sample comment"

  group_bys = [{
    enable_grouping = true
    value           = "Availability zone"
    value_type      = "NORMAL"
  }]

  query_items = [{
    field_type = "NORMAL"
    name       = "type"
    op_match   = true
    operator   = "EQ"
    value = {
      value_string = "Network"
    }
    value_type = "ENUM"
  }]
}
