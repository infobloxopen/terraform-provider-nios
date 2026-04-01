// Retrieve a specific Microsoft Server by filters
data "nios_microsoft_msserver" "get_msserver_using_filters" {
  filters = {
    address = "10.10.0.1"
  }
}
// Retrieve specific Microsoft Server using Extensible Attributes
data "nios_microsoft_msserver" "get_msserver_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Microsoft Servers
data "nios_microsoft_msserver" "get_all_msservers" {}
