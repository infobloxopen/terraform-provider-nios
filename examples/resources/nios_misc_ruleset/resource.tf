//Create a Ruleset with Basic Fields
resource "nios_misc_ruleset" "ruleset_basic" {
  name = "example_ruleset_1"
  type = "BLACKLIST"
}

//Create a Ruleset with Additional Fields
resource "nios_misc_ruleset" "ruleset_with_additional_fields" {
  name     = "example_ruleset_2"
  type     = "NXDOMAIN"
  comment  = "This ruleset handles NXDOMAIN redirection"
  disabled = false

  nxdomain_rules = [
    {
      action  = "PASS"
      pattern = "example.com"
    },
    {
      action  = "REDIRECT"
      pattern = "test.org"
    }
  ]
}
