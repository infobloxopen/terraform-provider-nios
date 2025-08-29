
// Create Global Smartfolder with Basic Fields
resource "nios_smartfolder_global" "create_smartfolder_global_basic" {
  name = "example-global-smartfolder"
}

// Create Global Smartfolder with Additional Fields
resource "nios_smartfolder_global" "create_smartfolder_global_additional" {
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
