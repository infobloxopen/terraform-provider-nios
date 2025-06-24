// forward-zone with forward_to and view
resource "nios_dns_zone_forward" "rec1" {
  fqdn     = "example_test1234.example.com"
  forward_to = [
    {
      name = "ns1.example.com"
      address = "1.1.1.1"
    }
  ]
}

//forward zone with full set of parameters
resource "nios_dns_zone_forward" "forward_zone_full_parameters" {
  fqdn = "max_params.ex.org"
  view = "custom_view"
  zone_format = "FORWARD"
  comment = "test sample forward zone"
  forward_to {
    name = "te32.dz.ex.com"
    address = "10.0.0.1"
  }
  forwarding_servers {
    name = "infoblox.172_28_82_248"
    forwarders_only = true
    use_override_forwarders = false
    forward_to {
      name = "cc.fwd.com"
      address = "10.1.1.1"
    }
  }
  forwarding_servers = [
    {
      name = "infoblox.172_28_82_248"
      forwarders_only = true
      use_override_forwarders = true
      forward_to = [
        {
          name = "kk.fwd.com"
          address = "10.2.1.31"
        }
      ]
    }
  ]
  extattrs = {
    Site = "Siteblr"
  }
}

//forward zone with ns_group and external_ns_group
resource "nios_dns_zone_forward" "forward_zone_nsGroup_externalNsGroup" {
  fqdn = "params_ns_ens.ex.org"
  ns_group = "fwd_member"
  external_ns_group = "nsg1"
}