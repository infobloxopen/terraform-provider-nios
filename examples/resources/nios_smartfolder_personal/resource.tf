// Create a Personal Smarfolder with Basic Fields
resource "nios_smartfolder_personal" "create_smartfolder_personal_basic" {
  name = "example-personal-smartfolder"
}

// Create a Personal Smartfolder with Additional Fields
resource "nios_smartfolder_personal" "create_smartfolder_personal_additional" {
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
