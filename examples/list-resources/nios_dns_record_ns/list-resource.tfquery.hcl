// List specific NS Records using filters
list "nios_dns_record_ns" "list_ns_records_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example.com"
    }
  }
}

// List NS Records with resource details included
list "nios_dns_record_ns" "list_ns_records_with_resource" {
  provider         = nios
  include_resource = true
}
