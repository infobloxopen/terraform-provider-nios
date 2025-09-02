// Retrieve a specific NS Group Delegation by filters
data "nios_dns_nsgroup_delegation" "get_ns_group_delegation_using_filters" {
  filters = {
    name = "example_ns_group_delegation"
  }
}

// Retrieve specific NS Group Delegations using Extensible Attributes
data "nios_dns_nsgroup_delegation" "get_ns_group_delegation_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all NS Group Delegations
data "nios_dns_nsgroup_delegation" "get_all_ns_group_delegations" {
}
