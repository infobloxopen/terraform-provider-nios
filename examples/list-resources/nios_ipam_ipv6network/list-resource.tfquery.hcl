// List specific IPv6 Networks using filters
list "nios_ipam_ipv6network" "list_ipv6networks_using_filters" {
  provider = nios
  config {
    filters = {
      network = "2001:db8::/32"
    }
  }
}

// List specific IPv6 Networks using Extensible Attributes
list "nios_ipam_ipv6network" "list_ipv6networks_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List IPv6 Networks with resource details included
list "nios_ipam_ipv6network" "list_ipv6networks_with_resource" {
  provider         = nios
  include_resource = true
}
