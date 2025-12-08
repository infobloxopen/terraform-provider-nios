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

resource "nios_dtc_record_aaaa" "example" {
  ipv6addr = "2001:0db8:85a3:0000:0000:8a2e:0370:7134"
  dtc_server = "dtc_server.com"
}
