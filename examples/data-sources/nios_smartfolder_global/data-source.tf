// Retrieve a specific smartfolder global by filters
data "nios_smartfolder_global" "get_smartfolder_using_filters" {
  filters = {
    name = "example-global-smartfolder"
  }
}

// Retrieve all smartfolder global records
data "nios_smartfolder_global" "get_all_smartfolders" {}
