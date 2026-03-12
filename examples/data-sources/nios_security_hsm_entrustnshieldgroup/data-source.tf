// Retrieve an Entrust nShield HSM Group by filters
data "nios_security_hsm_entrustnshieldgroup" "get_hsm_entrustnshieldgroup_using_filters" {
  filters = {
    name = "example-hsm-entrustnshieldgroup"
  }
}

// Retrieve all Entrust nShield HSM Groups
data "nios_security_hsm_entrustnshieldgroup" "get_all_hsm_entrustnshieldgroups" {}
