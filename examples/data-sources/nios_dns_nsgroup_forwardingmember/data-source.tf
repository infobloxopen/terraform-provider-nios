// Retrieve a NS Group forwardingmember by filters
data "nios_dns_nsgroup_forwardingmember" "get_nsgroup_forwardingmember_using_filters" {
  filters = {
    name = "example_nsgroup_forwardingmember"
  }
}

// Retrieve a NS Group forwardingmember using Extensible Attributes
data "nios_dns_nsgroup_forwardingmember" "get_nsgroup_forwardingmember_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all NS Group forwardingmember
data "nios_dns_nsgroup_forwardingmember" "get_all_ns_group_forwardingmembers" {}
