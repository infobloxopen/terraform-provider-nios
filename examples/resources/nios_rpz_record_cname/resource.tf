// Create Parent RP Zone
resource "nios_dns_zone_rp" "parent_zone" {
  fqdn = "rpz.example.com"
}

// Create Record RPZ CNAME with Basic Fields
resource "nios_rpz_record_cname" "create_record_rpz_cname" {
  name      = "record1.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = "canonical1.${nios_dns_zone_rp.parent_zone.fqdn}"
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ CNAME with Additional Fields
resource "nios_rpz_record_cname" "create_record_rpz_cname_with_additional_fields" {
  name      = "record2.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = "canonical2.${nios_dns_zone_rp.parent_zone.fqdn}"
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
  view      = "default"
  use_ttl   = true
  ttl       = 10
  comment   = "Example RPZ CNAME record"
  extattrs = {
    Site = "location-1"
  }
}

// Create Record RPZ CNAME - Block Domain (No Such Domain Rule)
resource "nios_rpz_record_cname" "create_record_rpz_cname_no_domain" {
  name      = "record3.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = ""
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ CNAME - Block Domain (No Data Rule)
resource "nios_rpz_record_cname" "create_record_rpz_cname_no_data" {
  name      = "record4.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = "*"
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ CNAME - Passthru Domain Name Rule
resource "nios_rpz_record_cname" "create_record_rpz_cname_passthru" {
  name      = "record5.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = "record5"
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ CNAME - Wildcard Passthru Domain Name Rule
// For wildcard names, canonical must be 'infoblox-passthru'
resource "nios_rpz_record_cname" "create_record_rpz_cname_wildcard" {
  name      = "*.record6.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = "infoblox-passthr"
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
}
