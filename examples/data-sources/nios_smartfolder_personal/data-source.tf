// Retrieve a specific Personal Smartfolder by filters
data "nios_smartfolder_personal" "get_personal_smartfolder_using_filters" {
  filters = {
    name = "example-personal-smartfolder"
  }
}

// Retrieve all Personal Smartfolders
data "nios_smartfolder_personal" "get_all_personal_smartfolders" {}
