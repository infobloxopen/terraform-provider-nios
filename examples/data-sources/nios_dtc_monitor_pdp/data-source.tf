// Retrieve a specific DTC PDP Monitor by filters
data "nios_dtc_monitor_pdp" "get_monitor_pdp_with_filters" {
  filters = {
    name = "example_dtc_monitor_pdp"
  }
}

// Retrieve specific DTC PDP Monitor using Extensible Attributes
data "nios_dtc_monitor_pdp" "get_monitor_pdp_with_extattr_filters" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DTC PDP Monitors
data "nios_dtc_monitor_pdp" "get_all_dtc_pdp_monitors" {}
