// Create awsrte53taskgroup with Basic Fields
resource "nios_cloud_awsrte53taskgroup" "awsrte53taskgroup_basic_fields" {
  name                        = "example_task_group"
  grid_member                 = "infoblox.localdomain"
  disabled                    = false
  sync_child_accounts         = false
  network_view_mapping_policy = "AUTO_CREATE"
}

// Create awsrte53taskgroup with Additional Fields
resource "nios_cloud_awsrte53taskgroup" "awsrte53taskgroup_additional_fields" {
  name                        = "example_task_group_2"
  grid_member                 = "infoblox.localdomain"
  disabled                    = false
  sync_child_accounts         = false
  network_view_mapping_policy = "AUTO_CREATE"

  task_list = [
    {
      name              = "test-task4"
      schedule_interval = "5"
      aws_user          = "awsuser/b25lLmF3c191c2VyJEFLSUFVNUpHWlRURVRTWEwyVEU0:AKIAU5JGZTTETSXL2TE4",
    },
    {
      name              = "test-task-17"
      schedule_interval = "10"
      aws_user          = "awsuser/b25lLmF3c191c2VyJEFLSUFVNUpHWlRURVRTWEwyVEU0:AKIAU5JGZTTETSXL2TE4",
    }
  ]
}

