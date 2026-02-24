// Manage DXL Endpoint  with Basic Fields
resource "nios_misc_dxl_endpoint" "misc_dxl_endpoint_basic" {
  client_certificate_file = "${path.module}/../../../internal/testdata/nios_misc_dxl_endpoint/client.pem"
  name                    = "example-dxl-endpoint"
  outbound_member_type    = "GM"
  brokers = [
    {
      host_name = "broker1.example.com",
      port      = 8443
    },
  ]
}

// Manage DXL Endpoint with Additional Fields
resource "nios_misc_dxl_endpoint" "misc_dxl_endpoint_with_additional_fields" {
  client_certificate_file = "${path.module}/../../../internal/testdata/nios_misc_dxl_endpoint/client.pem"
  name                    = "example-dxl-endpoint2"
  outbound_member_type    = "MEMBER"
  outbound_members        = ["infoblox.grid_master_candidate1"]
  brokers_import_file     = "${path.module}/../../../internal/testdata/nios_misc_dxl_endpoint/brokerlist.properties"
  template_instance = {
    template = "Version5_DXL_Session_Template"
  }

  log_level          = "INFO"
  wapi_user_name     = "admin"
  wapi_user_password = "password"

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}
