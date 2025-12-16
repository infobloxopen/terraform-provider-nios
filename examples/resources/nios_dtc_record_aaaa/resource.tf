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

// Create a DTC Server (Required as parent)
resource "nios_dtc_server" "create_dtc_server" {
  name = "example-server"
  host = "2.3.3.4"
}

// Create a DTC Record A with basic fields
resource "nios_dtc_record_aaaa" "record_aaaa_with_basic_fields" {
  dtc_server = nios_dtc_server.create_dtc_server.name
  ipv6addr   = "2001:0db8:85a3:0000:0000:8a2e:0370:6335"
}

# // Create a DTC Record A with additional fields
# resource "nios_dtc_record_aaaa" "record_a_with_additional_fields" {
#   dtc_server = nios_dtc_server.create_dtc_server.name
#   ipv6addr   = "2001:db8:85a3::8a2e:370:7337"
#   comment    = "Creating DTC Record AAAA"
#   disable    = false
#   ttl        = 20
#   use_ttl    = true
# }
