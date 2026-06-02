// List specific Network Templates using filters
list "nios_ipam_networktemplate" "list_network_templates_using_filters" {
  provider = nios
  config {
    filters = {
      name = "example_network_template"
    }
  }
}

// List specific Network Templates using Extensible Attributes
list "nios_ipam_networktemplate" "list_network_templates_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Network Templates with resource details included
list "nios_ipam_networktemplate" "list_network_templates_with_resource" {
  provider         = nios
  include_resource = true
}