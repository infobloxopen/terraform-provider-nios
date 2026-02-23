// Retrieve a specific DXL Endpoint by filters
data "nios_misc_dxl_endpoint" "get_dxl_endpoint_using_filters" {
  filters = {
    name                 = "example-dxl-endpoint"
    outbound_member_type = "GM"
  }
}

// Retrieve specific DXL Endpoints using Extensible Attributes
data "nios_misc_dxl_endpoint" "get_dxl_endpoint_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DXL Endpoints
data "nios_misc_dxl_endpoint" "get_all_dxl_endpoints" {}
