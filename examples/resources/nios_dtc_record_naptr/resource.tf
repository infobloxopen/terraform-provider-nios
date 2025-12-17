// Create a DTC Server (Required as parent)
resource "nios_dtc_server" "create_dtc_server" {
  name = "example-server"
  host = "2.3.3.4"
}

// Create a DTC NAPTR Record with basic fields
resource "nios_dtc_record_naptr" "record_naptr_basic_fields" {
  dtc_server  = nios_dtc_server.create_dtc_server.name
  order       = 100
  preference  = 10
  replacement = "naptr.example.com"
}

// Create a DTC NAPTR Record with additional fields
resource "nios_dtc_record_naptr" "record_naptr_additional_fields" {
  dtc_server  = nios_dtc_server.create_dtc_server.name
  comment     = "This is a sample NAPTR record"
  order       = 200
  preference  = 20
  flags       = "S"
  services    = "SIP+D2U"
  regexp      = "!^.*$!sip:example@domain.com!"
  replacement = "naptr.example.com"
  ttl         = 3600
  use_ttl     = true
  disable     = false
}