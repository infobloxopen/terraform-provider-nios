// Create DTC LBDN with basic fields
resource "nios_dtc_lbdn" "lbdn_basic_fields" {
  name      = "example_lbdn_1"
  lb_method = "SOURCE_IP_HASH"
}

// Create DTC LBDN with additional fields
resource "nios_dtc_lbdn" "lbdn_additional_fields" {
  name = "example_lbdn_2"
  auth_zones = ["zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5yZWNvcmRfdGVzdA:wapi.com/default",
    "zone_auth/ZG5zLnpvbmUkLjEuY29tLnRlc3Q:info.com/default.custom_view"
  ]
  comment = "lbdn with additional parameters"
  extattrs = {
    Site = "Yoshino"
  }
  lb_method = "TOPOLOGY"
  patterns  = ["*wapi.com", "info.com*"]
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
      pool  = "dtc:pool/ZG5zLmlkbnNfcG9vbCR0ZXN0LXBvb2w:pool6"
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
