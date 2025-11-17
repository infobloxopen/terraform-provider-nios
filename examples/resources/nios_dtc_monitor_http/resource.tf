terraform {
  required_providers {
    nios = {
      source  = "infobloxopen/nios"
      version = "1.0.0"
    }
  }
}

provider "nios" {
  nios_host_url = "https://172.28.82.33"
  nios_username = "admin" 
  nios_password = "Infoblox@123" 
}

// Create a DTC Monitor HTTP with Basic Fields
resource "nios_dtc_monitor_http" "create_dtc_monitor_http" {
  name = "example-monitor-http"
  request = "POST / HTTP/1.1\n\n"
}
