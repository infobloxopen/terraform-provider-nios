// Create an IPv4 PTR record with Basic fields
#resource "nios_dns_record_ptr" "create_ptr_record_with_ipv4addr" {
#  ptrdname = "example_record.example.com"
#  ipv4addr = "10.20.1.2"
#  name = ""
#  view     = "default"
#  extattrs = {
#    Site = "location-1"
#  }
#}

// Create an IPv6 PTR record with Basic fields
#resource "nios_dns_record_ptr" "create_ptr_record_with_ipv6addr" {
#  ptrdname = "example_record.example.com"
#  ipv6addr = "2001::123"
#  view     = "default"
#  extattrs = {
#    Site = "location-2"
#  }
#}

// Create an IPv4 PTR record by name with Basic fields
resource "nios_dns_record_ptr" "ptr123" {
#  name = "example_record.example.com"
  ptrdname     = "0.192.in-addr"
  ipv4addr = "192.0.1.22"
  view     = "default"
  extattrs = {
    Site = "location-3"
  }
}

// Create an IPv4 PTR record by name with Additional fields
#resource "nios_dns_record_ptr" "create_ptr_record_with_additional_fields" {
#  ptrdname = "example_record.example.com"
#  name     = "22.0.0.11.in-addr.arpa"
#
#  // Additional Fields
#  view    = "default"
#  use_ttl = true
#  ttl     = 10
#  creator = "DYNAMIC"
#  comment = "Example PTR record"
#
#  // Extensible Attributes
#  extattrs = {
#    Site = "location-4"
#  }
#}

// Create an PTR record using function call to retrieve ipv4addr
#resource "nios_dns_record_ptr" "create_ptr_record_with_func_call" {
#  ptrdname = "example_func_call.example.com"
#  func_call = {
#    attribute_name  = "ipv4addr"
#    object_function = "next_available_ip"
#    result_field    = "ips"
#    object          = "network"
#    object_parameters = {
#      network      = "85.85.0.0/16"
#      network_view = "default"
#    }
#  }
#  view    = "default"
#  comment = "Updated comment"
#}
