// Retrieve a specifi Vdiscoverytask by filters
data "nios_discovery_vdiscoverytask" "get_vdiscoverytask_using_filters" {
  filters = {
    name = "AWS-Vdiscoverytask"
  }
}

// Retrieve all Vdiscoverytasks
data "nios_discovery_vdiscoverytask" "get_all_vdiscoverytasks" {}
