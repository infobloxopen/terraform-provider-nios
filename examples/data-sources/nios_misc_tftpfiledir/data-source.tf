// Retrieve a specific TFTP file/directory object by filters
data "nios_misc_tftpfiledir" "get_misc_tftpfiledir_using_filters" {
  filters = {
    name = "NAME_REPLACE_ME"
    type = "TYPE_REPLACE_ME"
  }
}

// Retrieve all TFTP file/directory objects
data "nios_misc_tftpfiledir" "get_all_misc_tftpfiledir" {}