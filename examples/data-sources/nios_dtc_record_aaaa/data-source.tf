// Note: DTC AAAA Record cannot be retrieved without filters. The dtc_server filter is mandatory for all queries.

// Retrieve a specific DTC record AAAA in a DTC server using filters 
data "nios_dtc_record_aaaa" "get_dtc_record_aaaa_in_dtc_server_using_filters" {
  filters = {
    dtc_server = "example-server"
    ipv6addr   = "2001:db8:85a3::8a2e:370:7335"
  }
}


// Retrieve all DTC AAAA records in a DTC server
data "nios_dtc_record_aaaa" "get_all_dtc_record_aaaa_in_dtc_server" {
  filters = {
    dtc_server = "example-server"
  }
}
