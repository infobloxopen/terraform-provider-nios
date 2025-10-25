// Create an Auth zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_auth_zone" {
  fqdn        = "example.com"
  zone_format = "FORWARD"
  view        = "default"
  comment     = "Parent zone for CNAME records"
}

// Create Record CNAME with Basic Fields
resource "nios_dns_record_cname" "create_record_basic" {
  name      = "example_record.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  canonical = "example-canonical-name.${nios_dns_zone_auth.parent_auth_zone.fqdn}"

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
  depends_on = [nios_dns_zone_auth.parent_auth_zone]
}

// Record CNAME with Additional Fields
resource "nios_dns_record_cname" "create_record_additional_fields" {
  // Basic Fields
  name      = "example_record2.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  canonical = "example-canonical-name2.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  view      = "default"

  // Additional Fields
  ttl                = 3600
  use_ttl            = true
  creator            = "DYNAMIC"
  forbid_reclamation = false

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
  depends_on = [nios_dns_zone_auth.parent_auth_zone]
}
