// Create Smartfolder Personal with Basic Fields (Required only)
resource "nios_smartfolder_personal" "basic" {
  name = "my-basic-smartfolder"
}

// Create Smartfolder Personal with Additional Fields
resource "nios_smartfolder_personal" "complete" {
  name    = "example-smartfolder"
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

# // Example with Multiple Group Bys
# resource "nios_smartfolder_personal" "multi_group_bys" {
#   name    = "multi-groupby-smartfolder"
#   comment = "Smartfolder with multiple group by configurations"

#   group_bys = [
#     {
#       enable_grouping = true
#       value           = "Availability zone"
#       value_type      = "NORMAL"
#     },
#     {
#       enable_grouping = false
#       value           = "Site"
#       value_type      = "EXTATTR"
#     }
#   ]
# }

# // Example with Multiple Query Items
# resource "nios_smartfolder_personal" "multi_query_items" {
#   name    = "multi-query-smartfolder"
#   comment = "Smartfolder with multiple query items"

#   query_items = [
#     {
#       field_type = "NORMAL"
#       name       = "type"
#       op_match   = true
#       operator   = "EQ"
#       value = {
#         value_string = "Network"
#       }
#       value_type = "ENUM"
#     },
#     {
#       field_type = "NORMAL"
#       name       = "type"
#       op_match   = true
#       operator   = "EQ"
#       value = {
#         value_string = "Zone"
#       }
#       value_type = "ENUM"
#     }
#   ]
# }