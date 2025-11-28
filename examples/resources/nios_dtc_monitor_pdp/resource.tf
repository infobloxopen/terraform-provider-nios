// Create a DTC PDP Monitor with Basic Fields 
resource "nios_dtc_monitor_pdp" "monitor_pdp_basic_fields" {
  name = "example_dtc_monitor_pdp"
}

// Create a DTC PDP Monitor with Additional Fields
resource "nios_dtc_monitor_pdp" "monitor_pdp_additional_fields" {
  name    = "example_dtc_monitor_pdp_additional_fields"
  comment = "This is a DTC Monitor PDP"
  extattrs = {
    Site = "location-1"
  }
  interval   = 5
  port       = 80
  retry_down = 2
  retry_up   = 2
  timeout    = 10
}
