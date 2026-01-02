// Note: DTC NAPTR Record cannot be retrieved without filters. The dtc_server filter is mandatory for all queries.

// Retrieve a specific DTC NAPTR Record in a DTC server using filters 
data "nios_dtc_record_naptr" "get_dtc_record_naptr_in_dtc_server_using_filters" {
  filters = {
    dtc_server = "example-server"
    order      = 100
  }
}


// Retrieve all DTC NAPTR Records in a DTC server
data "nios_dtc_record_naptr" "get_all_dtc_record_naptr_in_dtc_server" {
  filters = {
    dtc_server = "example-server"
  }
}
