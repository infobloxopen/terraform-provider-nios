// Create a Shared Record Group (Required as Parent)
resource "nios_dns_sharedrecordgroup" "parent_sharedrecord_group" {
  name = "example-sharedrecordgroup"
}

// Create a Shared A Record with Basic Fields
resource "nios_dns_sharedrecord_a" "shared_record_a_with_basic_fields" {
  name                = "sharedrecord_a_basic"
  ipv4addr            = "10.0.0.10"
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}

// Create a Shared A Record with Additional Fields
resource "nios_dns_sharedrecord_a" "shared_record_a_with_additional_fields" {
  name                = "sharedrecord_a_additional_fields"
  ipv4addr            = "20.0.0.0"
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name

  // Additional Fields
  extattrs = {
    Site = "location-1"
  }
  comment = "Example Sharedrecord A"
  disable = false
  ttl     = 7200
  use_ttl = true
}
