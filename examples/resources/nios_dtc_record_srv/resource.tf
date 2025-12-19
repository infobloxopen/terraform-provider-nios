// Create a DTC Server (Required as parent)
resource "nios_dtc_server" "create_dtc_server" {
  name = "example-server"
  host = "2.3.3.4"
}

// Create a DTC SRV Record with basic fields
resource "nios_dtc_record_srv" "srv_record_basic_fields" {
  dtc_server = nios_dtc_server.create_dtc_server.name
  port       = 8081
  priority   = 5
  target     = "service.basic.com"
  weight     = 10
}

// Create a DTC SRV Record with additional fields
resource "nios_dtc_record_srv" "srv_record_with_additional_fields" {
  dtc_server = nios_dtc_server.create_dtc_server.name
  port       = 8080
  priority   = 10
  target     = "service.example.com"
  weight     = 5
  comment    = "Example SRV record"
  disable    = false
  ttl        = 3600
  use_ttl    = true
  name       = "_example._tcp.example.com"
}
