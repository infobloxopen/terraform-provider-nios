// Create an Auth zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_zone1" {
  fqdn = "example.com"
}

// Create a DNAME record with Basic Fields
resource "nios_dns_record_dname" "create_record_dname_with_basic_fields" {
  target = "example-dname-1.com"
  name   = nios_dns_zone_auth.parent_zone1.fqdn
}

// Create an Auth Zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_zone2" {
  fqdn = "example-1.com"
}

// Create a DNAME record with Additional Fields
resource "nios_dns_record_dname" "create_record_dname_with_additional_fields" {
  target = "example-dname-2.com"
  name   = nios_dns_zone_auth.parent_zone2.fqdn
  extattrs = {
    Site = "location-1"
  }
  comment = "DNAME record created by Terraform"
  ttl     = 10
  use_ttl = true
}

// Create an IPV4 reverse mapping zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_zone3" {
  fqdn        = "10.0.0.0/24"
  zone_format = "IPV4"
}

// Create a DNAME record in IPV4 reverse mapping zone
resource "nios_dns_record_dname" "create_record_dname1" {
  target = "example-dname-1.com"
  // We use display_domain for reverse mapping zones as arpa format is required for name
  name = nios_dns_zone_auth.parent_zone3.display_domain
}

// Create an IPV6 reverse mapping zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_zone4" {
  fqdn        = "2002:1100::/64"
  zone_format = "IPV6"
}

// Create a DNAME record in IPV6 reverse mapping zone
resource "nios_dns_record_dname" "create_record_dname2" {
  target = "example-dname-1.com"
  // We use display_domain for reverse mapping zones as arpa format is required for name
  name = nios_dns_zone_auth.parent_zone4.display_domain
}
