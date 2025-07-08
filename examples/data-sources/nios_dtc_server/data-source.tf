// Retrieve a specific DTC Server by name
data "nios_dtc_server" "get_server_using_filters" {
  filters = {
    name = "server"
  }
}

// Retrieve specific DTC Server using Extensible Attributes
data "nios_dtc_server" "get_servers_using_extensible_attributes" {
  extattrfilters = {
    Site = "Bangalore"
  }
}

// Retrieve all DTC Servers
data "nios_dtc_server" "dtc_server_read_all" {
}
