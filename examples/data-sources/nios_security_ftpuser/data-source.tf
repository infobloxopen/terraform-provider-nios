// Retrieve a specific ftp user by filters
data "nios_security_ftpuser" "get_ftpuser_using_filters" {
  filters = {
    username = "example_ftpuser"
  }
}
// Retrieve specific ftp users using Extensible Attributes
data "nios_security_ftpuser" "get_ftpusers_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}
// Retrieve all ftp users
data "nios_security_ftpuser" "get_all_ftpusers" {}
