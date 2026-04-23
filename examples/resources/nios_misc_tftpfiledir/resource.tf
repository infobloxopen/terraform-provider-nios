// Manage TFTP file/directory object with Basic Fields
resource "nios_misc_tftpfiledir" "misc_tftpfiledir_basic" {
  name = "example-tftpfiledir"
  type = "DIRECTORY"
}

// Manage TFTP file/directory object with Additional Fields
// Additional Configuration required to allow the client `10.0.0.0/24` to perform file transfers for the member `infoblox.localdomain`
resource "nios_misc_tftpfiledir" "misc_tftpfiledir_with_additional_fields" {
  name      = "example-tftpfiledir2"
  type      = "DIRECTORY"
  directory = "/ftpusers"
  vtftp_dir_members = [
    {
      member  = "infoblox.localdomain"
      ip_type = "NETWORK"
      network = "10.0.0.0"
      cidr    = 24
    }
  ]
}
