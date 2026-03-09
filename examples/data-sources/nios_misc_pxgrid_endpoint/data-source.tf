// Retrieve a specific Pxgrid Endpoint by filters
data "nios_misc_pxgrid_endpoint" "get_misc_pxgrid_endpoint_using_filters" {
  filters = {
    address              = "10.0.0.0"
    name                 = "example_pxgrid_endpoint"
    outbound_member_type = "GM"
  }
}

// Retrieve specific Pxgrid Endpoints using Extensible Attributes
data "nios_misc_pxgrid_endpoint" "get_misc_pxgrid_endpoints_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all Pxgrid Endpoints
data "nios_misc_pxgrid_endpoint" "get_all_misc_pxgrid_endpoints" {}
