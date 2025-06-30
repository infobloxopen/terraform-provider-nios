resource "nios_dtc_lbdn" "lbdn_record" {
  name = "testLbdn22"
  lb_method = "SOURCE_IP_HASH"
  extattrs = {
    Site = "Spain"
  }
}

data "nios_dtc_lbdn" "get_record_with_filter" {
  filters = {
    "name" = nios_dtc_lbdn.lbdn_record.name
  }
}

data "nios_dtc_lbdn" "get_record_with_extattr_filter" {
  extattrfilters = {
    "Site" = "Spain"
  }
}