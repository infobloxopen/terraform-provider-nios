// Create an Auth Zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_auth_zone" {
  fqdn        = "example.com"
  zone_format = "FORWARD"
  view        = "default"
  comment     = "Parent zone for NS records"
}

// Create Record NS with Basic Fields
resource "nios_dns_record_ns" "record1" {
  name       = nios_dns_zone_auth.parent_auth_zone.fqdn
  nameserver = "nsrec1.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  addresses = [{
    address         = "192.168.1.10"
    auto_create_ptr = false
  }]
  view       = "default"
  depends_on = [nios_dns_zone_auth.parent_auth_zone]
}

// Create Record NS with PTR Record creation enabled
resource "nios_dns_record_ns" "record2" {
  name       = nios_dns_zone_auth.parent_auth_zone.fqdn
  nameserver = "nsrec2.${nios_dns_zone_auth.parent_auth_zone.fqdn}"
  addresses = [{
    address         = "192.168.1.11"
    auto_create_ptr = true
  }]
  view       = "default"
  depends_on = [nios_dns_zone_auth.parent_auth_zone]
}
