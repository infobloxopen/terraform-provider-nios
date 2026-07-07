// Manage Microsoft Server with Basic Fields
resource "nios_microsoft_msserver" "msserver_basic" {
  address    = "10.10.0.1"
  login_name = "example_login"
}

// Manage Microsoft Server with Additional Fields
resource "nios_microsoft_msserver" "msserver_with_additional_fields" {
  address    = "10.10.0.2"
  login_name = "example_login2"

  // Additional Fields
  comment                   = "Example Microsoft Server"
  synchronization_min_delay = 5

  //Extensible Attributes
  extattrs = {
    Site = "location-1"
  }
}

// Manage Microsoft Server with AD Sites
resource "nios_microsoft_msserver" "msserver_with_ad_sites" {
  address    = "10.10.0.3"
  login_name = "example_login3"

  // Additional Fields
  comment                   = "Example Microsoft Server"
  synchronization_min_delay = 5

  ad_sites = {
    login_name                    = "example_login4"
    use_login                     = true
    synchronization_min_delay     = 10
    use_synchronization_min_delay = true
  }
}
