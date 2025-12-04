// Create Parent RP Zone
resource "nios_dns_zone_rp" "parent_zone" {
  fqdn = "rpz.example.com"
}

// Create Record RPZ SRV with Basic Fields
resource "nios_rpz_record_srv" "create_record_rpz_srv" {
  name     = "record1.${nios_dns_zone_rp.parent_zone.fqdn}"
  target   = "record1.target.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  rp_zone  = nios_dns_zone_rp.parent_zone.fqdn
  port     = 80
  priority = 10
  weight   = 20
}

// Create Record RPZ SRV with Additional Fields
resource "nios_rpz_record_srv" "create_record_rpz_srv_with_additional_fields" {
  name     = "record2.${nios_dns_zone_rp.parent_zone.fqdn}"
  target   = "record2.target.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  rp_zone  = nios_dns_zone_rp.parent_zone.fqdn
  port     = 443
  priority = 5
  weight   = 50
  view     = "default"
  use_ttl  = true
  ttl      = 10
  comment  = "Example RPZ SRV record"
  extattrs = {
    Site = "location-1"
  }
}
