// Create Parent RP Zone
resource "nios_dns_zone_rp" "parent_zone" {
  fqdn = "rpz.example.com"
}

// Create Record RPZ AAAA IPADDRESS with Basic Fields
resource "nios_rpz_record_aaaa_ipaddress" "create_record_rpz_aaaa_ipaddress" {
  name     = "2001:db8::/64.${nios_dns_zone_rp.parent_zone.fqdn}"
  ipv6addr = "2001:db8::1"
  rp_zone  = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ AAAA IPADDRESS with Additional Fields
resource "nios_rpz_record_aaaa_ipaddress" "create_record_rpz_aaaa_ipaddress_with_additional_fields" {
  name     = "2001:db9::1.${nios_dns_zone_rp.parent_zone.fqdn}"
  ipv6addr = "2001:db8::1"
  rp_zone  = nios_dns_zone_rp.parent_zone.fqdn
  view     = "default"
  use_ttl  = true
  ttl      = 10
  comment  = "Example RPZ AAAA IPADDRESS record"
  extattrs = {
    Site = "location-1"
  }
}
