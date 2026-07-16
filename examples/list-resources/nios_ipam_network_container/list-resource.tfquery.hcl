// List specific Network Containers using filters
list "nios_ipam_network_container" "list_containers_using_filters" {
  provider = nios
  config {
    filters = {
      network = "10.0.0.0/8"
    }
  }
}

// List specific Network Containers using Extensible Attributes
list "nios_ipam_network_container" "list_containers_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Network Containers with resource details included
list "nios_ipam_network_container" "list_containers_with_resource" {
  provider         = nios
  include_resource = true
}
