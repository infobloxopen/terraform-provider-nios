// Create AWS User with Basic Fields
resource "nios_cloud_aws_user" "aws_user_basic_fields" {
  access_key_id     = "AKIAexample1"
  account_id        = "337773173961"
  name              = "aws-user"
  secret_access_key = "S1JGWfwcZWESkfpyhxigL9A/u96mY"
}

// Create AWS User with Additional Fields
resource "nios_cloud_aws_user" "aws_user_additional_fields" {
  access_key_id     = "AKIAexample2"
  account_id        = "337773173962"
  name              = "aws-user-2"
  secret_access_key = "S1JGWfwcZWkfpyhxigL9A/ua6mZ"
  govcloud_enabled  = false
  nios_user_name    = "niosuser"
}
