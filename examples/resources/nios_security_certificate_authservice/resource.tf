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

resource "nios_security_certificate_authservice" "certificate_authservice_basic_fields"{
    name = "example_certificate_authservice2"
    ca_certificates = ["cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22"]
  remote_lookup_service = "ad_auth_service/b25lLmFkX2F1dGhfc2VydmljZSRhY3RpdmUrZGly:active%2Bdir"
    #file_name = "/Users/chaithra/go/src/github.com/infobloxopen/terraform-provider-nios/internal/utils/cert.pem"
    # ocsp_responders =[
    #   {
    #     fqdn_or_ip = "2.2.2.2"
    #   }
    # ]
}