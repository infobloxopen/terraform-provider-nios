// Retrieve a specific DTC SIP Monitor by filters
data "nios_dtc_monitor_sip" "get_dtc_sip_monitor_with_filter" {
  filters = {
    name = "example-monitor-sip"
  }
}

// Retrieve specific DTC SIP Monitors using Extensible Attributes
data "nios_dtc_monitor_sip" "get_dtc_sip_monitor_with_extattr_filter" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DTC SIP Monitors
data "nios_dtc_monitor_sip" "get_all_dtc_sip_monitors" {}
