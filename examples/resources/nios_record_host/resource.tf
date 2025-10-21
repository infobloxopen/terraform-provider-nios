# IP address allocation with static IPv4 address and location metadata
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

# IP address allocation with both IPv4 and IPv6 addresses
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

# IP address allocation using function call to retrieve ipv4addr
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

# Associate MAC address with host1's IP without DHCP configuration
resource "nios_ip_association" "association1" {
  ref                = nios_ip_allocation.allocation1.ref
  mac                = "12:00:43:fe:9a:8c"
  configure_for_dhcp = false
}

# Associate MAC and DUID with host2's IP with DHCP configuration enabled
resource "nios_ip_association" "association2" {
  ref                = nios_ip_allocation.allocation2.ref
  mac                = "12:43:fd:ba:9c:c9"
  duid               = "00:01:5f:3a:1b:2c:12:34:56:78:9a:bc"
  match_client       = "DUID"
  configure_for_dhcp = true
}

# Associate MAC address with host3's IP without DHCP configuration
resource "nios_ip_association" "association3" {
  ref                = nios_ip_allocation.allocation3.ref
  mac                = "12:00:43:fe:9a:8d"
  configure_for_dhcp = false
}
