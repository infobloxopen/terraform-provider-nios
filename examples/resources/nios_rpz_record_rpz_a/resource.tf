// Create Record RPZ A with Basic Fields
resource "nios_rpz_record_rpz_a" "create_record_rpz_a" {
  name     = "record1.rpz.example.com"
  ipv4addr = "10.10.0.1"
  rp_zone  = "rpz.example.com"
  extattrs = {
    Site = "location-1"
  }
}

// Create Record RPZ A with Additional Fields
resource "nios_rpz_record_rpz_a" "create_record_rpz_a_with_additional_fields" {
  name     = "record2.rpz.example.com"
  ipv4addr = "10.10.0.2"
  rp_zone  = "rpz.example.com"
  view     = "default"
  use_ttl  = true
  ttl      = 10
  comment  = "Example RPZ A record"
  extattrs = {
    Site = "location-1"
  }
}
