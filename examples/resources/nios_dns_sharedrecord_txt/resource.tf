// Create a Shared Record Group (Required as Parent)
resource "nios_dns_sharedrecordgroup" "parent_sharedrecord_group" {
  fqdn = "example-sharedrecordgroup"
}

// Create a Shared TXT Record with Basic Fields
resource "nios_dns_sharedrecord_txt" "shared_record_txt_with_basic_fields" {
  name                = "example-shared-record-txt"
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  text                = "Example TXT Shared Record"
}

// Create a Shared TXT Record with Additional Fields
resource "nios_dns_sharedrecord_txt" "shared_record_txt_with_additional_fields" {
  name                = "example-shared-record-txt2"
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  text                = "Example TXT Shared Record"

  // Additional Fields
  extattrs = {
    Site = "location-1"
  }
  comment = "Shared TXT Record created by Terraform"
  disable = false
  ttl     = 3600
  use_ttl = true
}
