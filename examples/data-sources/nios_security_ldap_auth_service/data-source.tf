// Retrieve a specific LDAP Auth Service by filters
data "nios_security_ldap_auth_service" "get_ldap_authservice_using_filters" {
  filters = {
    name = "example_ldap_authservice"
  }
}

// Retrieve all LDAP Auth Services
data "nios_security_ldap_auth_service" "get_all_ldap_authservices" {}
