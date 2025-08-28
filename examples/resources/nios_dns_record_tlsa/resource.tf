// Create an Auth Zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_zone" {
  fqdn = "example.com"
}

// Create a TLSA record with Basic Fields
resource "nios_dns_record_tlsa" "create_record_tlsa_with_basic_fields" {
  name              = "record-tlsa.${nios_dns_zone_auth.parent_zone.fqdn}"
  certificate_data  = "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971"
  certificate_usage = 2
  matched_type      = 10
  selector          = 20
}

// Create a TLSA record with Additional Fields
resource "nios_dns_record_tlsa" "create_record_tlsa_with_additional_fields" {
  name              = "record-tlsa-1.${nios_dns_zone_auth.parent_zone.fqdn}"
  certificate_data  = "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971"
  certificate_usage = 2
  matched_type      = 10
  selector          = 15

  // Additional Fields
  extattrs = {
    Site = "location-1"
  }
  ttl     = 3600
  use_ttl = true
  disable = false
  creator = "STATIC"
  comment = "TLSA record created by Terraform"
}
