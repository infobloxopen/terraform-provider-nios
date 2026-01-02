// Create a DTC Server (Required as parent)
resource "nios_dtc_server" "create_dtc_server" {
  name                    = "example-dtc-server"
  host                    = "2.3.3.4"
  auto_create_host_record = false
}

// Create a DTC CNAME Record with basic fields
resource "nios_dtc_record_cname" "record_cname_basic_fields" {
  dtc_server = nios_dtc_server.create_dtc_server.name
  canonical  = "example.com"
}

// Create a DTC Server (Required as parent)
resource "nios_dtc_server" "create_dtc_server2" {
  name                    = "example-dtc-server2"
  host                    = "2.3.3.4"
  auto_create_host_record = false
}
// Create a DTC CNAME Record with additional fields
resource "nios_dtc_record_cname" "record_cname_additional_fields" {
  dtc_server = nios_dtc_server.create_dtc_server2.name
  canonical  = "example-a.com"
  comment    = "Creating DTC CNAME Record"
  disable    = false
  ttl        = 20
  use_ttl    = true
}
