// Create IPAM Network Views (Required as Parent)
resource "nios_ipam_network_view" "example_network_view_parent_1" {
  name = "example_network_view_1"
}

resource "nios_ipam_network_view" "example_network_view_parent_2" {
  name = "example_network_view_2"
}

// Manage Pxgrid Endpoint with Basic Fields
resource "nios_misc_pxgrid_endpoint" "misc_pxgrid_endpoint_basic" {
  address                 = "10.0.0.0"
  client_certificate_file = "${path.module}/../../../internal/testdata/nios_misc_pxgrid_endpoint/client.pem"
  name                    = "example_pxgrid_endpoint"
  outbound_member_type    = "GM"
  subscribe_settings = {
    enabled_attributes = ["ENDPOINT_PROFILE", "DOMAINNAME"]
  }
  publish_settings = {
    enabled_attributes = ["IPADDRESS"]
  }
  network_view = nios_ipam_network_view.example_network_view_parent_1.name
}

// Manage Pxgrid Endpoint with Additional Fields
resource "nios_misc_pxgrid_endpoint" "misc_pxgrid_endpoint_with_additional_fields" {
  address                 = "11.0.0.0"
  client_certificate_file = "${path.module}/../../../internal/testdata/nios_misc_pxgrid_endpoint/client.pem"
  name                    = "example_pxgrid_endpoint_2"
  outbound_member_type    = "MEMBER"
  outbound_members        = ["infoblox.grid_master_candidate1"]
  subscribe_settings = {
    enabled_attributes = ["ENDPOINT_PROFILE", "DOMAINNAME", "USERNAME"]
  }
  publish_settings = {
    enabled_attributes = ["IPADDRESS"]
  }
  log_level = "INFO"
  timeout   = 100

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
  network_view = nios_ipam_network_view.example_network_view_parent_2.name
}
