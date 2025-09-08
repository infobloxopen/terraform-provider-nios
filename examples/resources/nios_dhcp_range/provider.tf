terraform {
  required_providers {
    nios = {
      source  = "infobloxopen/nios"
      version = "0.0.1"
    }
  }
}

provider "nios" {
  nios_host_url = "<NIOS_HOST_URL>"
  nios_username = "<NIOS_USERNAME>"
  nios_password = "<NIOS_PASSWORD>"
}
