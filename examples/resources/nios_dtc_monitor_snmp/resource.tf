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

resource "nios_dtc_monitor_snmp" "create_monitor" {
  name = "dtc_monitor2"
  oids = [
    {
      oid       = ".2"
      condition = "EXACT"
      first     = "10"
    },
    {
      oid = ".02"
    },
    {
      oid       = ".1"
      condition = "EXACT"
      first     = "20"
    }
  ]
}