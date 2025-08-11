// Create an Admin User with Basic Fields
resource "nios_security_admin_user" "admin_user_basic_fields" {
  name         = "example_admin_user"
  password     = "Example-password1!1223a"
  admin_groups = ["admin-group"]
}

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
  time_zone     = "singapore"
  use_time_zone = true
  #  ca_certificate_issuer = ""
  #  client_certificate_serial_number = ""
  #  enable_certificate_authentication = true
  ssh_keys = [
    {
      key_name  = "sample-key"
      key_type  = "RSA"
      key_value = "-----BEGIN RSA PRIVATE KEY-----MIIEpAIBAAKCAQEAv1ZzY9kzv+3x9Qk3vWvYz3JkQ9vZz8X9Qk3vWvYz3JkQ9vZz8X9Qk3vWvYz3JkQ9vZz8X9Qk3vWvYz3JkQ9vZz8X9Qk3vWvYz3JkQ9vZz8X9Qk3vWvYz3JkQ9vZz8X9Qk3vWvYz3JkQ9vZz8X9Qk3vWvYz3JkQ9vZz8X9QIDAQABAoIBAQCDUMMYKEYDATA1234567890abcdefg==-----END RSA PRIVATE KEY-----"
    }
  ]
  use_ssh_keys = true
}