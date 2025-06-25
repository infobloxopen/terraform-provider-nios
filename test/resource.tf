resource "nios_resource_nios_Networkcontainer" "example_container" {
  network      = "10.0.0.0/24"
  network_view = "default"
  comment      = "test comment"

  // Optional: Configure extensible attributes
  extattrs = {
    "Site" = {
      value = "DataCenter1"
      flags = {}
    }
  }
}

# resource "nios_resource_nios_Network" "example_network" {
#   network      = "10.0.0.0/30"
#   network_view = "default"
#   comment      = "test comment"

#   // Optional: Configure extensible attributes
#   extattrs = {
#     "Site" = {
#       value = "DataCenter1"
#       flags = {}
#     }
#   }
# }
