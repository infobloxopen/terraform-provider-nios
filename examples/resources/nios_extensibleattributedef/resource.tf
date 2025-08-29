//Create a extensibleattributedef with basic fields
resource "nios_grid_extensibleattributedef" "test" {
  name = "tf_test_ea_basic"
  type = "STRING"
}

//Create a extensibleattributedef with additional fields.
resource "nios_grid_extensibleattributedef" "test" {
  name = "tf_test_ea_additional"
  type = "STRING"
  allowed_object_types = ["NetworkContainer",
    "IPv6NetworkContainer",
    "Network", "IPv6Network",
  "FixedAddress", "IPv6FixedAddress"]
  comment = "Test extensible attribute definition"
}
