// Create an Auth Zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_zone" {
  fqdn = "example.com"
}

// Create a NAPTR Record with Basic Fields
resource "nios_dns_record_naptr" "record_naptr_with_basic_fields" {
  name        = "naptr_record.${nios_dns_zone_auth.parent_zone.fqdn}"
  order       = 10
  preference  = 10
  replacement = "."
  extattrs = {
    Site = "location-1"
  }
}

// Create a NAPTR Record with Additional Fields
resource "nios_dns_record_naptr" "record_naptr_with_additional_fields" {
  name        = "naptr_record1.${nios_dns_zone_auth.parent_zone.fqdn}"
  order       = 10
  preference  = 10
  replacement = "."
  comment     = "NAPTR record created by Terraform"
  ttl         = 10
  use_ttl     = true
  flags       = "U"
  services    = "SIP+D2U"
  regexp      = "!^.*$!sip:jdoe@corpxyz.com!"
  extattrs = {
    Site = "location-1"
  }
}
