// Create Parent RP Zone
resource "nios_dns_zone_rp" "parent_zone" {
  fqdn = "rpz.example.com"
}

// Create Record RPZ A with Basic Fields
resource "nios_rpz_record_a" "create_record_rpz_a" {
  name     = "record1.${nios_dns_zone_rp.parent_zone.fqdn}"
  ipv4addr = "10.10.0.1"
  rp_zone  = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ A with Additional Fields
resource "nios_rpz_record_a" "create_record_rpz_a_with_additional_fields" {
  name     = "record2.${nios_dns_zone_rp.parent_zone.fqdn}"
  ipv4addr = "10.10.0.2"
  rp_zone  = nios_dns_zone_rp.parent_zone.fqdn
  view     = "default"
  use_ttl  = true
  ttl      = 10
  comment  = "Example RPZ A record"
  extattrs = {
    Site = "location-1"
  }
}
