// Create an FTPuser with Basic Fields
resource "nios_security_ftpuser" "ftpusers_with_basic_fields" {
  username = "example_ftpuser"
  password = "example_password"
}
// Create an FTPuser with Additional Fields
resource "nios_security_ftpuser" "ftpusers_with_additional_fields" {
  username        = "example_ftpuser2"
  password        = "example_password"
  permission      = "RW"
  create_home_dir = false
  home_dir        = "/ftpusers"
  extattrs = {
    Site = "location-1"
  }
}


