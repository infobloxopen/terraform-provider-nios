// Retrieve a specific AWS User by filters
data "nios_cloud_awsuser" "awsuser_name" {
  filters = {
    name = "aws-user"
  }
}

// Retrieve all AWS Users
data "nios_cloud_awsuser" "awsuser_all" {}
