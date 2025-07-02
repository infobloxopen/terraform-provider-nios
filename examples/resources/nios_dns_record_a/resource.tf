# resource "nios_dns_record_a" "create_record" {
#   count    = 1000
#   name     = "example_record${count.index}.example.com"
#   ipv4addr = "10.20.1.2"
#   view     = "default"
#   extattrs = {
#     Site = "Siteblr"
#   }
# }

# resource "nios_dns_record_cname" "create_record" {
#   count     = 1000
#   name      = "example_record${count.index}.example.com"
#   canonical = "canonical${count.index}.example.com"
#   view     = "default"
# }

resource "nios_dns_record_cname" "create_record" {
  count     = 1
  name      = "example_record3.example.com"
  canonical = "canonical3.example.com"
  view     = "default"
}
