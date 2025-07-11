terraform {
  required_providers {
    nios = {
      source  = "infoblox-cto/nios"
      version = "1.0.0"
    }
  }
}

provider "nios" {
  nios_host_url = "https://172.28.82.248"
  nios_username = "admin"
  nios_password = "Infoblox@123"
}

resource "nios_dhcp_range" "test"{
    start_addr = "20.0.0.40"
    end_addr   = "20.0.0.50"
}