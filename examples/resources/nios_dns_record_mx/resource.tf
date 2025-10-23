// Create Authoritative zone (required as parent)
resource "nios_dns_zone_auth" "parent_auth_zone" {
  fqdn        = "example.com"
  zone_format = "FORWARD"
  view        = "default"
  comment     = "Parent zone for MX records"
}

// Create MX Record with Basic Fields
resource "nios_dns_record_mx" "record1" {
  name           = "mx_record.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  mail_exchanger = "mail.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  preference     = 10
  view           = "default"
  depends_on     = [nios_dns_zone_auth.parent_auth_zone]
}

// Create MX Record with additional fields
resource "nios_dns_record_mx" "record2" {
  name           = nios_dns_zone_auth.parent_auth_zone.fqdn
  mail_exchanger = "mail1.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  preference     = 20
  view           = "default"
  use_ttl        = true
  ttl            = 3600
  comment        = "Example MX Record"
  extattrs = {
    Site = "location-1"
  }
  depends_on = [nios_dns_zone_auth.parent_auth_zone]
}
