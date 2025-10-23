// Create Authoritative zone (required as parent)
resource "nios_dns_zone_auth" "parent_auth_zone" {
  fqdn        = "example.com"
  zone_format = "FORWARD"
  view        = "default"
  comment     = "Parent zone for SRV and TXT records"
}

// Create Record TXT with Basic Fields
resource "nios_dns_record_txt" "create_record" {
  name = "example-txt-record.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  text = "Example TXT Record"

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}

// Create Record TXT with additional fields
resource "nios_dns_record_txt" "create_with_additional_config" {
  name = "example-txt-record-with-config.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  text = "Example TXT Record with Additional Config"

  // Additional Fields
  view    = "default"
  use_ttl = true
  ttl     = 10
  creator = "DYNAMIC"
  comment = "Example TXT record"

  // Extensible Attributes
  extattrs = {
    Site = "location-2"
  }
}
