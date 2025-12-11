// Retrieve a specific DTC record SRV in a DTC server using filters 
data "nios_dtc_record_srv" "get_dtc_record_srv_in_dtc_server_using_filters" {
  filters = {
    dtc_server = "example-server2326763"
    name       = "_example._tcp.example.com"
  }
}


// Retrieve all DTC record SRV in a DTC server
data "nios_dtc_record_srv" "get_all_dtc_record_srv_in_dtc_server" {
  filters = {
    dtc_server = "example-server2326763"
  }
}
