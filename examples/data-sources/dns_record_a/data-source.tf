data "nios_dns_record_a" "get_record" {
    extattrfilters =  {
    "Site" = "Siteblr"
  }
}