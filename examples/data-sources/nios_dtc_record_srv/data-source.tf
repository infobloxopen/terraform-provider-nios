// Note: DTC SRV Record cannot be retrieved without filters. The dtc_server filter is mandatory for all queries.

// Retrieve a specific DTC SRV Record in a DTC server using filters 
data "nios_dtc_record_srv" "get_dtc_record_srv_in_dtc_server_using_filters" {
  filters = {
    dtc_server = "example-server"
    name       = "example_record._tcp.example.com"
  }
}


// Retrieve all DTC SRV Records in a DTC server
data "nios_dtc_record_srv" "get_all_dtc_record_srv_in_dtc_server" {
  filters = {
    dtc_server = "example-server"
  }
}
