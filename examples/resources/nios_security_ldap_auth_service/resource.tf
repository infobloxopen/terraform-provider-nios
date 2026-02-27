// Create LDAP Authservice with Basic Fields
resource "nios_security_ldap_auth_service" "ldap_authservice_with_basic_fields" {
  name = "example_ldap_authservice"
  servers = [
    {
      address             = "2.2.2.2"
      authentication_type = "ANONYMOUS"
      base_dn             = "ou=People,dc=example,dc=com"
      disable             = false
      encryption          = "SSL"
      port                = 636
      use_mgmt_port       = false
      version             = "V3"
    }
  ]
  ldap_user_attribute = "adminID"
  recovery_interval   = 60
  retries             = 3
  timeout             = 60
}

// Create LDAP Authservice with Additional Fields
resource "nios_security_ldap_auth_service" "ldap_authservice_with_additional_fields" {
  name    = "example_ldap_authservice2"
  comment = "This is a comment"
  disable = false
  ea_mapping = [
    {
      mapped_ea = "Availability zone"
      name      = "ldapfield"
    }
  ]
  ldap_group_attribute           = "namecn"
  ldap_group_authentication_type = "GROUP_ATTRIBUTE"
  ldap_user_attribute            = "adminID"
  mode                           = "ROUND_ROBIN"
  search_scope                   = "SUBTREE"
  servers = [
    {
      address             = "2.2.2.2"
      authentication_type = "ANONYMOUS"
      base_dn             = "ou=People,dc=example,dc=com"
      disable             = false
      encryption          = "SSL"
      port                = 636
      use_mgmt_port       = false
      version             = "V3"
    }
  ]
  recovery_interval = 60
  retries           = 3
  timeout           = 60
}
