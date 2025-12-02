// Retrieve a specific DTC TCP Monitor by filters
data "nios_dtc_monitor_tcp" "get_tcp_monitor_with_filters" {
  filters = {
    name = "example_monitor_tcp"
  }
}

// Retrieve specific DTC TCP Monitor using Extensible Attributes
data "nios_dtc_monitor_tcp" "get_tcp_monitor_with_extattr_filters" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DTC TCP Monitors
data "nios_dtc_monitor_tcp" "get_all_tcp_monitors" {}
