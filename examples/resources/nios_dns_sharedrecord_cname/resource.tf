// Create a Shared Record Group (Required as Parent)
resource "nios_dns_sharedrecordgroup" "parent_sharedrecord_group" {
  name = "example-sharedrecordgroup"
}

// Create a Shared CNAME Record with Basic Fields
resource "nios_dns_sharedrecord_cname" "shared_cname_record_with_basic_fields" {
  name                = "example-shared-record-cname"
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  canonical           = "example.canonical.com"
}

// Create a Shared CNAME Record with Additional Fields
resource "nios_dns_sharedrecord_cname" "shared_cname_record_with_additional_fields" {
  name                = "example-shared-record-cname2"
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  canonical           = "example.canonical2.com"

  // Additional Fields
  extattrs = {
    Site = "location-1"
  }
  comment = "Shared CNAME Record created by Terraform"
  disable = false
  ttl     = 3600
  use_ttl = true
}
