// Create an Auth Zone (Required as Parent)
resource "nios_dns_zone_auth" "parent_zone" {
  fqdn = "example_auth_zone.com"
}

// Create an Shared Record Group with Basic Fields
resource "nios_dns_sharedrecordgroup" "shared_record_group_with_basic_fields" {
  name = "example-shared-record-group"
}

// Create an Shared Record Group with Additional Fields
resource "nios_dns_sharedrecordgroup" "shared_record_group_with_additional_fields" {
  name = "example-shared-record-group-2"

  // Additional Fields
  extattrs = {
    Site = "location-1"
  }
  record_name_policy     = "Allow Any"
  use_record_name_policy = true
  comment                = "Shared Record Group created by Terraform"
  zone_associations = [
    {
      fqdn = nios_dns_zone_auth.parent_zone.fqdn
      view = nios_dns_zone_auth.parent_zone.view
    }
  ]
}
