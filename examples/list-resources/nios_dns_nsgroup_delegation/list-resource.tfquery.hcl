// List specific NS Group Delegations using filters
list "nios_dns_nsgroup_delegation" "list_ns_group_delegations_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example_ns_group_del"
    }
  }
}

// List specific NS Group Delegations using Extensible Attributes
list "nios_dns_nsgroup_delegation" "list_ns_group_delegations_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List NS Group Delegations with resource details included
list "nios_dns_nsgroup_delegation" "list_ns_group_delegations_with_resource" {
  provider         = nios
  include_resource = true
}
