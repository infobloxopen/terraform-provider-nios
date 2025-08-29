// Create nios_misc_bfdtemplate with Basic Fields
resource "nios_misc_bfdtemplate" "test_bfdtemplate_basic" {
  name = "tf_test_bfd_name"
}

// Create nios_misc_bfdtemplate with additional fields
resource "nios_misc_bfdtemplate" "test_bfdtemplate_additional" {
  name                  = "tf_test_bfd_name_additional"
  authentication_key_id = 4
  authentication_type   = "MD5"
  detection_multiplier  = 5
  min_rx_interval       = 1000
  min_tx_interval       = 1000
}
