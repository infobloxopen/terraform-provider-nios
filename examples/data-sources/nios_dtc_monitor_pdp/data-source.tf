// Retrieve a specific DTC pdp Monitor by filters
data "nios_dtc_monitor_pdp" "get_monitor_pdp_with_filter" {
  filters = {
    name = "example_dtc_monitor_pdp"
  }
}

// Retrieve specific DTC pdp Monitor using Extensible Attributes
data "nios_dtc_monitor_pdp" "get_monitor_pdp_with_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DTC pdp Monitors
data "nios_dtc_monitor_pdp" "get_all_pdp_monitors" {}
