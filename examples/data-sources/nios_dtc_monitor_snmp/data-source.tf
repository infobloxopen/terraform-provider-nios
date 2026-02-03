// Retrieve a specific DTC SNMP monitor by filters
data "nios_dtc_monitor_snmp" "get_monitor_snmp_using_filters" {
  filters = {
    name = "dtc_monitor_snmp"
  }
}

// Retrieve specific DTC SNMP monitors using Extensible Attributes
data "nios_dtc_monitor_snmp" "get_monitor_snmp_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DTC SNMP monitors 
data "nios_dtc_monitor_snmp" "get_all_snmp_monitors" {}
