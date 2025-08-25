// Create Smartfolder Personal with Basic Fields (Required only)
resource "nios_smartfolder_personal" "basic" {
  name = "example-personal-smartfolder"
}

// Create Smartfolder Personal with Additional Fields
resource "nios_smartfolder_personal" "complete" {
  name    = "example-personal-smartfolder-2"
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
