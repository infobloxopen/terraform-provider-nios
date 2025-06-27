resource "nios_resource_nios_DtcLbdn" "lbdn1" {
  name = "testLbdn2"
  lb_method = "ROUND_ROBIN"
  #  topology = "test-topo"
  types = ["A", "AAAA","NAPTR"]
}


resource "nios_resource_nios_DtcLbdn" "lbdn2" {
  name = "lbdn123"
  auth_zones = ["zone_auth/ZG5zLnpvbmUkLjEuY29tLnRlc3Qx:test1.com/default.view2",
    #    "zone_auth/ZG5zLnpvbmUkLjEuY29tLnRlc3Q:test.com/default.view2",
    "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS50ZXN0:test.com/default"
  ]
  comment = "test"
  extattrs = {
    "Site" = {
      value = "Kyoto1"
    }
  }
  lb_method = "TOPOLOGY"
  patterns = ["test.com","test1.com*"]
  pools = [
    {
      pool  = "dtc:pool/ZG5zLmlkbnNfcG9vbCRwb29sMg:pool2"
      ratio = 2
    },
    #    {
    #      pool  = "dtc:pool/ZG5zLmlkbnNfcG9vbCRwb29sNA:pool4"
    #      ratio = 3
    #    },
    {
      pool  = "dtc:pool/ZG5zLmlkbnNfcG9vbCR0ZXN0LXBvb2w:test-pool"
      ratio = 6
    }
  ]
  ttl = 0
  use_ttl = false
  topology = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wbzE:topo1"
  disable = true
  types = ["A", "CNAME"]
  persistence = 100
  priority = 1
}