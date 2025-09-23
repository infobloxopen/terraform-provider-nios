// Retrieve a specific certificate Authservice by filters
data "nios_security_certificate_authservice" "get_certificate_authservice_using_filters" {
  filters = {
    name = "example_certificate_authservice2"
  }
}

// Retrieve all certificate Authservices
data "nios_security_certificate_authservice" "get_all_certificate_authservices" {}
