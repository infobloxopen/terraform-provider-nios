// Create a BFD template with Basic Fields
resource "nios_misc_bfdtemplate" "bfd_template_with_basic_fields" {
  name = "example_bfd_name"
}

// Create a BFD template with Additional Fields
resource "nios_misc_bfdtemplate" "bfd_template_with_additional_fields" {
  name                  = "example_bfd_name_additional"
  authentication_key_id = 4
  authentication_type   = "MD5"
  detection_multiplier  = 5
  min_rx_interval       = 1000
  min_tx_interval       = 1000
}
