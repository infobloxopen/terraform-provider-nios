// Retrieve a NS Group stub members by filters
data "nios_dns_nsgroup_stubmember" "get_ns_group_stubmember_using_filters" {
  filters = {
    name = "example_ns_group"
  }
}

// Retrieve a NS Group stub members using Extensible Attributes
data "nios_dns_nsgroup_stubmember" "get_ns_group_stubmember_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all NS Group stub members
data "nios_dns_nsgroup_stubmember" "get_all_ns_group_stubmembers" {}
