// Create an Auth Zone
resource "nios_dns_zone_auth" "test" {
  fqdn = "example.com"
}

// Create a NAPTR Record with Basic Fields
resource "nios_dns_record_naptr" "record_naptr_with_basic_fields" {
  name        = nios_dns_zone_auth.test.fqdn
  order       = 10
  preference  = 10
  replacement = "."
  extattrs = {
    Site = "location-1"
  }
}

// Create a NAPTR Record with Additional Fields
resource "nios_dns_record_naptr" "record_naptr_with_additional_fields" {
  name        = nios_dns_zone_auth.test.fqdn
  order       = 10
  preference  = 10
  replacement = "."
  comment     = "NAPTR record created by Terraform"
  ttl         = 10
  use_ttl     = "true"
  flags       = "U"
  services    = "SIP+D2U"
  regexp      = "!^.*$!sip:jdoe@corpxyz.com!"
  extattrs = {
    Site = "location-1"
  }
}
