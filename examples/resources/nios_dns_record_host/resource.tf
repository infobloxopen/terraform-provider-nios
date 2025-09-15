# Create a DNS host record with static IPv4 address and location metadata
resource "nios_dns_record_host" "allocation1" {
  name = "host1.example.com"
  view = "default"
  ipv4addrs = [
    {
      ipv4addr = "10.101.1.110"
    }
  ]
  extattrs = {
    Site = "location-1"
  }
}

# Create a dual-stack DNS host record with both IPv4 and IPv6 addresses
resource "nios_dns_record_host" "allocation2" {
  name = "host2.example.com"
  view = "default"
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

# Create a DNS host record with dynamically allocated IPv4 address from network pool
resource "nios_dns_record_host" "allocation3" {
  name = "host3.example.com"
  view = "default"

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
      }
    }
  ]
}

# Associate MAC address with host1's IP without DHCP configuration
resource "nios_dhcp_ip_association" "association1" {
  ref                = nios_dns_record_host.allocation1.ref
  mac                = "12:00:43:fe:9a:8c"
  configure_for_dhcp = false
  internal_id        = nios_dns_record_host.allocation1.internal_id
}

# Associate MAC and DUID with host2's IP with DHCP configuration enabled
resource "nios_dhcp_ip_association" "association2" {
  ref                = nios_dns_record_host.allocation2.ref
  mac                = "12:43:fd:ba:9c:c9"
  duid               = "00:43:d2:0a:11:e6"
  configure_for_dhcp = true
  internal_id        = nios_dns_record_host.allocation2.internal_id
}
