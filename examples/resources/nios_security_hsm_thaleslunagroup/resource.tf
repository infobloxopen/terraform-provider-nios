# Create a Thales Luna HSM Group with Basic Fields
resource "nios_security_hsm_thaleslunagroup" "hsm_thaleslunagroup_with_basic_fields" {
  name        = "example-thalesluna-hsm1"
  hsm_version = "Luna_7_CPL"
  pass_phrase = "< Enter Password for the Luna HSM >"

  thalesluna = [
    {
      name                    = "< Enter Name for Luna HSM Device 1>"
      partition_serial_number = "< Enter Partition Serial Number >"
      server_cert_file_path   = "< Enter Path to Server Certificate >"
    }
  ]
}

# Create a Thales Luna HSM Group with Additional Fields
resource "nios_security_hsm_thaleslunagroup" "hsm_thaleslunagroup_with_additional_fields" {
  name        = "example-thalesluna-hsm"
  hsm_version = "Luna_7_CPL"
  pass_phrase = "< Enter Password for the Luna HSM >"

  comment = "Group for Thales Luna 7 HSM1 "

  thalesluna = [
    {
      name                    = "< Enter Name for Luna HSM Device 1>"
      partition_serial_number = "< Enter Partition Serial Number >"
      server_cert_file_path   = "< Enter Path to Server Certificate >"
      disable                 = false
    }
  ]
}
