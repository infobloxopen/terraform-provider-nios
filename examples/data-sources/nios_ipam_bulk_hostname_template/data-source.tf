// Retrieve a specific IPAM Bulk Hostname Templates by filters
data "nios_ipam_bulk_hostname_template" "bulk_hostname_template" {
  filters = {
    template_name = "one-octet"
  }
}

// Retrieve all IPAM Bulk Hostname Templates
data "nios_ipam_bulk_hostname_template" "all_templates" {}
