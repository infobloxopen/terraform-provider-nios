// Retrieve a specific AWS User by filters
data "nios_cloud_aws_user" "get_aws_user_by_filters" {
  filters = {
    name = "aws-user"
  }
}

// Retrieve all AWS Users
data "nios_cloud_aws_user" "get_all_aws_users" {}
