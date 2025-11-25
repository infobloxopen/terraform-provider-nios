// Create a DTC monitor icmp with Basic Fields 
resource "nios_dtc_monitor_icmp" "example" {
  name = "example_icmp_monitor"
}

// Create a DTC monitor icmp with Additional Fields
resource "nios_dtc_monitor_icmp" "example2" {
  name    = "example_icmp_monitor2"
  comment = "DTC ICMP monitor creation"
  extattrs = {
    Site = "location-1"
  }
  interval   = 45
  timeout    = 10
  retry_up   = 2
  retry_down = 3
}
