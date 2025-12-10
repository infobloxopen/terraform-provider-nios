// Retrieve a specific DTC record A in a DTC server using filters 
data "nios_dtc_record_a" "get_dtc_record_a_in_dtc_server_using_filters" {
  filters = {
    dtc_server = "example-server"
    ipv4addr   = "2.21.2.3"
  }
}


// Retrieve all DTC record A in a DTC server
data "nios_dtc_record_a" "get_all_dtc_record_a_in_dtc_server" {
  filters = {
    dtc_server = "example-server"
  }
}
