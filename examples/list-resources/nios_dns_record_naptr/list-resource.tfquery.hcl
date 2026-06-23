// List specific NAPTR Records using filters
list "nios_dns_record_naptr" "list_naptr_records_using_filters" {
  provider = nios
  config {
    filters = {
      name = "record-naptr.example.com"
    }
  }
}

// List specific NAPTR Records using Extensible Attributes
list "nios_dns_record_naptr" "list_naptr_records_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List NAPTR Records with resource details included
list "nios_dns_record_naptr" "list_naptr_records_with_resource" {
  provider         = nios
  include_resource = true
}
