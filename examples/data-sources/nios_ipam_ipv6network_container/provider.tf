terraform {
  required_providers {
    nios = {
      source  = "infobloxopen/nios"
      version = "1.0.0"
    }
  }
}

provider "nios" {
  nios_username="admin"
  nios_password="Infoblox@123"
  nios_host_url="https://172.28.82.250"
}