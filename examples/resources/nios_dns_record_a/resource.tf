# resource "nios_dns_record_a" "create_record" {
#   count    = 1000
#   name     = "example_test_${count.index}.example1"
#   ipv4addr = "10.20.1.2"
#   view     = "default"
#   extattrs = {
#     Site = "Siteblr"
#   }
# }

resource "nios_dns_record_cname" "create_record" {
  count     = 1000
  name      = "example_test_${count.index}.example1"
  canonical = "canonical2.example1.com"
  view     = "default"
  comment = "apply"
}

# resource "nios_dns_record_cname" "create_record" {
#   count     = 1
#   name      = "example_record.example.com"
#   canonical = "canonical3.example.com"
#   view     = "default"
# }
