// Retrieve a specific Shared Record Group by filters
data "nios_dns_sharedrecordgroup" "get_sharedrecordgroup_using_filters" {
  filters = {
    name = "example-shared-record-group"
  }
}

// Retrieve specific Shared Record Groups using Extensible Attributes
data "nios_dns_sharedrecordgroup" "get_sharedrecordgroups_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Shared Record Groups
data "nios_dns_sharedrecordgroup" "get_all_sharedrecordgroup" {}
