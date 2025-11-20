// Create a Shared Record Group (Required as Parent)
resource "nios_dns_sharedrecordgroup" "parent_sharedrecord_group" {
  name = "example-sharedrecordgroup"
}

// Create a Shared AAAA Record with Basic Fields
resource "nios_dns_sharedrecord_aaaa" "shared_record_aaaa_with_basic_fields" {
  name                = "sharedrecord_aaaa_basic"
  ipv6addr            = "2001:db8::1"
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}

// Create a Shared AAAA Record with Additional Fields
resource "nios_dns_sharedrecord_aaaa" "shared_record_aaaa_with_additional_fields" {
  name                = "sharedrecord_aaaa_additional_fields"
  ipv6addr            = "2001:db8::10"
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name

  // Additional Fields
  extattrs = {
    Site = "location-1"
  }

  comment = "Example Sharedrecord AAAA"
  disable = false
  ttl     = 7200
  use_ttl = true
}
