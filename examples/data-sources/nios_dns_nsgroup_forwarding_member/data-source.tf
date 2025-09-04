// Retrieve an NS Group Forwarding Member by filters
data "nios_dns_nsgroup_forwarding_member" "get_nsgroup_forwarding_member_using_filters" {
  filters = {
    name = "example_nsgroup_forwarding_member"
  }
}

// Retrieve an NS Group Forwarding Member using Extensible Attributes
data "nios_dns_nsgroup_forwarding_member" "get_nsgroup_forwarding_member_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all NS Group Forwarding Members
data "nios_dns_nsgroup_forwarding_member" "get_all_ns_group_forwarding_members" {}
