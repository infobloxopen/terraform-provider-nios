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
      name              = "example-task4"
      schedule_interval = "5"
      aws_user          = "awsuser/c25lLmF3c19HWlRURVRTWEwyV1EU0:KIAU5JXL2TE4",
    },
    {
      name              = "example-task17"
      schedule_interval = "10"
      aws_user          = "awsuser/c25lLmF3c19FVNUpURVRTWEwy1VEU0:IAU5SXL2TE4",
    }
  ]
}

