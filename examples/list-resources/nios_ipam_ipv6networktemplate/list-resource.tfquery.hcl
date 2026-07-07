// List specific IPv6 Network Templates using filters
list "nios_ipam_ipv6networktemplate" "list_templates_using_filters" {
  provider = nios
  config {
    filters = {
      name = "network-template-example"
    }
  }
}

// List specific IPv6 Network Templates using Extensible Attributes
list "nios_ipam_ipv6networktemplate" "list_templates_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      "Tenant ID" = "tenant-123"
    }
  }
}

// List IPv6 Network Templates with resource details included
list "nios_ipam_ipv6networktemplate" "list_templates_with_resource" {
  provider         = nios
  include_resource = true
  limit            = 5
}
