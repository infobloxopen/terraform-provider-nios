// Retrieve a specific DTC ICMP Monitor by filters
data "nios_dtc_monitor_icmp" "get_monitor_icmp_with_filter" {
  filters = {
    name = "example_monitor_icmp"
  }
}

// Retrieve specific DTC ICMP Monitor using Extensible Attributes
data "nios_dtc_monitor_icmp" "get_monitor_icmp_with_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DTC ICMP Monitors
data "nios_dtc_monitor_icmp" "get_all_dtc_monitor_icmp" {}
