// Retrieve a specific Global Smartfolder by filters
data "nios_smartfolder_global" "get_global_smartfolder_using_filters" {
  filters = {
    name = "example-global-smartfolder"
  }
}

// Retrieve all Global Smartfolders
data "nios_smartfolder_global" "get_all_global_smartfolders" {}
