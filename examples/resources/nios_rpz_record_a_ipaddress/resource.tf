// Create Parent RP Zone
resource "nios_dns_zone_rp" "parent_zone" {
  fqdn = "rpz.example.com"
}

// Create Record RPZ A IPADDRESS with Basic Fields
resource "nios_rpz_record_a_ipaddress" "create_record_rpz_a_ipaddress" {
  name     = "10.10.0.1.${nios_dns_zone_rp.parent_zone.fqdn}"
  ipv4addr = "10.10.0.10"
  rp_zone  = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ A IPADDRESS with Additional Fields
resource "nios_rpz_record_a_ipaddress" "create_record_rpz_a_ipaddress_with_additional_fields" {
  name     = "10.10.0.0/16.${nios_dns_zone_rp.parent_zone.fqdn}"
  ipv4addr = "10.10.0.10"
  rp_zone  = nios_dns_zone_rp.parent_zone.fqdn
  view     = "default"
  use_ttl  = true
  ttl      = 10
  comment  = "Example RPZ A IPADDRESS record"
  extattrs = {
    Site = "location-1"
  }
}
