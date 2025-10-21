import {
  to = nios_ip_allocation.allocation1
  id = "record:host/ZG5zLmhvc3QkLl9kZWZhdWx0LmNvbS5leGFtcGxlLnNhbXBsZV9yZWNvcmQ:sample_record.example.com/default"
}

resource "nios_ip_allocation" "allocation1" {
  name = "sample_record.example.com"
  view = "default"
  ipv4addrs = [
    {
      ipv4addr = "10.101.1.110"
    }
  ]
}

import {
  to = nios_ip_association.association1
  id = "record:host/ZG5zLmhvc3QkLl9kZWZhdWx0LmNvbS5leGFtcGxlLnNhbXBsZV9yZWNvcmQ:sample_record.example.com/default"
}

resource "nios_ip_association" "association1" {
  ref                = nios_ip_allocation.allocation1.ref
  mac                = "12:00:43:fe:9a:8c"
  configure_for_dhcp = false
  internal_id        = nios_ip_allocation.allocation1.internal_id
}
