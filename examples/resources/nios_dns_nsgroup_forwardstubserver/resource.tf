// Create NS group Forward Stub Server with Basic Fields
resource "nios_dns_nsgroup_forwardstubserver" "nsgroup_forward_stub_server_with_basic_fields" {
  name = "example_ns_group_forward_stub_server"
  external_servers = [
    {
      name    = "example.com"
      address = "2.3.4.4"
    }
  ]
}

// Create NS Group Forward Stub Server with Additional Fields
resource "nios_dns_nsgroup_forwardstubserver" "nsgroup_forward_stub_server_with_additional_fields" {
  name = "example_ns_group_forward_stub_server1"
  external_servers = [
    {
      name    = "example.com"
      address = "2.3.4.4"
    }
  ]
  // Additional Fields
  comment = "Example NS Group Forward Stub Server"
  extattrs = {
    Site = "location-1"
  }
}
