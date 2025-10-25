// Create an Auth zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_auth_zone" {
  fqdn        = "example.com"
  zone_format = "FORWARD"
  view        = "default"
  comment     = "Parent zone for AAAA records"
}

// Create an IPV6 network for function call (Required as Parent)
resource "nios_ipam_ipv6network" "example_ipv6_network" {
  network      = "2001:db8:abcd:12::/64"
  network_view = "default"
  comment      = "IPv6 network for AAAA record IP allocation"
}

// Create Record AAAA with Basic Fields
resource "nios_dns_record_aaaa" "create_record_aaaa_with_basic_fields" {
  name       = "example_record.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  ipv6addr   = "2002:1111::1401"
  view       = "default"
  depends_on = [nios_dns_zone_auth.parent_auth_zone]
}

// Create Record AAAA with additional fields
resource "nios_dns_record_aaaa" "create_record_aaaa_with_additional_fields" {
  name     = "example_record_with_ttl.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  ipv6addr = "2002:1111::1401"
  view     = "default"
  use_ttl  = true
  ttl      = 10
  comment  = "Example AAAA record"
  extattrs = {
    Site = "location-1"
  }
  depends_on = [nios_dns_zone_auth.parent_auth_zone]
}

// Create Record AAAA using function call to retrieve ipv6addr
resource "nios_dns_record_aaaa" "create_record_aaaa_with_func_call" {
  name = "example_record_with_func_call.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  func_call = {
    attribute_name  = "ipv6addr"
    object_function = "next_available_ip"
    result_field    = "ips"
    object          = "ipv6network"
    object_parameters = {
      network      = "2001:db8:abcd:12::/64"
      network_view = "default"
    }
  }
  view    = "default"
  comment = "AAAA record with function call"
  depends_on = [
    nios_dns_zone_auth.parent_auth_zone,
    nios_ipam_ipv6network.example_ipv6_network
  ]
}
