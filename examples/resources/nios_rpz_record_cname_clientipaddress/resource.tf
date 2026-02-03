// Create Parent RP Zone
resource "nios_dns_zone_rp" "parent_zone" {
  fqdn = "rpzclientipaddress.example.com"
}

// Create Record RPZ CNAME client IP address with basic fields having passthru IP address rule.
resource "nios_rpz_record_cname_clientipaddress" "create_record_rpz_cname_clientipaddress" {
  name      = "12.0.0.0.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = "rpz-passthru"
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ CNAME client IP address with additional fields having passthru IP address rule.
resource "nios_rpz_record_cname_clientipaddress" "create_record_rpz_cname_clientipaddress_with_additional_fields" {
  name      = "12.0.0.1.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = "rpz-passthru"
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
  view      = "default"
  use_ttl   = true
  ttl       = 10
  comment   = "Example RPZ CNAME clientipaddress record"
  extattrs = {
    Site = "location-1"
  }
}

// Create Record RPZ CNAME client IP address - Block IP Address (No Such Domain) Rule.
resource "nios_rpz_record_cname_clientipaddress" "create_record_rpz_cname_clientipaddress_no_domain" {
  name      = "12.0.0.2.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = ""
  view      = "default"
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ CNAME client IP address - Block IP Address (No Data) Rule.
resource "nios_rpz_record_cname_clientipaddress" "create_record_rpz_cname_clientipaddress_no_data" {
  name      = "12.0.0.3.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = "*"
  view      = "default"
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
}
