resource "nios_dtc_pool" "dtc_pool1"{
    name = "dtc_pool23"
    lb_preferred_method="ROUND_ROBIN"
    extattrs = {
    Site = "Siteblr"
  }
}

data "nios_dtc_pool" "test" {
  filters = {
    name = nios_dtc_pool.dtc_pool1.name 
  }
}

data "nios_dtc_pool" "test2" {
  extattrfilters = {
	"Site" = nios_dtc_pool.dtc_pool1.extattrs.Site
  }
}