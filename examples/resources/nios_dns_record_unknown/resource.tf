// Create an Auth zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_zone" {
  fqdn = "example.com"
}

// Create an Unknown record with Basic Fields (Type: SPF)
resource "nios_dns_record_unknown" "record_spf_with_basic_fields" {
  name        = "record-spf.${nios_dns_zone_auth.parent_zone.fqdn}"
  record_type = "SPF"
  subfield_values = [
    {
      field_type     = "T"
      field_value    = "example-text",
      include_length = "8_BIT"
    }
  ]
}

// Create an Unknown record with Basic Fields (Type: RP)
resource "nios_dns_record_unknown" "record_rp_with_basic_fields" {
  name        = "record-rp.${nios_dns_zone_auth.parent_zone.fqdn}"
  record_type = "RP"
  subfield_values = [
    {
      field_type     = "N"
      field_value    = "example1.com",
      include_length = "NONE"
    },
    {
      field_type     = "N",
      field_value    = "example2.com",
      include_length = "NONE"
    }
  ]
}

// Create an Unknown record with Basic Fields (Type: HINFO)
resource "nios_dns_record_unknown" "record_hinfo_with_basic_fields" {
  name        = "record-hinfo.${nios_dns_zone_auth.parent_zone.fqdn}"
  record_type = "HINFO"
  subfield_values = [
    {
      field_type     = "T"
      field_value    = "INTEL-386",
      include_length = "8_BIT"
    },
    {
      field_type     = "T",
      field_value    = "WIN32",
      include_length = "8_BIT"
    }
  ]
}

// Create an Unknown record with Additional Fields (Type: SPF)
resource "nios_dns_record_unknown" "record_unknown_with_additional_fields" {
  name        = "record-spf.${nios_dns_zone_auth.parent_zone.fqdn}"
  record_type = "SPF"
  subfield_values = [
    {
      field_type     = "T"
      field_value    = "example-text-1",
      include_length = "8_BIT"
    }
  ]

  // Additional Fields
  extattrs = {
    Site = "location-1"
  }
  ttl     = 3600
  use_ttl = true
  disable = false
  creator = "STATIC"
  comment = "Unknown record created by Terraform"
}
