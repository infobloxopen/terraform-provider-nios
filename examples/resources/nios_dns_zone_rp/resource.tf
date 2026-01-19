// Create IPV4 forward mapping zone with Basic Fields
resource "nios_dns_zone_rp" "zone_rp_basic_fields" {
  fqdn = "example1.com"
  view = "default"
  extattrs = {
    Site = "location-1"
  }
}

// Create IPV4 forward mapping zone with Additional Fields
resource "nios_dns_zone_rp" "zone_rp_additional_fields" {
  // Basic Fields
  fqdn = "example2.com"
  view = "default"

  // Additional Fields
  grid_primary = [
    {
      name = "infoblox.member",
    }
  ]

  soa_default_ttl     = 37000
  soa_expire          = 92000
  soa_negative_ttl    = 900
  soa_refresh         = 2100
  soa_retry           = 800
  use_grid_zone_timer = true


  comment = "Comment for Zone RP"
  extattrs = {
    Site = "location-2"
  }
}
