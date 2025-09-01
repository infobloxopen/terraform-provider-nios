// Retrieve a specific AWS User by filters
data "nios_cloud_aws_user" "awsuser_name" {
  filters = {
    name = "aws-user"
  }
}

// Retrieve all AWS Users
data "nios_cloud_aws_user" "awsuser_all" {}
