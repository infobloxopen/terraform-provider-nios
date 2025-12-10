// Create an Auth Zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_zone" {
  fqdn = "example.com"
}

// Create IPv4 Reverse Mapping Zones (Required as Parent)
resource "nios_dns_zone_auth" "reverse_zone1" {
  fqdn        = "10.20.1.0/24"
  view        = "default"
  zone_format = "IPV4"
  comment     = "Reverse zone for 10.20.1.0/24 network"
}

resource "nios_dns_zone_auth" "reverse_zone2" {
  fqdn        = "22.0.0.0/24"
  view        = "default"
  zone_format = "IPV4"
  comment     = "Reverse zone for 22.0.0.0/24 network"
}

// Create an IPv6 Reverse Mapping Zone (Required as Parent)
resource "nios_dns_zone_auth" "reverse_zone_ipv6" {
  fqdn        = "2001::/64"
  view        = "default"
  zone_format = "IPV6"
  comment     = "Reverse zone for 2001::/64 network"
}

// Create an IPv4 Reverse MappingZone for function call network (Required for function call PTR)
resource "nios_dns_zone_auth" "reverse_zone3" {
  fqdn        = "85.85.0.0/16"
  view        = "default"
  zone_format = "IPV4"
  comment     = "Reverse zone for 85.85.0.0/16 network - required for function call PTR"
}

// Create Network for function call (Required for next_available_ip)
resource "nios_ipam_network" "func_call_network" {
  network      = "85.86.0.0/16"
  network_view = "default"
  comment      = "Network for PTR record function call"
}

// Create an IPv4 PTR record with Basic Fields
resource "nios_dns_record_ptr" "create_ptr_record_with_ipv4addr" {
  ptrdname = "example_record1.${nios_dns_zone_auth.parent_zone.fqdn}"
  ipv4addr = "10.20.1.2"
  view     = "default"
  extattrs = {
    Site = "location-1"
  }
  depends_on = [nios_dns_zone_auth.reverse_zone1]
}

// Create an IPv6 PTR record with Basic Fields
resource "nios_dns_record_ptr" "create_ptr_record_with_ipv6addr" {
  ptrdname = "example_record2.${nios_dns_zone_auth.parent_zone.fqdn}"
  ipv6addr = "2001::123"
  view     = "default"
  extattrs = {
    Site = "location-2"
  }
  depends_on = [nios_dns_zone_auth.reverse_zone_ipv6]
}

// Create an IPv4 PTR record by name with Basic Fields
resource "nios_dns_record_ptr" "create_ptr_record_with_name" {
  ptrdname = "example_record3.${nios_dns_zone_auth.parent_zone.fqdn}"
  name     = "11.0.0.22.in-addr.arpa"
  view     = "default"
  extattrs = {
    Site = "location-3"
  }
  depends_on = [nios_dns_zone_auth.reverse_zone2]
}

// Create an IPv4 PTR record by name with Additional Fields
resource "nios_dns_record_ptr" "create_ptr_record_with_additional_fields" {
  ptrdname = "example_record4.${nios_dns_zone_auth.parent_zone.fqdn}"
  name     = "12.0.0.22.in-addr.arpa"

  // Additional Fields
  view    = "default"
  use_ttl = true
  ttl     = 10
  creator = "DYNAMIC"
  comment = "Example PTR record"

  // Extensible Attributes
  extattrs = {
    Site = "location-4"
  }
  depends_on = [nios_dns_zone_auth.reverse_zone2]
}

// Create an PTR record using function call to retrieve ipv4addr
resource "nios_dns_record_ptr" "create_ptr_record_with_func_call" {
  ptrdname = "example_func_call.${nios_dns_zone_auth.parent_zone.fqdn}"
  func_call = {
    attribute_name  = "ipv4addr"
    object_function = "next_available_ip"
    result_field    = "ips"
    object          = "network"
    object_parameters = {
      network      = "85.85.0.0/16"
      network_view = "default"
    }
  }
  view    = "default"
  comment = "Updated comment"
  depends_on = [
    nios_ipam_network.func_call_network,
    nios_dns_zone_auth.reverse_zone3
  ]
}

// Create an IPV4 reverse mapping zone (Required as Parent)
resource "nios_dns_zone_auth" "create_zone1" {
  fqdn        = "60.0.0.0/24"
  view        = "default"
  zone_format = "IPV4"
}

// Create an IPv4 PTR record by name with arpa notation
resource "nios_dns_record_ptr" "create_ptr_record_with_ipv4_arpa" {
  name     = "5.${nios_dns_zone_auth.create_zone1.display_domain}"
  ptrdname = "0.0.192.in-addr"
  view     = "default"
  extattrs = {
    Site = "location-3"
  }
}

// Create an IPV6 reverse mapping zone (Required as Parent)
resource "nios_dns_zone_auth" "create_zone2" {
  fqdn        = "2002:1100::/64"
  view        = "default"
  zone_format = "IPV6"
  extattrs = {
    Site = "location-1"
  }
}

// Create an IPv6 PTR record by name with arpa notation
resource "nios_dns_record_ptr" "create_ptr_record_with_ipv6_arpa" {
  ptrdname = "example_record.example.com"
  name     = "7.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.${nios_dns_zone_auth.create_zone2.display_domain}"
  view     = "default"
  extattrs = {
    Site = "location-2"
  }
}
