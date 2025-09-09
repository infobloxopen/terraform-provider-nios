// Retrieve a specific NS Group Forward Stub Server by filters
data "nios_dns_nsgroup_forwardstubserver" "get_nsgroup_forward_stub_server_using_filters" {
  filters = {
    name = "example_ns_group_forward_stub_server"
  }
}

// Retrieve specific NS Group Forward Stub Servers using Extensible Attributes
data "nios_dns_nsgroup_forwardstubserver" "get_nsgroup_forward_stub_server_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all NS Groups Forward Stub Servers
data "nios_dns_nsgroup_forwardstubserver" "get_nsgroup_forward_stub_server_in_specific_view" {
}
