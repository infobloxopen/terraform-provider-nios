// Manage TFTP file/directory object with Basic Fields
resource "nios_misc_tftpfiledir" "misc_tftpfiledir_basic" {
  name = "example-tftpfiledir"
  type = "DIRECTORY"
}

// Manage TFTP file/directory object with Additional Fields
resource "nios_misc_tftpfiledir" "misc_tftpfiledir_with_additional_fields" {
  name = "example-tftpfiledir2"
  type = "FILE"
  members = [
    {
      member  = "infoblox.localdomain"
      ip_type = "NETWORK"
      network = "192.16.1.1"
      cidr    = 24
    }
  ]
}
