// Manage Microsoft Server ADSites with Basic Fields
resource "nios_microsoft_msserver_adsites_site" "msserver_adsites_basic" {
  domain = "msserver:adsites:domain/ZG5zLm1zX2FkX3NpdGVzX2RvbWFpbiQwLkFELTE3MA:example.local/default"
  name   = "example_adsite_1"
}


// Create an IPAM Network(Required as parent)
resource "nios_ipam_network" "example_network" {
  network = "13.0.0.0/24"
}


// Manage Microsoft Server ADSites with Additional Fields
resource "nios_microsoft_msserver_adsites_site" "msserver_adsites_additional" {
  domain = "msserver:adsites:domain/ZG5zLm1zX2FkX3NpdGVzX2RvbWFpbiQwLkFELTE3MA:example.local/default"
  name   = "example_adsite_2"
  networks = [
    nios_ipam_network.example_network.ref
  ]
}
