// Retrieve a specific DTC Topology by name
data "nios_dtc_topology" "get_topology_using_filters" {
  filters = {
    name = "example_dtc_topology"
  }
}

// Retrieve specific DTC Topology using Extensible Attributes
data "nios_dtc_topology" "get_topologies_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all DTC Topologies
data "nios_dtc_topology" "get_all_dtc_topologies" {
}
