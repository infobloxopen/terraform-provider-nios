// Note: DTC CNAME Record cannot be retrieved without filters. The dtc_server filter is mandatory for all queries.

// Retrieve a specific DTC CNAME Record in a DTC server using filters 
data "nios_dtc_record_cname" "get_dtc_cname_record_in_dtc_server_using_filters" {
  filters = {
    dtc_server = "example-dtc-server"
    canonical  = "example_canonical_name.com"
  }
}


// Retrieve all DTC CNAME Records in a DTC server
data "nios_dtc_record_cname" "get_all_dtc_cname_records_in_dtc_server" {
  filters = {
    dtc_server = "example-dtc-server"
  }
}
