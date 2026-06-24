// List specific Bulk Hostname Templates using filters
list "nios_ipam_bulk_hostname_template" "list_bulk_hostname_templates_using_filters" {
  provider = nios
  config {
    filters = {
      template_name = "example_template"
    }
  }
}

// List Bulk Hostname Templates with resource details included
list "nios_ipam_bulk_hostname_template" "list_bulk_hostname_templates_with_resource" {
  provider         = nios
  include_resource = true
}

