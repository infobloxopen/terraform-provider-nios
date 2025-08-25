// Create awsrte53taskgroup with Basic Fields
resource "nios_cloud_awsrte53taskgroup" "test" {
  name                        = "test-taskgroup"
  grid_member                 = "infoblox.localdomain"
  disabled                    = false
  sync_child_accounts         = false
  network_view_mapping_policy = "AUTO_CREATE"
}

// Create awsrte53taskgroup with Advanced Fields
resource "nios_cloud_awsrte53taskgroup" "test_advanced" {
  name                        = "test-taskgroup-advanced"
  grid_member                 = "infoblox.localdomain"
  disabled                    = false
  sync_child_accounts         = false
  network_view_mapping_policy = "AUTO_CREATE"

  task_list {
    name              = "test-task4"
    schedule_interval = 5
    aws_user          = "awsuser/***:***"
  }

  task_list {
    name              = "test-task5"
    schedule_interval = 5
    aws_user          = "awsuser/***:***"
  }
}

