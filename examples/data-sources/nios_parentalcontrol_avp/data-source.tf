// Retrieve a specific Attribute Value Pair by filters
data "nios_parentalcontrol_avp" "get_attribute_value_pair_using_filter" {
  filters = {
    name = "example-avp-1"
  }
}

// Retrieve all Attribute Value Pairs
data "nios_parentalcontrol_avp" "get_all_attribute_value_pairs" {}
