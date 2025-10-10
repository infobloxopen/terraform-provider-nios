// Create a Notification REST Endpoint with Basic Fields
resource "nios_notification_rest_endpoint" "notification_rest_endpoint_with_basic_fields" {
  name                 = "example-notification-rest-endpoint-1"
  outbound_member_type = "GM"
  uri                  = "https://example.com"
}

// Create a Notification REST Endpoint with Additional Fields
resource "nios_notification_rest_endpoint" "notification_rest_endpoint_with_additional_fields" {
  name                 = "example-notification-rest-endpoint-2"
  outbound_member_type = "MEMBER"
  uri                  = "https://example.com"
  outbound_members = [
    "infoblox.grid_master_candidate1",
  ]

  // Additional Fields
  extattrs = {
    Site = "location-1"
  }
  username               = "example_username"
  password               = "example_password"
  server_cert_validation = "NO_VALIDATION"
  sync_disabled          = true
  template_instance = {
    parameters = [
      {
        name   = "SPECIAL1"
        syntax = "INT"
      },
      {
        name   = "SPECIAL2"
        syntax = "STR"
      },
      {
        name   = "SPECIAL3"
        syntax = "BOOL"
      }
    ]
    template = "REST API Template"
  }
  client_certificate_file = "<path-to-the-client-certificate-file>"
}
