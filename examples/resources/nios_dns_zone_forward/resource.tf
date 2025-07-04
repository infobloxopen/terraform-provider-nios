resource "nios_dns_zone_forward" "rec1" {
  fqdn = "example_test1.example.com"
  forward_to = [
    {
      name = "ns1.example.com"
      address = "1.1.1.1"
    }
  ]
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
  view     = "default"
  extattrs = {
    Site = "Hungary"
  }
}