terraform {
  required_providers {
    nios = {
      source  = "infoblox-cto/nios"
      version = "1.0.0"
    }
  }
}

provider "nios" {
  nios_host_url = "https://172.28.83.136"
  nios_username = "admin"
  nios_password = "Infoblox@123"
}
