terraform {
  required_providers {
    nios = {
      source  = "Infoblox-CTO/nios"
      version = "1.0.0"
    }
  }
}

provider "nios" {
    nios_username = "admin"
    nios_password = "Infoblox@123"
    nios_host_url="https://172.28.82.250"
    delete_non_terraform_resources = true
    remove_records_if_found = "A"
}
