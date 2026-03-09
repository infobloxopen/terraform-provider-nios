// Create a Parental Control Subscriber Record with Basic Fields
resource "nios_parentalcontrol_subscriberrecord" "subscriber_record_with_basic_fields" {
  ip_addr       = "12.1.1.1"
  ipsd          = "example_ipsd"
  localid       = "example_local_id"
  prefix        = 65
  site          = "example_site"
  subscriber_id = "example_subscriber_id"
}

// Create a Parental Control Subscriber Record with Additional Fields
resource "nios_parentalcontrol_subscriberrecord" "subscriber_record_with_additional_fields" {
  ip_addr                = "12.1.1.1"
  ipsd = "example_ipsd"
  localid          = "example_local_id"
  prefix = 65
  site = "example_site"
  subscriber_id = "example_subscriber_id"
}