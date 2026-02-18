# Associate MAC address with host1's IP without DHCP configuration
resource "nios_ip_association" "association1" {
  ref                = nios_ip_allocation.allocation1.ref
  mac                = "12:00:43:fe:9a:8c"
  configure_for_dhcp = false
}

# Associate MAC and DUID with host2's IP with DHCP configuration enabled
resource "nios_ip_association" "association2" {
  ref                = nios_ip_allocation.allocation2.ref
  mac                = "12:43:fd:ba:9c:c9"
  duid               = "00:01:5f:3a:1b:2c:12:34:56:78:9a:bc"
  match_client       = "DUID"
  configure_for_dhcp = true
}

# Associate MAC address with host3's IP without DHCP configuration
resource "nios_ip_association" "association3" {
  ref                = nios_ip_allocation.allocation3.ref
  mac                = "12:00:43:fe:9a:8d"
  configure_for_dhcp = false
}
