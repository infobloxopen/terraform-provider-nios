// Create an Admin User with Basic Fields
resource "nios_security_admin_user" "admin_user_basic_fields" {
  name         = "example_admin_user"
  password     = "Example-password1!"
  admin_groups = ["admin-group"]
}

// TODO: Retrieve references based on the provided distinguished name of the object: cacertificate
// Create an Admin User with Additional Fields
resource "nios_security_admin_user" "admin_user_additional_fields" {
  name         = "example_admin_user2"
  admin_groups = ["cloud-api-only"]
  password     = "Example-password1!"
  auth_type    = "LOCAL"
  auth_method  = "KEYPAIR_PASSWORD"
  comment      = "Example Admin User with additional fields"
  email        = "admin-user2@example.com"
  disable      = false
  extattrs = {
    Site = "location-1"
  }
  time_zone                         = "Asia/Kolkata"
  use_time_zone                     = true
  ca_certificate_issuer             = "cacertificate/b25lLmVhcF9jYV9jZXJ0JDAuNzg5Y2IyOGVkZDgyMDE5MTYzODljOGQ5MGI2MTM4YmFlNDIxODY1YmY2YWZlMTdiMmEyNDRjNTIwNDRkMGQ3NWFiMGY0MGFjNTBmYzc3ZGMwM2YwOTI2NWRhNDRkYzllMjQ0OTBkZmMyMWEyOWVlYmIxODhlMDFlMWY5OGYwOTg:CN%3D%22ib-root-ca%22"
  client_certificate_serial_number  = "4e7c675cd972ecd2e5b895ad6cb4e38e6d77b4b4"
  enable_certificate_authentication = true
  ssh_keys = [
    {
      key_name = "sample-key"
      key_type = "RSA"
      // A Public key file is required for the example to function
      key_value = replace(file("${path.module}/sample_key.pub"), "\n", "")
    }
  ]
  use_ssh_keys = true
}
