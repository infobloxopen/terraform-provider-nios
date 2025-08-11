// Retrieve IPAM bulk hostname templates using filters
data "nios_ipam_bulkhostnametemplate" "test" {
  filters = {
    template_name   = nios_ipam_bulkhostnametemplate.test2.template_name
    template_format = nios_ipam_bulkhostnametemplate.test2.template_format
  }
}

// Retrieve all IPAM bulk hostname templates
data "nios_ipam_bulkhostnametemplate" "all_templates" {}
