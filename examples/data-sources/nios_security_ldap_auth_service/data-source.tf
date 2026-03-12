// Retrieve a specific LDAP Authservice by filters
data "nios_security_ldap_auth_service" "get_ldap_authservice_using_filters" {
  filters = {
    name = "example_ldap_authservice"
  }
}

// Retrieve all LDAP Authservices
data "nios_security_ldap_auth_service" "get_all_ldap_authservices" {}
