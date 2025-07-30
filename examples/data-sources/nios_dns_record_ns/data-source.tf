// Retrieve a specific NS record by filters
data "nios_dns_record_ns" "get_record_using_filters" {
  filters = {
    name       = "example.com"
    nameserver = "nsrec1.example.com"
  }
}

// Retrieve all NS records for a specific zone
data "nios_dns_record_ns" "get_all_records_in_default_view" {}
