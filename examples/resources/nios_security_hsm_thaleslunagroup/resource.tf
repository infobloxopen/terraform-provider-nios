// Create a Thales Luna HSM Group with Basic Fields
resource "nios_security_hsm_thaleslunagroup" "hsm_thaleslunagroup_with_basic_fields" {
  name        = "example-thalesluna-hsm1"
  hsm_version = "Luna_7_CPL"
  pass_phrase = "examplePass123"

  thalesluna = [
    {
      name                    = "10.0.0.1"
      partition_serial_number = "123456789"
      server_cert_file_path   = "/path/to/server.pem"
    }
  ]
}

// Create a Thales Luna HSM Group with Additional Fields
resource "nios_security_hsm_thaleslunagroup" "hsm_thaleslunagroup_with_additional_fields" {
  name        = "example-thalesluna-hsm"
  hsm_version = "Luna_7_CPL"
  pass_phrase = "examplePass123"

  comment = "Group for Thales Luna 7 HSM1 "

  thalesluna = [
    {
      name                    = "10.0.0.1"
      partition_serial_number = "123456789"
      server_cert_file_path   = "/path/to/server.pem"
      disable                 = false
    }
  ]
}
