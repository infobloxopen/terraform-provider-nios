// Retrieve a specific smartfolder personal by filters
data "nios_smartfolder_personal" "get_smartfolder_personal_using_filters" {
  filters = {
    name = "example-personal-smartfolder"
  }
}

// Retrieve all smartfolder personal records
data "nios_smartfolder_personal" "get_all_smartfolders_personal" {}
