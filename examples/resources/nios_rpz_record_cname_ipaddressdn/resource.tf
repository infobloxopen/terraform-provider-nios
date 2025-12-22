// Create Parent RP Zone
resource "nios_dns_zone_rp" "parent_zone" {
  fqdn = "rpz.example.com"
}

// Create Record RPZ CNAME IP Address DN with Basic Fields
resource "nios_rpz_record_cname_ipaddressdn" "create_record_rpz_cname_ipaddressdn" {
  name      = "10.10.0.1.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = "canonical1.${nios_dns_zone_rp.parent_zone.fqdn}"
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ CNAME IP Address DN with Additional Fields
resource "nios_rpz_record_cname_ipaddressdn" "create_record_rpz_cname_ipaddressdn_with_additional_fields" {
  name      = "10.10.0.0/16.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = "canonical2.${nios_dns_zone_rp.parent_zone.fqdn}"
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
  view      = "default"
  use_ttl   = true
  ttl       = 10
  comment   = "Example RPZ CNAME IP Address DN record"
  extattrs = {
    Site = "location-1"
  }
}

// Create Record RPZ CNAME IP Address DN with IPV6 Address
resource "nios_rpz_record_cname_ipaddressdn" "create_record_rpz_cname_ipaddressdn_with_additional_fields" {
  name      = "2001:db9::1.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = "canonical3.${nios_dns_zone_rp.parent_zone.fqdn}"
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ CNAME IP Address DN with IPV6 Network
resource "nios_rpz_record_cname_ipaddressdn" "create_record_rpz_cname_ipaddressdn_with_additional_fields" {
  name      = "2001:db8::/64.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = "canonical4.${nios_dns_zone_rp.parent_zone.fqdn}"
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
}
