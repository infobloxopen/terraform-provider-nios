// Retrieve a specific Global smartfolder by filters
data "nios_smartfolder_global" "get_global_smartfolder_using_filters" {
  filters = {
    name = "example-global-smartfolder"
  }
}

// Retrieve all Global smartfolders
data "nios_smartfolder_global" "get_all_global_smartfolders" {}
