// Create an Extensible Attribute definition with Basic Fields
resource "nios_grid_extensibleattributedef" "extensibleattributedef_basic_fields" {
  name = "example_ea_1"
  type = "STRING"
}

// Create an Extensible Attribute definition with Additional Fields.
resource "nios_grid_extensibleattributedef" "extensibleattributedef_additional_fields" {
  name = "example_ea_2"
  type = "STRING"
  allowed_object_types = ["NetworkContainer",
    "IPv6NetworkContainer",
    "Network", "IPv6Network",
  "FixedAddress", "IPv6FixedAddress"]
  comment = "Extensible attribute definition"
}
