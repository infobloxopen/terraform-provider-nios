// Create Parent RP Zone
resource "nios_dns_zone_rp" "parent_zone" {
  fqdn = "rpz.example.com"
}

// Create Record RPZ NAPTR with Basic Fields
resource "nios_rpz_record_naptr" "create_record_rpz_naptr" {
  name        = "record1.${nios_dns_zone_rp.parent_zone.fqdn}"
  rp_zone     = nios_dns_zone_rp.parent_zone.fqdn
  order       = 10
  preference  = 10
  replacement = "."
}

// Create Record RPZ NAPTR with Additional Fields
resource "nios_rpz_record_naptr" "create_record_rpz_naptr_with_additional_fields" {
  name        = "record2.${nios_dns_zone_rp.parent_zone.fqdn}"
  rp_zone     = nios_dns_zone_rp.parent_zone.fqdn
  view        = "default"
  order       = 10
  preference  = 10
  replacement = "."
  flags       = "U"
  services    = "SIP+D2U"
  regexp      = "!^.*$!sip:jdoe@corpxyz.com!"
  use_ttl     = true
  ttl         = 10
  comment     = "Example RPZ NAPTR record"
  extattrs = {
    Site = "location-1"
  }
}
