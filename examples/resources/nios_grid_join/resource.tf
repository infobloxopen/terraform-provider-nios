// Create an Offline Member with IPV4 config
resource "nios_grid_member" "example_grid_member_1" {
  host_name        = "infoblox.member1"
  config_addr_type = "IPV4"
  platform         = "VNIOS"
  vip_setting = {
    address     = "172.28.82.231"
    gateway     = "172.28.82.1"
    subnet_mask = "255.255.254.0"
  }
}

// Join the member to the grid master
resource "nios_grid_join" "member_join_1" {
  member_url      = "https://172.28.83.231"
  member_username = "username"
  member_password = "password"
  grid_name       = "Infoblox"
  master          = "172.28.82.32"
  shared_secret   = "secret"
  depends_on      = [nios_grid_member.example_grid_member_1]
}

// Create an Offline Member with IPV6 config
resource "nios_grid_member" "example_grid_member_2" {
  host_name        = "infoblox.member2"
  config_addr_type = "IPV6"
  platform         = "VNIOS"
  ipv6_setting = {
    virtual_ip  = "2600:1f1c:e86:5e01:a1df:ceec:98c4:2b30"
    cidr_prefix = 64
    gateway     = "fe80::4e5:f7ff:fe5d:84e3"
  }
}

// Join the member to the grid master
resource "nios_grid_join" "member_join_2" {
  member_url      = "https://172.28.83.21"
  member_username = "username"
  member_password = "password"
  grid_name       = "Infoblox"
  master          = "2600:1f1c:e86:5e01:fb1d:dc76:3c28:ec48"
  shared_secret   = "secret"
  depends_on      = [nios_grid_member.example_grid_member_2]
}

// Create an Offline Member with both IPV4 and IPV6 config
resource "nios_grid_member" "example_grid_member_3" {
  host_name        = "infoblox.member3"
  config_addr_type = "BOTH"
  platform         = "VNIOS"
  vip_setting = {
    address     = "172.28.83.21"
    gateway     = "172.28.82.1"
    subnet_mask = "255.255.254.0"
  }
  ipv6_setting = {
    virtual_ip  = "2600:1f1c:e86:5e01:a1df:ceec:98c4:2b30"
    cidr_prefix = 64
    gateway     = "fe80::4e5:f7ff:fe5d:84e3"
  }
}

resource "nios_grid_join" "member_join_3" {
  member_url      = "https://172.28.83.21"
  member_username = "username"
  member_password = "password"
  grid_name       = "Infoblox"
  master          = "172.28.82.171"
  shared_secret   = "secret"
  depends_on      = [nios_grid_member.example_grid_member_3]
}

// For SA-HA config, import the existing grid master and update the resource to set ha_on_cloud to true and provide the cloud attributes.
// This forms the HA with one Active and one Passive node.
resource "nios_grid_member" "example_grid_member_4" {
  host_name         = "infoblox.localdomain"
  config_addr_type  = "IPV4"
  platform          = "VNIOS"
  upgrade_group     = "Grid Master"
  enable_ha         = true
  router_id         = 100
  ha_on_cloud       = true
  ha_cloud_platform = "AWS"
  vip_setting = {
    address     = "172.28.83.245"
    gateway     = "172.28.82.1"
    subnet_mask = "255.255.254.0"
  }

  node_info = [
    {
      lan_ha_port_setting = {
        ha_ip_address      = "172.28.82.239"
        mgmt_lan           = "172.28.82.31"
        ha_cloud_attribute = "<eni-0a0a9a4>"
      }
    },
    {
      lan_ha_port_setting = {
        ha_ip_address      = "172.28.82.233"
        mgmt_lan           = "172.28.83.237"
        ha_cloud_attribute = "<eni-014efdde>"
      }
    }
  ]

  // To configure grid level dns resolver settings, use the grid_level_dns_resolver_setting attribute 
  grid_level_dns_resolver_setting = {
    resolvers = [
      "10.10.10.10"
  ] }
}

// Initiate the grid join
resource "nios_grid_join" "member_join_4" {
  member_url      = "https://172.28.83.237"
  member_username = "username"
  member_password = "password"
  grid_name       = "Infoblox"
  master          = "172.28.83.245"
  shared_secret   = "secret"
  depends_on      = [nios_grid_member.example_grid_member_4]
}
