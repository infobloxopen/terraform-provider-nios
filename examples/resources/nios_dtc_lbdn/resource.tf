// Create DTC LBDN with basic fields
resource "nios_dtc_lbdn" "lbdn_basic_fields" {
  name      = "example_lbdn_1"
  lb_method = "SOURCE_IP_HASH"
  types     = ["A", "CNAME"]
}

// Authoritative DNS zone in the default view with a grid primary association
resource "nios_dns_zone_auth" "parent_zone" {
  fqdn = "wapi.com"
  view = "default"
  grid_primary = [
    {
      name = "infoblox.localdomain",
    }
  ]
}

// Create a custom DNS view
resource "nios_dns_view" "custom_view1" {
  name = "custom_view"
}

// Authoritative DNS zone in the custom view with a grid primary association
resource "nios_dns_zone_auth" "parent_zone2" {
  fqdn = "info.com"
  view = nios_dns_view.custom_view1.name
  grid_primary = [
    {
      name = "infoblox.localdomain",
    }
  ]
}

// Define the DTC pools for LBDN association
resource "nios_dtc_pool" "dtc_pool1" {
  name                = "pool2"
  lb_preferred_method = "ROUND_ROBIN"
}

resource "nios_dtc_pool" "dtc_pool2" {
  name                = "pool4"
  lb_preferred_method = "ROUND_ROBIN"
}

resource "nios_dtc_pool" "dtc_pool3" {
  name                = "pool6"
  lb_preferred_method = "ROUND_ROBIN"
}

// Create DTC LBDN with additional fields
resource "nios_dtc_lbdn" "lbdn_additional_fields" {
  name = "example_lbdn_2"
  auth_zones = [nios_dns_zone_auth.parent_zone.ref,
    nios_dns_zone_auth.parent_zone2.ref
  ]
  comment = "lbdn with additional parameters"
  extattrs = {
    Site = "location-1"
  }
  lb_method = "TOPOLOGY"
  //The topology used here must have any one of the pools configured in its topology members
  topology = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wbzE:topo1"
  patterns = ["*wapi.com", "info.com*"]
  pools = [
    {
      pool  = nios_dtc_pool.dtc_pool1.ref
      ratio = 2
    },
    {
      pool  = nios_dtc_pool.dtc_pool2.ref
      ratio = 3
    },
    {
      pool  = nios_dtc_pool.dtc_pool3.ref
      ratio = 6
    }
  ]
  ttl         = 0
  use_ttl     = false
  disable     = true
  types       = ["A", "CNAME"]
  persistence = 100
  priority    = 1
}
