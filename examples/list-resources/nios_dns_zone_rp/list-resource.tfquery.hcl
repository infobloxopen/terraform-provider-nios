// List specific Response Policy Zones using filters
list "nios_dns_zone_rp" "list_response_policy_zones_using_filters" {
  provider = nios
  config {
    filters = {
      fqdn = "example.com"
    }
  }
}

// List specific Response Policy Zones using Extensible Attributes
list "nios_dns_zone_rp" "list_response_policy_zones_using_extensible_attributes" {
  provider = nios
  config {
    extattrfilters = {
      Site = "location-1"
    }
  }
}

// List Response Policy Zones with resource details included
list "nios_dns_zone_rp" "list_response_policy_zones_with_resource" {
  provider         = nios
  include_resource = true
}
