// Retrieve a specific Named ACL by filters
data "nios_acl_namedacl" "get_namedacl_using_filters" {
  filters = {
    name = "example-network-acl"
  }
}

// Retrieve specific Named ACL using Extensible Attributes
data "nios_acl_namedacl" "get_namedacl_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Named ACLs
data "nios_acl_namedacl" "get_all_namedacls" {}
