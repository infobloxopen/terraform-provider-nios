// Create a Shared Record Group (Required as Parent)
resource "nios_dns_sharedrecordgroup" "parent_sharedrecordgroup" {
  name = "example-sharedrecordgroup"
}

// Create a Shared SRV Record with Basic Fields
resource "nios_dns_sharedrecord_srv" "sharedrecord_srv_basic_fields" {
  name                = "sharedrecord_srv.example.com"
  port                = 443
  priority            = 10
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecordgroup.name
  target              = "server.example.com"
  weight              = 5
}

// Create a Shared SRV Record with Additional Fields
resource "nios_dns_sharedrecord_srv" "sharedrecord_srv_additional_fields" {
  name                = "_http._tcp.example.com"
  port                = 80
  priority            = 20
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecordgroup.name
  target              = "webserver.example.com"
  weight              = 10
  comment             = "Example Shared SRV Record"
  disable             = true
  extattrs = {
    Site = "location-1"
  }
  use_ttl = true
  ttl     = 7200
}
