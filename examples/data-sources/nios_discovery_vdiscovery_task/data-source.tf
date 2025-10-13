// Retrieve a specific vDiscovery Task by filters
data "nios_discovery_vdiscovery_task" "get_vdiscoverytask_using_filters" {
  filters = {
    name = "AWS-vDiscovery-task"
  }
}

// Retrieve all vDiscovery Tasks
data "nios_discovery_vdiscovery_task" "get_all_vdiscoverytasks" {}
