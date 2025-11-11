// Create an DTC Server(Required as Parent)
resource "nios_dtc_server" "create_dtc_server" {
  name = "example-server"
  host = "2.3.3.4"
}

// Create a DTC Topology with Basic Fields
resource "nios_dtc_topology" "create_dtc_topology" {
  name = "example_dtc_topology"
}

// Create a DTC Topology with Additional Fields
resource "nios_dtc_topology" "create_dtc_topology_with_additional_fields" {
  name    = "example_dtc_topology_2"
  comment = "DTC topology additional"
  rules = [
    {
      dest_type        = "SERVER"
      destination_link = nios_dtc_server.create_dtc_server.ref
    }
  ]
  extattrs = {
    Site = "location-1"
  }
}