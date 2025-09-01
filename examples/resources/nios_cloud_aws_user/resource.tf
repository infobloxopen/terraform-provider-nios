// Create AWS User with Basic Fields
resource "nios_cloud_aws_user" "awsusertest" {
  access_key_id     = "AKIAexample1"
  account_id        = "337773173961"
  name              = "aws-user"
  secret_access_key = "S1JGWfwcZWEY+uduuujhSkfpyhxigL9A/uaJ96mY"
}

// Create AWS User with Additional Fields
resource "nios_cloud_aws_user" "test2" {
  access_key_id     = "AKIAexample2"
  account_id        = "337773173962"
  name              = "aws-user-2"
  secret_access_key = "S1JGWfwcZWEY+uduuujhSkfpyhxigL9A/uaJ96mZ"
  govcloud_enabled  = false
  nios_user_name    = "niosuser"
}
