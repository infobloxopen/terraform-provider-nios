// Retrieve a specific Microsoft Super Scope by filters
data "nios_microsoft_mssuperscope" "get_microsoft_mssuperscope_using_filters" {
  filters = {
    name = "example_mssuperscope"
  }
}
// Retrieve specific Microsoft Super Scopes using Extensible Attributes
data "nios_microsoft_mssuperscope" "get_microsoft_mssuperscopes_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Microsoft Super Scopes
data "nios_microsoft_mssuperscope" "get_all_microsoft_mssuperscopes" {}
