// Create Record CNAME with Basic Fields
resource "nios_dns_record_cname" "create_record_basic" {
  name      = "example_record.example.com"
  canonical = "example-canonical-name"

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}

// Record CNAME with Additional Fields
resource "nios_dns_record_cname" "create_record_additional_fields" {
  // Basic Fields
  name      = "example_record2.example.com"
  canonical = "example-canonical-name2"
  view      = "default"

  // Additional Fields
  ttl                = 3600
  use_ttl            = true
  creator            = "DYNAMIC"
  forbid_reclamation = false

  // Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}
