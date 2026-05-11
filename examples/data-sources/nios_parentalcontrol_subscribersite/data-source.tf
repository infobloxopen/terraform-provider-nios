// Retrieve a specific Parental Control Subscriber Site by filters
data "nios_parentalcontrol_subscribersite" "get_subscriber_site_using_filters" {
  filters = {
    name = "example_subscriber_site"
  }
}

// Retrieve all Parental Control Subscriber Sites
data "nios_parentalcontrol_subscribersite" "get_all_subscriber_sites" {}
