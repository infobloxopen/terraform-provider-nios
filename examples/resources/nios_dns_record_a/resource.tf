// Create an Auth zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_auth_zone" {
  fqdn        = "example_auth.com"
  zone_format = "FORWARD"
  view        = "default"
  comment     = "Parent zone for A records"
}

// Create network for function call (required as parent)
resource "nios_ipam_network" "example_network" {
  network      = "85.85.0.0/16"
  network_view = "default"
  comment      = "Network for A record IP allocation"
}

// Create Record A with Basic Fields
resource "nios_dns_record_a" "create_record_a" {
  name     = "a-record.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  ipv4addr = "10.20.1.2"
  view     = "default"
  extattrs = {
    Site = "location-1"
  }
}

// Create Record A with additional fields
resource "nios_dns_record_a" "create_record_a_with_additional_fields" {
  name     = "name.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  ipv4addr = "10.20.1.3"
  view     = "default"
  use_ttl  = true
  ttl      = 10
  comment  = "Example A record"
  extattrs = {
    Site = "location-1"
  }
}

// Create Record A using function call to retrieve ipv4addr
resource "nios_dns_record_a" "create_record_a_with_func_call" {
  name = "example_func_call.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
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
}
