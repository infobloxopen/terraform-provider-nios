// Retrieve a specific BFD Template by filters
data "nios_misc_bfdtemplate" "test_bfdtemplate_filters" {
  filters = {
    name = "tf_test_bfd_name"
  }
}
