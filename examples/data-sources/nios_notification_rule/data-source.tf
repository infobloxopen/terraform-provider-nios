// Retrieve a specific Notification Rule by filters
data "nios_notification_rule" "get_notification_rule_using_filters" {
  filters = {
    name = "example_notification_rule"
  }
}

// Retrieve all Notification Rules
data "nios_notification_rule" "get_all_notification_rules" {}
