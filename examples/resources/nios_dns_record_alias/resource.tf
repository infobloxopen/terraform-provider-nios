// Create Record Alias with Basic Fields
resource "nios_dns_record_alias" "create_alias_record" {
  name        = "alias-record.example.com"
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
    Site = "location-1"
  }
  ttl     = 20
  use_ttl = true
}

import {
  id = "record:alias/ZG5zLmFsaWFzX3JlY29yZCQuX2RlZmF1bHQuY29tLmluZm8uaW1wb3J0LWFsaWFzLk5BUFRS:import-alias.info.com/default"
  to = nios_dns_record_alias.import4321
}

resource "nios_dns_record_alias" "import2" {


}