resource "nios_dns_view" "create_view" {
  name = "example_view"
}

resource "nios_dns_zone_auth" "create_zone1" {
  fqdn = "example.com"
  view = nios_dns_view.create_view.name
}

// Create a DNAME record with Basic Fields
resource "nios_dns_record_dname" "create_record_dname_with_basic_fields" {
  target = "example-dname-1.com"
  name   = nios_dns_zone_auth.create_zone1.fqdn
  view   = nios_dns_zone_auth.create_zone1.view
}

resource "nios_dns_zone_auth" "create_zone2" {
  fqdn = "example-1.com"
  view = nios_dns_view.create_view.name
}

// Create a DNAME record with Additional Fields
resource "nios_dns_record_dname" "create_record_dname_with_additional_fields" {
  target = "example-dname-2.com"
  name   = nios_dns_zone_auth.create_zone2.fqdn
  view   = nios_dns_zone_auth.create_zone2.view

  extattrs = {
    Site = "location-1"
  }
  comment = "DNAME record created by Terraform"
}

// Create IPV4 reverse mapping zone with Basic Fields
resource "nios_dns_zone_auth" "create_zone2" {
  fqdn        = "10.0.0.0/24"
  view        = nios_dns_view.create_view.name
  zone_format = "IPV4"
}

resource "nios_dns_record_dname" "create_record_dname1" {
  target = "example-dname-1.com"
  name   = nios_dns_zone_auth.create_zone2.display_domain
  view   = nios_dns_zone_auth.create_zone2.view
}

// Create IPV6 reverse mapping zone with Basic Fields
resource "nios_dns_zone_auth" "create_zone4" {
  fqdn        = "2002:1100::/64"
  view        = nios_dns_view.create_view.name
  zone_format = "IPV6"
}

resource "nios_dns_record_dname" "create_record_dname2" {
  target = "example-dname-1.com"
  name   = nios_dns_zone_auth.create_zone4.display_domain
  view   = nios_dns_zone_auth.create_zone4.view
}