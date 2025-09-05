// Create aws user with Basic Fields
resource "nios_cloud_aws_user" "aws_user_basic_fields" {
  access_key_id     = "AKIAexample1"
  account_id        = "337773173961"
  name              = "aws-user"
  secret_access_key = "S1JGWfwcZWESkfpyhxigL9A/u96mY"
}

// Create awsrte53taskgroup with Basic Fields
resource "nios_cloud_aws_route53_task_group" "awsrte53taskgroup_basic_fields" {
  name                        = "example_task_group"
  grid_member                 = "infoblox.localdomain"
  disabled                    = false
  sync_child_accounts         = false
  network_view_mapping_policy = "AUTO_CREATE"
}

// Create awsrte53taskgroup with Additional Fields
resource "nios_cloud_aws_route53_task_group" "awsrte53taskgroup_additional_fields" {
  name                        = "example_task_group_2"
  grid_member                 = "infoblox.localdomain"
  disabled                    = false
  sync_child_accounts         = false
  network_view_mapping_policy = "AUTO_CREATE"

  task_list = [
    {
      name              = "example-task2"
      schedule_interval = "5"
      aws_user          = nios_cloud_aws_user.aws_user_basic_fields.ref,
    },
    {
      name              = "example-task1"
      schedule_interval = "10"
      aws_user          = nios_cloud_aws_user.aws_user_basic_fields.ref,
    }
  ]
}

