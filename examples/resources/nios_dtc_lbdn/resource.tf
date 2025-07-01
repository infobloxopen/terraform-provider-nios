resource "nios_dtc_lbdn" "lbdn1" {
  name      = "testLbdn21"
  lb_method = "SOURCE_IP_HASH"
}


resource "nios_dtc_lbdn" "lbdn2" {
  name = "lbdn123"
  auth_zones = ["zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5yZWNvcmRfdGVzdA:record_test.com/default",
    "zone_auth/ZG5zLnpvbmUkLjEuY29tLnRlc3Q:test.com/default.custom_view"
  ]
  comment = "test"
  extattrs = {
    Site = "Yoshino"
  }
  lb_method = "TOPOLOGY"
  patterns  = ["*record_test.com", "test.com*"]
  pools = [
    {
      pool  = "dtc:pool/ZG5zLmlkbnNfcG9vbCRwb29sMg:pool2"
      ratio = 2
    },
    {
      pool  = "dtc:pool/ZG5zLmlkbnNfcG9vbCRwb29sNA:pool4"
      ratio = 3
    },
    {
      pool  = "dtc:pool/ZG5zLmlkbnNfcG9vbCR0ZXN0LXBvb2w:test-pool"
      ratio = 6
    }
  ]
  ttl         = 0
  use_ttl     = false
  topology    = "dtc:topology/ZG5zLmlkbnNfdG9wb2xvZ3kkdG9wbzE:topo1"
  disable     = true
  types       = ["A", "CNAME"]
  persistence = 100
  priority    = 1
}