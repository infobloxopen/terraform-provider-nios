// Create a Parental Control Subscriber Site with Basic Fields
resource "nios_parentalcontrol_subscribersite" "subscriber_site_with_basic_fields" {
  name = "example_subscriber_site-1"
}

// Create a Parental Control Subscriber Site with Additional Fields
resource "nios_parentalcontrol_subscribersite" "subscriber_site_with_additional_fields" {
  name = "example_subscriber_site_2"

  // Additional Fields
  comment = "Example Subscriber Site"
  nas_gateways = [
    {
      "ip_address" = "12.1.1.1",
      "name"       = "nas_gateway_1",
      "send_ack" : false,
      "shared_secret" : "secret123"
    }
  ]

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}
