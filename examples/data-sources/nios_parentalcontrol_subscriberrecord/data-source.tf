// Retrieve a specific Parental Control Subscriber Record by filters
data "nios_parentalcontrol_subscriberrecord" "get_parentalcontrol_subscriberrecord_using_filters" {
  filters = {
    site = "example_subscriber_site"
  }
}
