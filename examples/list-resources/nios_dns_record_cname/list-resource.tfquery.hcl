// List specific CNAME Records using filters
list "nios_dns_record_cname" "list_cname_records_using_filters" {
  provider = nios
  config {
    filters = {
      name = "alias.example.com"
    }
  }
}

// List specific CNAME Records using Extensible Attributes
list "nios_dns_record_cname" "list_cname_records_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List CNAME Records with resource details included
list "nios_dns_record_cname" "list_cname_records_with_resource" {
  provider         = nios
  include_resource = true
}
