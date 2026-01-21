// Retrieve a specific DTC A Record in a DTC server using filters 
data "nios_dtc_record_a" "get_dtc_record_a_in_dtc_server_using_filters" {
  filters = {
    dtc_server = "example-server"
    ipv4addr   = "2.2.2.2"
  }
}


// Retrieve all DTC A Records in a DTC server
data "nios_dtc_record_a" "get_all_dtc_record_a_in_dtc_server" {
  filters = {
    dtc_server = "example-server"
  }
}
