// Create a SAML Authservice with Basic Fields
resource "nios_security_saml_authservice" "saml_authservice_with_basic_fields" {
  name = "saml_authservice"
  idp = {
    idp_type           = "OKTA"
    metadata_file_path = "<path-to-the-metadata-file>"
  }
}

//Create a SAML Authservice with Additional Fields
resource "nios_security_saml_authservice" "saml_authservice_with_additional_fields" {
  name    = "saml_authservice_2"
  comment = "Example SAML Authservice with Additional Fields"
  idp = {
    idp_type           = "OKTA"
    metadata_file_path = "<path-to-the-metadata-file>"
    groupname          = "group1"
    comment            = "IDP for SAML Authserv232ce"
    sso_redirect_url   = "2.2.2.2"
  }
  session_timeout = 120
}