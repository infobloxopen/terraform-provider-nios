// Retrieve a specific Microsoft Server Adsites using filters 
data "nios_microsoft_msserver_adsites_site" "get_msserver_adsites_using_filters" {
  filters = {
    name = "example_adsite_1"
  }
}

// Retrieve all Microsoft Servers
data "nios_microsoft_msserver_adsites_site" "get_all_msserver_adsites" {}
