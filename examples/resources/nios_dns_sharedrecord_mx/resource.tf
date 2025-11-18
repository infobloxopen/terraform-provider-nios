// Create Parent Shared Record Group
resource "nios_dns_sharedrecordgroup" "parent_sharedrecordgroup" {
  name = "example-sharedrecordgroup"
}

// Create a Shared MX Record with Basic Fields
resource "nios_dns_sharedrecord_mx" "sharedrecord_mx_basic_fields" {
  mail_exchanger      = "mail.example.com"
  name                = "sharedrecord_mx_basic"
  preference          = 10
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecordgroup.name
}

// Create a Shared MX Record with Additional Fields
resource "nios_dns_sharedrecord_mx" "sharedrecord_mx_additional_fields" {
  mail_exchanger      = "mail.example.com"
  name                = "sharedrecord_mx_additional_fields"
  preference          = 20
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecordgroup.name
  comment             = "Example MX Shared Record"
  disable             = true
  extattrs = {
    Site = "location-1"
  }
  use_ttl = true
  ttl     = 7200
}
