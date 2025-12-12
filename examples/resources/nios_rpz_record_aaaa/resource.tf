// Create Parent RP Zone
resource "nios_dns_zone_rp" "parent_zone" {
  fqdn = "rpz.example.com"
}

// Create Record RPZ AAAA with Basic Fields
resource "nios_rpz_record_aaaa" "create_record_rpz_aaaa" {
  name     = "record1.${nios_dns_zone_rp.parent_zone.fqdn}"
  ipv6addr = "2002:1f93::12:1"
  rp_zone  = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ AAAA with Additional Fields
resource "nios_rpz_record_aaaa" "create_record_rpz_aaaa_with_additional_fields" {
  name     = "record2.${nios_dns_zone_rp.parent_zone.fqdn}"
  ipv6addr = "2002:1f93::12:10"
  rp_zone  = nios_dns_zone_rp.parent_zone.fqdn
  view     = "default"
  use_ttl  = true
  ttl      = 10
  comment  = "Example RPZ AAAA record"
  extattrs = {
    Site = "location-1"
  }
}
