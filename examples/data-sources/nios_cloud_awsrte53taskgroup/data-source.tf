// Retrieve a specific awsrte53taskgroup by filters
data "nios_cloud_awsrte53taskgroup" "get_awsrte53taskgroup_by_filters" {
  filters = {
    name = "test-taskgroup"
  }
}

// Retrieve all awsrte53taskgroups
data "nios_cloud_awsrte53taskgroup" "all" {
}
