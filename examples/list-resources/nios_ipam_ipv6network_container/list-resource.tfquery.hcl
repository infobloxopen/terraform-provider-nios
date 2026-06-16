// List all IPv6 Network Containers
list "nios_ipam_ipv6network_container" "test" {
  provider         = nios
  include_resource = true
}

// List specific IPv6 Network Containers using filters
list "nios_ipam_ipv6network_container" "list_containers_using_filters" {
  provider = nios
  config {
    filters = {
      network = "2001:db8::/32"
    }
  }
}

// List specific IPv6 Network Containers using Extensible Attributes
list "nios_ipam_ipv6network_container" "list_containers_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List IPv6 Network Containers with resource details included
list "nios_ipam_ipv6network_container" "list_containers_with_resource" {
  provider         = nios
  include_resource = true
}
