// Create Parent RP Zone
resource "nios_dns_zone_rp" "parent_zone" {
  fqdn = "rpz.example.com"
}

// Create Record RPZ PTR with Basic Fields
resource "nios_rpz_record_ptr" "create_record_rpz_ptr" {
  ptrdname = "record1.${nios_dns_zone_rp.parent_zone.fqdn}"
  ipv4addr = "10.10.0.1"
  rp_zone  = nios_dns_zone_rp.parent_zone.fqdn
}

// Create Record RPZ PTR with Additional Fields
resource "nios_rpz_record_ptr" "create_record_rpz_ptr_with_additional_fields" {
  ptrdname = "record2.${nios_dns_zone_rp.parent_zone.fqdn}"
  ipv4addr = "10.10.0.2"
  rp_zone  = nios_dns_zone_rp.parent_zone.fqdn
  view     = "default"
  use_ttl  = true
  ttl      = 10
  comment  = "Example RPZ PTR record"
  extattrs = {
    Site = "location-1"
  }
}

// Create Record RPZ PTR with Name
resource "nios_rpz_record_ptr" "create_record_rpz_ptr_with_name" {
  ptrdname = "record3.${nios_dns_zone_rp.parent_zone.fqdn}"
  name     = "3.0.10.10.in-addr.arpa.${nios_dns_zone_rp.parent_zone.fqdn}"
  rp_zone  = nios_dns_zone_rp.parent_zone.fqdn
  view     = "default"
  extattrs = {
    Site = "location-1"
  }
}

// Create Record RPZ PTR with IPv6 Address
resource "nios_rpz_record_ptr" "create_record_rpz_ptr_with_ipv6addr" {
  ptrdname = "record4.${nios_dns_zone_rp.parent_zone.fqdn}"
  ipv6addr = "2002:1f93::12:1"
  rp_zone  = nios_dns_zone_rp.parent_zone.fqdn
  view     = "default"
  extattrs = {
    Site = "location-1"
  }
}
