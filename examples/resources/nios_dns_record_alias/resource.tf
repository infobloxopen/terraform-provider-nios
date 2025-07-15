// Create Record Alias with Basic Required Fields
resource "nios_dns_record_alias" "create_alias_record" {
  name        = "alias-record1.example.com"
  target_name = "server.example.com"
  target_type = "A"
  view        = "default"
}

// Create Record Alias with Additional Fields
resource "nios_dns_record_alias" "create_alias_record_with_additional_fields" {
  name        = "alias-record2.example.com"
  target_name = "webserver.example.com"
  target_type = "A"
  view        = "default"

  // Optional fields
  comment = "Alias record with additional parameters"
  disable = false
  extattrs = {
    Site = "site-1"
  }
  ttl     = 20
  use_ttl = true
}
