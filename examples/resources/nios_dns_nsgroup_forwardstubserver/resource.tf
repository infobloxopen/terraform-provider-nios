// Create NS group Forward Stub Server with Basic Fields
resource "nios_dns_nsgroup_forwardstubserver" "create_ns_group_forward_stub_server" {
  name = "example_ns_group_forward_stub_server"
  external_servers = [
    {
      name    = "infoblox.localdomain"
      address = "2.3.4.4"
    }
  ]
}

// Create NS Group Forward Stub Server with Additional Fields
resource "nios_dns_nsgroup_forwardstubserver" "create_ns_group_forward_stub_server_with_additional_fields" {
  name    = "example_ns_group_forward_stub_server1"
  comment = "Example NS Group Forward Stub Server"
  external_servers = [
    {
      name    = "infoblox.localdomain"
      address = "2.3.4.4"
    }
  ]
  //extensible attributes
  extattrs = {
    Site = "location-1"
  }
}