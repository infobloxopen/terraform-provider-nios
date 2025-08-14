// Retrieve a specific smartfolder personal by filters
data "nios_smartfolder_personal" "get_smartfolder_using_filters" {
  filters = {
    name = "my-basic-smartfolder"
  }
}

// Retrieve all smartfolder personal records
data "nios_smartfolder_personal" "get_all_smartfolders" {}

# // Retrieve smartfolder personal by name only
# data "nios_smartfolder_personal" "get_smartfolder_by_name" {
#   filters = {
#     name = "my-complete-smartfolder"
#   }
# }

# // Retrieve smartfolders with specific query items
# data "nios_smartfolder_personal" "get_smartfolder_with_query_items" {
#   filters = {
#     name = "multi-query-smartfolder"
#   }
# }

output "name" {
  value = data.nios_smartfolder_personal.get_smartfolder_using_filters

}
output "complete_smartfolder_name" {
  value = data.nios_smartfolder_personal.get_all_smartfolders
}