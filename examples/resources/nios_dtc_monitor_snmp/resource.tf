// Create DTC Monitor snmp with Basic Fields
resource "nios_dtc_monitor_snmp" "create_monitor_snmp" {
  name = "dtc_monitor_snmp"
}

// Create DTC Monitor snmp with Additional Fields
resource "nios_dtc_monitor_snmp" "create_monitor_snmp_with_additional_fields" {
  name    = "dtc_monitor_snmp2"
  comment = "Example DTC Monitor SNMP"
  extattrs = {
    Site = "location-1"
  }
  oids = [
    {
      oid       = ".2"
      condition = "EXACT"
      first     = "10"
    },
    {
      oid = ".02"
    },
    {
      oid       = ".1"
      condition = "EXACT"
      first     = "20"
    }
  ]
  community  = "private"
  context    = "snmpv3-context"
  version    = "V3"
  user       = "admin"
  engine_id  = "0x2b060102010500"
  interval   = 20
  timeout    = 5
  port       = 10161
  retry_down = 2
  retry_up   = 3
}