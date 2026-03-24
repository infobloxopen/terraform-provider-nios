resource "nios_grid_member" "test_enable_ha" {
  host_name                  = "infoblox-ha-node.localdomain"
  config_addr_type           = "IPV4"
  platform                   = "VNIOS"
  service_type_configuration = "ALL_V4"

  ipv6_setting = {
    auto_router_config_enabled = false
    enabled                    = false
    primary                    = true
    use_dscp                   = false
    dscp                       = 0
  }

  vip_setting = {
    address     = "172.28.83.101"
    gateway     = "172.28.82.1"
    subnet_mask = "255.255.254.0"
    primary     = true
    use_dscp    = false
    dscp        = 0
  }
  use_snmp_setting = true
  syslog_proxy_setting = {
    client_acls = [
      {
        struct     = "addressac"
        address    = "192.0.0.1"
        permission = "ALLOW"
      }
    ]
    enable     = false
    tcp_enable = false
    tcp_port   = 514
    udp_enable = true
    udp_port   = 514
  }
  use_syslog_proxy_setting = true
}
