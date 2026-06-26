// List specific Networks using filters
list "nios_ipam_network" "list_networks_using_filters" {
  provider = nios
  config {
    filters = {
      network = "10.0.0.0/24"
    }
  }
}

// List specific Networks using Extensible Attributes
list "nios_ipam_network" "list_networks_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Networks with resource details included
list "nios_ipam_network" "list_networks_with_resource" {
  provider         = nios
  include_resource = true
}
