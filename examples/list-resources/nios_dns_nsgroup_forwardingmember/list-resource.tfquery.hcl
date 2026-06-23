// List DNS NS Group Forwarding Members using filters
list "nios_dns_nsgroup_forwardingmember" "list_ns_group_forwarding_member_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example_nsgroup_forwarding_member"
    }
  }
}

// List DNS NS Group Forwarding Members using Extensible Attributes
list "nios_dns_nsgroup_forwardingmember" "list_ns_group_forwarding_member_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List DNS NS Group Forwarding Members with resource details included
list "nios_dns_nsgroup_forwardingmember" "list_nsgroups_with_resource" {
  provider         = nios
  include_resource = true
}
