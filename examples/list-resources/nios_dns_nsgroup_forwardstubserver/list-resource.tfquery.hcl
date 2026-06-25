// List DNS NS Group Forward Stub Servers using filters
list "nios_dns_nsgroup_forwardstubserver" "list_ns_group_forward_stub_server_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example_ns_group_forward_stub_server"
    }
  }
}

// List DNS NS Group Forward Stub Servers using Extensible Attributes
list "nios_dns_nsgroup_forwardstubserver" "list_ns_group_forward_stub_server_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List DNS NS Group Forward Stub Servers with resource details included
list "nios_dns_nsgroup_forwardstubserver" "list_ns_group_forward_stub_server_with_resource" {
  provider         = nios
  include_resource = true
}
