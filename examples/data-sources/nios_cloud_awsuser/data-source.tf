// Retrieve a specific AWS User by filters
data "nios_cloud_awsuser" "awsuser_name" {
  filters = {
    nios_user_name = "AWS_USER_NAME"
  }
}

// Retrieve all AWS Users
data "nios_cloud_awsuser" "awsuser_all" {}
