// Retrieve a specific SAML Authservice by filters
data "nios_security_saml_authservice" "get_saml_authservice_using_filters" {
  filters = {
    name = "saml_authservice"
  }
}

// Retrieve all Certificate Authservices
data "nios_security_saml_authservice" "get_all_saml_authservices" {}
