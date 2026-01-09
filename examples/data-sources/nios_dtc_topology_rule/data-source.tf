// Create an DTC Server (Required as Parent)
resource "nios_dtc_server" "create_dtc_server" {
  name = "example-server"
  host = "2.3.3.4"
}

// Create DTC Topology (Required as Parent) 
resource "nios_dtc_topology" "create_dtc_topology" {
  name = "example_dtc_topology_2"
  rules = [
    {
      dest_type        = "SERVER"
      destination_link = nios_dtc_server.create_dtc_server.ref
    },
    {
      dest_type        = "SERVER"
      destination_link = nios_dtc_server.create_dtc_server.ref
      sources = [
        {
          source_op    = "IS",
          source_type  = "CONTINENT",
          source_value = "Africa"
        }
      ]
    }
  ]
  extattrs = {
    Site = "location-1"
  }
}

// Retrieve all DTC Topology rule in a DTC Topology 
data "nios_dtc_topology_rule" "get_topology_rules_using_filters" {
  filters = {
    topology = nios_dtc_topology.create_dtc_topology.ref
  }
}

// Retrieve all DTC Topology rule 
data "nios_dtc_topology_rule" "get_all_topology_rules" {
}
