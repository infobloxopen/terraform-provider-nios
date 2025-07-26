// Retrieve a specific Network View by filters
data "nios_ipam_networkview" "get_record_using_filters" {
  filters = {
    name = "my_network_view"
  }
}
output "record_using_filters" {
  value = data.nios_ipam_networkview.get_record_using_filters
}


// Retrieve specific Network Views using Extensible Attributes
data "nios_ipam_networkview" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Network Views
data "nios_ipam_networkview" "get_all_records" {}

