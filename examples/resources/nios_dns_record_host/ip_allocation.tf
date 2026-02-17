# Manage IP address allocation with static IPv4 address and location metadata
resource "nios_ip_allocation" "allocation1" {
  name              = "host1.example.com"
  view              = "default"
  configure_for_dns = true
  ipv4addrs = [
    {
      ipv4addr = "10.101.1.110"
    }
  ]
  extattrs = {
    Site = "location-1"
  }
}

# Manage IP address allocation with both IPv4 and IPv6 addresses
resource "nios_ip_allocation" "allocation2" {
  name              = "host2.example.com"
  view              = "default"
  configure_for_dns = true
  ipv4addrs = [
    {
      ipv4addr = "10.101.1.112"
    }
  ]
  ipv6addrs = [
    {
      ipv6addr = "2002:1f93::12:2"
    }
  ]
  extattrs = {
    Site = "location-1"
  }
}

# Manage IP address allocation using function call to retrieve ipv4addr
resource "nios_ip_allocation" "allocation3" {
  name              = "host3.example.com"
  view              = "default"
  configure_for_dns = true
  ipv4addrs = [
    {
      func_call = {
        attribute_name  = "ipv4addr"
        object_function = "next_available_ip"
        result_field    = "ips"
        object          = "network"
        object_parameters = {
          network      = "10.10.0.0/16"
          network_view = "default"
        }
        parameters = {
          exclude = jsonencode(["10.10.0.1", "10.10.0.2"]),
        }
      }
    }
  ]
  extattrs = {
    Site = "location-1"
  }
}