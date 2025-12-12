// Create Parent RP Zone
resource "nios_dns_zone_rp" "parent_zone" {
  fqdn = "rpz.example.com"
}

// Create Record RPZ TXT with Basic Fields
resource "nios_rpz_record_txt" "create_record_rpz_txt" {
  name    = "record1.${nios_dns_zone_rp.parent_zone.fqdn}"
  text    = "Example text"
  rp_zone = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ TXT with Additional Fields
resource "nios_rpz_record_txt" "create_record_rpz_txt_with_additional_fields" {
  name    = "record2.${nios_dns_zone_rp.parent_zone.fqdn}"
  text    = "Example text with Additional Config"
  rp_zone = nios_dns_zone_rp.parent_zone.fqdn
  view    = "default"
  use_ttl = true
  ttl     = 10
  comment = "Example RPZ TXT record"
  extattrs = {
    Site = "location-1"
  }
}
