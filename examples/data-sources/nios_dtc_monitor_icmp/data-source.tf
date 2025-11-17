// Retrieve a specific DTC Monitor ICMP by filters
data "nios_dtc_monitor_icmp" "get_monitor_icmp_with_filter" {
  filters = {
    name = "example_monitor_icmp"
  }
}

// Retrieve specific DTC Monitor ICMP using Extensible Attributes
data "nios_dtc_monitor_icmp" "get_monitor_icmp_with_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DTC Monitor ICMP records
data "nios_dtc_monitor_icmp" "get_all_records" {}
