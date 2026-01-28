#import {
#  to = nios_ip_allocation.allocation1
#  id = "record:host/ZG5zLmhvc3QkLl9kZWZhdWx0LmNvbS5leGFtcGxlLnNhbXBsZV9yZWNvcmQ:sample_record.example.com/default"
#}

resource "nios_ip_allocation" "allocation1" {
  name = "sample_record.example.com"
  view = "default"
  ipv4addrs = [
    {
      ipv4addr = "10.101.1.110"
    }
  ]
}

import {
  to = nios_ip_association.association1
  id = "record:host/ZG5zLmhvc3QkLl9kZWZhdWx0LmNvbS5leGFtcGxlLnNhbXBsZV9yZWNvcmQ:sample_record.example.com/default"
}

resource "nios_ip_association" "association1" {
  ref = nios_ip_allocation.allocation1.ref
}

terraform {
  required_providers {
    nios = {
      source  = "infobloxopen/nios"
      version = "0.0.1"
    }
  }
}

provider "nios" {
  nios_host_url = "https://172.28.83.204"
  nios_username = "admin"
  nios_password = "Infoblox@123"
}

resource "nios_ip_allocation" "all1" {
  name              = "host177.example.com"
  view              = "default"
#  aliases = ["alias23.example.com", "alias24.example.com"]
  comment = null
  configure_for_dns = true
  ipv4addrs = [
    {
      ipv4addr = "10.101.10.110"
    }
  ]
  extattrs = {
    Site = "location-6"
  }
  ddns_protected   = true
  device_description = "This is host1 device description"
  restart_if_needed   = true
  rrset_order = "random"
  use_cli_credentials = true
  snmp3_credential = {
    user = "NIOS_USER"
    authentication_protocol = "SHA"
    authentication_password = "PASShhj777l"
    privacy_protocol = "AES"
    privacy_password = "PASS28889k"
    comment = "Comment"
    credential_group = "default"
  }
  use_snmp3_credential = true
}
