// Retrieve a specific awsrte53taskgroup by filters
data "nios_cloud_aws_route53_task_group" "get_awsrte53taskgroup_by_filters" {
  filters = {
    name = "example_task_group"
  }
}

// Retrieve all awsrte53taskgroups
data "nios_cloud_aws_route53_task_group" "get_all_awsrte53taskgroups" {}
