// List specific TXT Records using filters
list "nios_dns_record_txt" "list_txt_records_using_filters" {
  provider = nios
  config {
    filters = {
      name = "record-txt.example.com"
    }
  }
}

// List specific TXT Records using Extensible Attributes
list "nios_dns_record_txt" "list_txt_records_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List TXT Records with resource details included
list "nios_dns_record_txt" "list_txt_records_with_resource" {
  provider         = nios
  include_resource = true
}
