resource "nios_dns_record_aaaa" "record" {
  name     = "example_test.example.com"
  ipv6addr = "2002:1111::1401"
  view     = "default"
  extattrs = {
    Site = "Blr"
  }
}

data "nios_dns_record_aaaa" "get_record_with_filter" {
  filters = {
   "name" = nios_dns_record_aaaa.record.name
  }
}

data "nios_dns_record_aaaa" "get_record_with_extattr_filter" {
  extattrfilters = {
    "Site" = "Blr"
  }
}