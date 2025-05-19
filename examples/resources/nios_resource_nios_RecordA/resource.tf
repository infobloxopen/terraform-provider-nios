resource "nios_resource_nios_RecordA" "create_record" {
  name     = "example_test.example.com"
  ipv4addr = "10.20.1.2"
  view     = "default"
  extattrs = {
    Site = {
      value = "Siteblr"
    }
  }
}
