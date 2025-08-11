#create a basic bulk host name template
resource "nios_ipam_bulkhostnametemplate" "test" {
  template_name   = "one-octet"
  template_format = "-$4"
}

#create a bulk host name template with two octets
resource "nios_ipam_bulkhostnametemplate" "test2" {
  template_name   = "two-octet"
  template_format = "host-$3-$4"
}
