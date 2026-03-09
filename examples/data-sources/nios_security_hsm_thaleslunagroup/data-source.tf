# Retrieve a specific Thales Luna HSM Group by filters
data "nios_security_hsm_thaleslunagroup" "get_hsm_thaleslunagroup_using_filter" {
  filters = {
    name = "example-thalesluna-hsm"
  }
}

# Retrieve all Thales Luna HSM Groups
data "nios_security_hsm_thaleslunagroup" "get_all_hsm_thaleslunagroups" {}
