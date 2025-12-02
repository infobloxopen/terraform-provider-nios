// Create a DTC monitor TCP with Basic Fields 
resource "nios_dtc_monitor_tcp" "dtc_monitor_tcp_basic_fields" {
  name = "example_tcp_monitor"
  port = 8080
}

// Create a DTC monitor TCP with Additional Fields
resource "nios_dtc_monitor_tcp" "dtc_monitor_tcp_additional_fields" {
  name    = "example_tcp_monitor2"
  port    = 9090
  comment = "DTC TCP monitor creation"
  extattrs = {
    Site = "location-1"
  }
  interval   = 45
  timeout    = 10
  retry_up   = 2
  retry_down = 3
}
