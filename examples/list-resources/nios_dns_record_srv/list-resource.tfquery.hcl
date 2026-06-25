// List specific SRV Records using filters
list "nios_dns_record_srv" "list_srv_records_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example-srv-record.example.com"
    }
  }
}

// List specific SRV Records using Extensible Attributes
list "nios_dns_record_srv" "list_srv_records_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List SRV Records with resource details included
list "nios_dns_record_srv" "list_srv_records_with_resource" {
  provider         = nios
  include_resource = true
}
