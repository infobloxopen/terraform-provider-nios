// List DNS NS Group Stub Members using filters
list "nios_dns_nsgroup_stubmember" "list_ns_group_stub_member_using_filters" {
  provider = nios
  config {
    filters = {
      name = "stubmember1"
    }
  }
}

// List DNS NS Group Stub Members using Extensible Attributes
list "nios_dns_nsgroup_stubmember" "list_ns_group_stub_member_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List DNS NS Group Stub Members with resource details included
list "nios_dns_nsgroup_stubmember" "list_ns_group_stub_member_with_resource" {
  provider         = nios
  include_resource = true
}
