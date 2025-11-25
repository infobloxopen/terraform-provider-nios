// Retrieve a specific DTC Monitor HTTP by filters
data "nios_dtc_monitor_http" "get_dtc_monitor_with_filters" {
  filters = {
    name = "example-monitor-http"
  }
}

// Retrieve specific DTC HTTP Monitor using Extensible Attributes
data "nios_dtc_monitor_http" "get_dtc_monitor_with_extattr_filters" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DTC HTTP Monitor
data "nios_dtc_monitor_http" "get_all_dtc_http_monitors" {}
