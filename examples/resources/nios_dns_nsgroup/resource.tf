// Create NS group with Basic Fields
resource "nios_dns_nsgroup" "create_ns_group" {
  name = "example_ns_group"
  grid_primary = [
    {
      name = "infoblox.localdomain"
    }
  ]
}

// Create NS Group with Additional Fields
resource "nios_dns_nsgroup" "create_ns_group_with_additional_fields" {
  name    = "example_ns_group_1"
  comment = "Example NS Group"

  grid_secondaries = [
    {
      name = "12.9.0.6",
    },
  ]
  external_primaries = [
    {
      name = "external.primary.1",
      address = "2.3.4.5",
    },
  ]
  use_external_primary = true
  extattrs = {
    Site = "location-1"
  }
}
