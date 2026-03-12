// Retrieve a specific TFTP file/directory object by filters
data "nios_misc_tftpfiledir" "get_misc_tftpfiledir_using_filters" {
  filters = {
    name = "example-tftpfiledir"
  }
}

// Retrieve all TFTP file/directory objects
data "nios_misc_tftpfiledir" "get_all_misc_tftpfiledir" {}
