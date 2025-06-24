terraform {
    required_providers {
        nios = {
            source  = "infoblox-cto/nios"
            version = "1.0.0"
        }
    }
}

provider "nios" {
    nios_host_url="https://172.28.83.91"
    nios_username="admin"
    nios_password="Infoblox@123"
}


resource "nios_dns_record_a" "create_record" {
  name     = "example_test71.example.com"
  ipv4addr = "10.20.1.2"
  view     = "default"
  comment = ""
}

resource "nios_dns_record_a" "create_record1" {
  name     = "example_test72.example.com"
  ipv4addr = "10.20.1.2"
  view     = "default"
  comment = ""
}