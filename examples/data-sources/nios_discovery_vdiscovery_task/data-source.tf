// Retrieve a specific Vdiscovery Task by filters
data "nios_discovery_vdiscovery_task" "get_vdiscoverytask_using_filters" {
  filters = {
    name = "AWS-Vdiscoverytask"
  }
}

// Retrieve all Vdiscovery Tasks
data "nios_discovery_vdiscovery_task" "get_all_vdiscoverytasks" {}
