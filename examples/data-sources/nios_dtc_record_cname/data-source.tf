// Retrieve a specific DTC record CNAME in a DTC server using filters 
data "nios_dtc_record_cname" "get_dtc_record_cname_in_dtc_server_using_filters" {
  filters = {
    dtc_server = "example-dtc-server"
    canonical  = "example.com"
  }
}


// Retrieve all DTC record CNAME in a DTC server
data "nios_dtc_record_cname" "get_all_dtc_record_cname_in_dtc_server" {
  filters = {
    dtc_server = "example-dtc-server"
  }
}
