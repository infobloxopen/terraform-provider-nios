// Retrieve a specific ruleset by filters
data "nios_misc_ruleset" "get_ruleset_by_name" {
  filters = {
    name = "example_ruleset_1"
  }
}

// Retrieve all rulesets
data "nios_misc_ruleset" "get_all_rulesets" {}
