// Create an Auth Zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_auth_zone" {
  fqdn        = "example1.com"
  zone_format = "FORWARD"
  view        = "default"
  comment     = "Parent zone for SVCB records"
}

// Create Record SVCB with Basic Fields
resource "nios_dns_record_svcb" "record_svcb_basic_fields" {
  name        = "example-svcb-record.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  target_name = "example.target.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  priority    = 10
}

// Create Record SVCB with Additional Fields
resource "nios_dns_record_svcb" "record_svcb_additional_fields" {
  name        = "example-svcb-record-1.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  target_name = "example_updated.target.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  priority    = 20

  // Additional Fields
  creator            = "DYNAMIC"
  comment            = "Example SVCB record"
  ddns_principal     = "example_principal"
  ddns_protected     = true
  disable            = false
  forbid_reclamation = true
  svc_parameters = [
    {
      "mandatory" : false,
      "svc_key" : "alpn",
      "svc_value" : [
        "123",
        "145"
      ]
    },
    {
      "mandatory" : false,
      "svc_key" : "no_default_alpn",
      "svc_value" : [
        "True"
      ]
    },
    {
      "mandatory" : true,
      "svc_key" : "port",
      "svc_value" : [
        "110"
      ]
    },
    {
      "mandatory" : false,
      "svc_key" : "ipv4hint",
      "svc_value" : [
        "11.11.1.1"
      ]
    },
    {
      "mandatory" : false,
      "svc_key" : "ipv6hint",
      "svc_value" : [
        "123::99:0"
      ]
    }
  ]
  view    = "default"
  use_ttl = true
  ttl     = 10

  // Extensible Attributes
  extattrs = {
    Site = "location-2"
  }
}
