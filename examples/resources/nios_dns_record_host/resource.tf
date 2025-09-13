resource "nios_dns_record_host" "example" {
  name = "test-host-pm3.test-zone.com"
  view = "default"

  ipv4addrs = [
    {
      ipv4addr = "10.101.0.1"
    }
  ]
  # ipv6addrs = [
  #   {
  #     ipv6addr = "2002:1f93::12:2"
  #   }
  # ]
  # ipv4addrs = [
  #   {
  #     func_call = {
  #       attribute_name  = "ipv4addr"
  #       object_function = "next_available_ip"
  #       result_field    = "ips"
  #       object          = "network"
  #       object_parameters = {
  #         network      = "10.10.0.0/16"
  #         network_view = "default"
  #       }
  #     }
  #   }
  # ]
}

resource "nios_dns_ip_association" "example_assoc" {
  ref = nios_dns_record_host.example.ref
  mac = "aa:bb:cc:dd:ee:ff"
  # duid = "00:43:d2:0a:11:e6"
  configure_for_dhcp = true
  internal_id        = nios_dns_record_host.example.internal_id
}
