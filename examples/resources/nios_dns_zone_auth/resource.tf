// Create IPV4 forward mapping zone with Basic Fields
resource "nios_dns_zone_auth" "create_zone1" {
  fqdn = "example1.com"
  view = "default"
  extattrs = {
    Site = "location-1"
  }
}

// Create IPV4 reverse mapping zone with Basic Fields
resource "nios_dns_zone_auth" "create_zone2" {
  fqdn        = "10.0.0.0/24"
  view        = "default"
  zone_format = "IPV4"
  extattrs = {
    Site = "location-1"
  }
}

// Create IPV6 reverse mapping zone with Basic Fields
resource "nios_dns_zone_auth" "create_zone4" {
  fqdn        = "2002:1100::/64"
  view        = "default"
  zone_format = "IPV6"
  extattrs = {
    Site = "location-1"
  }
}

// Create IPV4 forward mapping zone with Additional Fields
resource "nios_dns_zone_auth" "create_zone5" {
  // Basic Fields
  fqdn = "example2.com"
  view = "default"

  // Additional Fields
  grid_primary = [
    {
      name = "infoblox.10_0_0_1",
    }
  ]
  restart_if_needed = true

  soa_default_ttl     = 37000
  soa_expire          = 92000
  soa_negative_ttl    = 900
  soa_refresh         = 2100
  soa_retry           = 800
  use_grid_zone_timer = true

  allow_query = [
    {
      struct     = "addressac"
      address    = "10.0.0.0"
      permission = "ALLOW"
    }
  ]
  use_allow_query = true

  comment = "IPV4 forward auth zone"
  extattrs = {
    Site = "location-1"
  }
}
