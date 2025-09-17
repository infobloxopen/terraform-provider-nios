// Retrieve a specific BFD Template by filters
data "nios_misc_bfdtemplate" "bfd_template_with_filters" {
  filters = {
    name = "example_bfd_name"
  }
}

// Retrieve all specific BFD Template
data "nios_misc_bfdtemplate" "get_all_bfd_template" {}
