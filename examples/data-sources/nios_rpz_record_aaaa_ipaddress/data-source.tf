// Retrieve a specific RPZ AAAA IPADDRESS record by filters
data "nios_rpz_record_aaaa_ipaddress" "get_record_using_filters" {
  filters = {
    name = "2001:db8::/64.rpz.example.com"
  }
}

// Retrieve specific RPZ AAAA IPADDRESS records using Extensible Attributes
data "nios_rpz_record_aaaa_ipaddress" "get_record_using_extensible_attributes" {
  extattrfilters = {
    Site = "location-1"
  }
}

// Retrieve all RPZ AAAA IPADDRESS records
data "nios_rpz_record_aaaa_ipaddress" "get_all_rpz_aaaa_ipaddress_records" {}
