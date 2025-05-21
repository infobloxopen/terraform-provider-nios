resource "nios_resource_nios_RecordA" "example_record_a" {
  name     = "example_recorda.example.com"
  ipv4addr = "10.20.1.2"

  # Other Optional fields
  view    = "default"
  comment = "Updated comment"
  use_ttl = true
  ttl     = 3600
  creator = "DYNAMIC"
  extattrs = {
    Site = {
      value = "us-west-1"
    }
  }
}
