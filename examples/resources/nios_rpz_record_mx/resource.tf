// Create Parent RP Zone
resource "nios_dns_zone_rp" "parent_zone" {
  fqdn = "rpz.example.com"
}

// Create Record RPZ MX with Basic Fields
resource "nios_rpz_record_mx" "create_record_rpz_mx" {
  name           = "rpz_record_mx.${nios_dns_zone_rp.parent_zone.fqdn}"
  mail_exchanger = "mailexchanger1.example.com"
  preference     = 10
  rp_zone        = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ MX with Additional Fields
resource "nios_rpz_record_mx" "create_record_rpz_mx_with_additional_fields" {
  name           = "rpz_record_mx2.${nios_dns_zone_rp.parent_zone.fqdn}"
  mail_exchanger = "mailexchanger2.example.com"
  preference     = 20
  rp_zone        = nios_dns_zone_rp.parent_zone.fqdn
  view           = "default"
  use_ttl        = true
  ttl            = 3600
  comment        = "Example RPZ MX record"
  disable        = false
  extattrs = {
    Site = "location-1"
  }
}
