# Filter by an Attribute
data "nios_datasource_nios_RecordA" "example_by_attribute" {
  filters = {
    name = "example_recorda"
  }
}

# Filter by Extensible Attributes
data "nios_datasource_nios_RecordA" "example_by_ea" {
  extattrfilters = {
    Site = "us-west-1"
  }
}

# Get all existing objects
data "nios_datasource_nios_RecordA" "example_all" {}
