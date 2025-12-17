// Create Parent RP Zone
resource "nios_dns_zone_rp" "parent_zone" {
  fqdn = "rpzip.example.com"
}

// Create an IPV4 network if not present (Required as Parent)
# resource "nios_ipam_network" "range_parent_network" {
#   network      = "11.0.0.0/8"
#   network_view = "default"
#   comment      = "Parent network for DHCP ranges"
# }

// Create Record RPZ CNAME ipaddress with Basic Fields
resource "nios_rpz_record_rpz_cname_ipaddress" "create_record_rpz_cname_ipaddress" {
  name      = "11.0.0.0.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = "11.0.0.0"
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ CNAME ipaddress with Additional Fields
resource "nios_rpz_record_rpz_cname_ipaddress" "create_record_rpz_cname_ipaddress_with_additional_fields" {
  name      = "11.0.0.1.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = "11.0.0.1"
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
  view      = "default"
  use_ttl   = true
  ttl       = 10
  comment   = "Example RPZ CNAME ipaddress record"
  extattrs = {
    Site = "location-1"
  }
}

// Create Record RPZ CNAME ipaddress - Block Domain (No Such Domain Rule)
resource "nios_rpz_record_rpz_cname_ipaddress" "create_record_rpz_cname_no_domain" {
  name      = "11.0.0.2.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = ""
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ CNAME ipaddress - Block Domain (No Data Rule)
resource "nios_rpz_record_rpz_cname_ipaddress" "create_record_rpz_cname_no_data" {
  name      = "11.0.0.3.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = "*"
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ CNAME ipaddress - Passthru Domain Name Rule
resource "nios_rpz_record_rpz_cname_ipaddress" "create_record_rpz_cname_passthru" {
  name      = "11.0.0.4.${nios_dns_zone_rp.parent_zone.fqdn}"
  canonical = "11.0.0.4"
  rp_zone   = nios_dns_zone_rp.parent_zone.fqdn
}
