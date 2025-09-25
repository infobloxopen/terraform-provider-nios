// Retrieve a specific FTP user by filters
data "nios_security_ftpuser" "get_ftpuser_using_filters" {
  filters = {
    username = "example_ftpuser"
  }
}
// Retrieve specific FTP users using Extensible Attributes
data "nios_security_ftpuser" "get_ftpusers_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}
// Retrieve all FTP users
data "nios_security_ftpuser" "get_all_ftpusers" {}
