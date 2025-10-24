// Create an Auth Zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_zone" {
  fqdn = "example.com"
}

// Create a CAA Record with Basic Fields
resource "nios_dns_record_caa" "record_caa_with_basic_fields" {
  name     = "caa_record.${nios_dns_zone_auth.parent_zone.fqdn}"
  ca_flag  = 1
  ca_tag   = "issue"
  ca_value = "digicert.com"
  extattrs = {
    Site = "location-1"
  }
}

// Create a CAA Record with Additional Fields
resource "nios_dns_record_caa" "record_caa_with_additional_fields" {
  name     = "caa_record1.${nios_dns_zone_auth.parent_zone.fqdn}"
  ca_flag  = 1
  ca_tag   = "issue"
  ca_value = "digicert.com"

  // Additional Fields
  extattrs = {
    Site = "location-1"
  }
  comment = "CAA Record created by Terraform"
  disable = false
  ttl     = 10
  use_ttl = true
}
