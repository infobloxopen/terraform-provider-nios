terraform {
  required_providers {
    nios = {
      source  = "infobloxopen/nios"
      version = "1.0.0"
    }
  }
}

provider "nios" {
  nios_host_url = "https://172.28.82.213"
  nios_username = "admin"
  nios_password = "Infoblox@123"
}


// Create Certificate Authservice with basic fields
resource "nios_security_certificate_authservice" "certificate_authservice_with_basic_fields" {
  name                    = "example_certificate_authservice2"
  ca_certificates         = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22"]
  ocsp_check              = "DISABLED"
  remote_lookup_service   = "ad_auth_service/b25lLmFkX2F1dGhfc2VydmljZSRhY3RpdmVfZGly:active_dir"
  enable_remote_lookup    = true
  remote_lookup_password  = "example-password"
  # remote_lookup_username  = "administrator"
  enable_password_request = false

}

// Create Certificate Authservice with additional fields
resource "nios_security_certificate_authservice" "certificate_authservice_with_additional_fields" {
  name            = "example_certificate_authservice3"
  ca_certificates = ["cacertificate/id"]

  //additonal fields
  comment                 = "Example Certificate Authservice with additional fields"
  auto_populate_login     = "SERIAL_NUMBER"
  remote_lookup_service   = "ad_auth_service/id"
  enable_remote_lookup    = true
  enable_password_request = false
  remote_lookup_password  = "example-password"
  remote_lookup_username  = "administrator"
  ocsp_check              = "MANUAL"
  ocsp_responders = [
    {
      certificate_file_path = "<path-to-the-certificate-file>"
      fqdn_or_ip            = "2.2.2.2"
    }
  ]
  recovery_interval = 20
  response_timeout  = 2000
  trust_model       = "DELEGATED"
  user_match_type   = "AUTO_MATCH"
  max_retries       = 2
  disabled          = false
}
