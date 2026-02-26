// Retrieve a specific Parental Control Blocking Policy by filters
data "nios_parentalcontrol_blockingpolicy" "get_parentalcontrol_blockingpolicy_using_filters" {
  filters = {
    name = "example_blockingpolicy"
  }
}

// Retrieve all Parental Control Blocking Policies
data "nios_parentalcontrol_blockingpolicy" "get_all_parentalcontrol_blockingpolicies" {}
