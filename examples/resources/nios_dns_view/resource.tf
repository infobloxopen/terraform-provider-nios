// Create DNS View with Basic Fields
resource "nios_dns_view" "create_view" {
  name = "example_view"
}

// Create DNS View with additional fields
resource "nios_dns_view" "create_with_additional_fields" {
  name         = "example_custom_view"
  comment      = "DNS View"
  network_view = "default"

  //forwarders settings 
  use_forwarders = true
  forwarders     = ["10.192.81.23"]
  forward_only   = true

  //match client and destinations
  match_destinations = [
    {
      struct     = "addressac"
      address    = "192.168.0.45"
      permission = "ALLOW"
    },
  ]
  match_clients = [
    {
      struct     = "addressac"
      address    = "92.168.0.23"
      permission = "ALLOW"
    },
  ]

  //extensible attributes
  extattrs = {
    Site = "location-1"
  }
}
