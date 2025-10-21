// Create parent authoritative zone first (required as parent)
resource "nios_dns_zone_auth" "parent_auth_zone" {
  fqdn        = "example.com"
  zone_format = "FORWARD"
  view        = "default"
  comment     = "Parent zone for SRV and TXT records"
}

// Create Record SRV with Basic Fields
resource "nios_dns_record_srv" "create_record" {
  name     = "example-srv-record.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  target   = "example.target.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  port     = 80
  priority = 4
  weight   = 50

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
  depends_on = [nios_dns_zone_auth.parent_auth_zone]
}

// Create Record SRV with additional fields
resource "nios_dns_record_srv" "create_with_additional_config" {
  name     = "example-srv-record-with-config.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  target   = "example_updated.target.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  port     = 8080
  priority = 2
  weight   = 100

  // Additional Fields
  view    = "default"
  use_ttl = true
  ttl     = 10
  creator = "DYNAMIC"
  comment = "Example SRV record"

  // Extensible Attributes
  extattrs = {
    Site = "location-2"
  }
  depends_on = [nios_dns_zone_auth.parent_auth_zone]
}

// Create Record TXT with Basic Fields
resource "nios_dns_record_txt" "create_record" {
  name = "example-txt-record.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  text = "Example TXT Record"

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
  depends_on = [nios_dns_zone_auth.parent_auth_zone]
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
  depends_on = [nios_dns_zone_auth.parent_auth_zone]
}
