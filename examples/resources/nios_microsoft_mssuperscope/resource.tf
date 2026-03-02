// Manage Microsoft Super Scope with Basic Fields
resource "nios_microsoft_mssuperscope" "microsoft_mssuperscope_basic" {
  name   = "example_mssuperscope"
  ranges = []
}

// Manage Microsoft Super Scope with Additional Fields
resource "nios_microsoft_mssuperscope" "microsoft_mssuperscope_with_additional_fields" {
  name   = "example_mssuperscope_additional"
  ranges = []

  //Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}