// Create an Auth Zone
resource "nios_dns_zone_auth" "create_zone_auth" {
  fqdn = "example.com"
}

// Create a TLSA record with Basic Fields
resource "nios_dns_record_tlsa" "create_record_tlsa_with_basic_fields" {
  name              = "example-record-tlsa.${nios_dns_zone_auth.create_zone_auth.fqdn}"
  certificate_data  = "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971"
  certificate_usage = 2
  matched_type      = 0
  selector          = 0
}

// Create a TLSA record with Additional Fields
resource "nios_dns_record_tlsa" "create_record_tlsa_with_additional_fields" {
  name              = "example-record-tlsa-1.${nios_dns_zone_auth.create_zone_auth.fqdn}"
  certificate_data  = "D2ABDE240D7CD3EE6B4B28C54DF034B97983A1D16E8A410E4561CB106618E971"
  certificate_usage = 2
  matched_type      = 0
  selector          = 0

  extattrs = {
    Site = "location-1"
  }
  comment = "TLSA record created by Terraform"
}
