<<<<<<<< HEAD:examples/resources/dns_record_a/resource.tf
resource "nios_dns_record_a" "create_record" {
  name     = "example_record.example.com"
========
// Create Record A with Basic Fields
resource "nios_dns_record_a" "create_record" {
  name     = "example_record.example.com"
>>>>>>>> 4b5356041d52a8d5dbf9a93959ad537449c1f0a3:examples/resources/nios_dns_record_a/resource.tf
  ipv4addr = "10.20.1.2"
  view     = "default"
  extattrs = {
    Site = "Siteblr"
  }
}
<<<<<<<< HEAD:examples/resources/dns_record_a/resource.tf
resource "nios_dns_record_a" "create_with_func_call" {
  name = "example_test_func_call.example.com"
========

// Create Record A using function call to retrieve ipv4addr
resource "nios_dns_record_a" "create_with_func_call" {
  name = "example_func_call.example.com"
>>>>>>>> 4b5356041d52a8d5dbf9a93959ad537449c1f0a3:examples/resources/nios_dns_record_a/resource.tf
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
