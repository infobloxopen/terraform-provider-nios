// Retrieve a specific RIR Organization by filters
data "nios_rir_organization" "get_rir_organization_using_filters" {
  filters = {
    name = "example_rir_organization"
  }
}
// Retrieve specific RIR Organizations using Extensible Attributes
data "nios_rir_organization" "get_rir_organization_using_extensible_attributes" {
  extattrfilters = {
    "RIPE Email" = "support@infoblox.com"
  }
}

// Retrieve all RIR Organizations
data "nios_rir_organization" "get_all_rir_organizations" {}
