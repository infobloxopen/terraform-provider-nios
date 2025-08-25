// Retrieve a specific awsrte53taskgroup by filters
data "nios_cloud_awsrte53taskgroup" "test" {
  filters = {
    name = "test-taskgroup"
  }
}

// Retrieve all awsrte53taskgroup
data "nios_cloud_awsrte53taskgroup" "all" {
}
