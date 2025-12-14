// Create a DTC Server (Required as parent)
resource "nios_dtc_server" "create_dtc_server" {
  name = "example-server"
  host = "2.3.3.4"
}

// Create a DTC A Record with basic fields
resource "nios_dtc_record_a" "record_a_basic_fields" {
  dtc_server = nios_dtc_server.create_dtc_server.name
  ipv4addr   = "2.2.2.2"
}

// Create a DTC A Record with additional fields
resource "nios_dtc_record_a" "record_a_additional_fields" {
  dtc_server = nios_dtc_server.create_dtc_server.name
  ipv4addr   = "2.2.2.3"
  comment    = "Creating DTC A Record"
  disable    = false
  ttl        = 20
  use_ttl    = true
}
