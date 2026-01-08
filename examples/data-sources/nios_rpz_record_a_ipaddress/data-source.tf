// Retrieve a specific RPZ A IPADDRESS record by filters
data "nios_rpz_record_a_ipaddress" "get_record_using_filters" {
  filters = {
    name = "10.10.0.1.rpz.example.com"
  }
}

// Retrieve specific RPZ A IPADDRESS records using Extensible Attributes
data "nios_rpz_record_a_ipaddress" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ A IPADDRESS records
data "nios_rpz_record_a_ipaddress" "get_all_rpz_a_ipaddress_records" {}
