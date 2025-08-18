// Retrieve a specific smartfolder global by filters
data "nios_smartfolder_global" "get_smartfolder_using_filters" {
  filters = {
    name = "example-global-smartfolder"
  }
}

// Retrieve all smartfolder global records
data "nios_smartfolder_global" "get_all_smartfolders" {}

output "name" {
  value = data.nios_smartfolder_global.get_smartfolder_using_filters

}
output "all_smartfolders" {
  value = data.nios_smartfolder_global.get_all_smartfolders
}