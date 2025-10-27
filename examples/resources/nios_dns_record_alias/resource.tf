// Create an Auth Zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_auth_zone" {
  fqdn        = "example.com"
  zone_format = "FORWARD"
  view        = "default"
  comment     = "Parent zone for Alias records"
}

// Create Record Alias with Basic Fields
resource "nios_dns_record_alias" "create_alias_record" {
  name        = "alias-record.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  target_name = "server.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  target_type = "A"
  view        = "default"
  depends_on  = [nios_dns_zone_auth.parent_auth_zone]
}

// Create Record Alias with Additional Fields
resource "nios_dns_record_alias" "create_alias_record_with_additional_fields" {
  name        = "alias-record2.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  target_name = "webserver.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  target_type = "A"
  view        = "default"

  // Optional fields
  comment = "Alias record with additional parameters"
  disable = false
  extattrs = {
    Site = "location-1"
  }
  ttl        = 20
  use_ttl    = true
  depends_on = [nios_dns_zone_auth.parent_auth_zone]
}
