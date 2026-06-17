// List specific DNAME Records using filters
list "nios_dns_record_dname" "list_dname_records_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example.com"
    }
  }
}

// List specific DNAME Records using Extensible Attributes
list "nios_dns_record_dname" "list_dname_records_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List DNAME Records with resource details included
list "nios_dns_record_dname" "list_dname_records_with_resource" {
  provider         = nios
  include_resource = true
}
