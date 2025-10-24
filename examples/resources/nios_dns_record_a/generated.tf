# __generated__ by Terraform
# Please review these resources and move them into your main configuration files.

# __generated__ by Terraform from "record:a/ZG5zLmJpbmRfYSQuMTMuY29tLnpvbmUsYV9yZWNvcmQsMi4yLjIuMg:a_record.zone.com/example_view"
resource "nios_dns_record_a" "import_rec" {
  comment               = "comment terr"
  creator               = "STATIC"
  ddns_principal        = null
  ddns_protected        = false
  disable               = false
  extattrs              = null
  forbid_reclamation    = false
  func_call             = null
  ipv4addr              = "2.2.2.2"
  name                  = "a_record.zone.com"
  remove_associated_ptr = null
  ttl                   = null
  use_ttl               = false
  view                  = "example_view"
}
