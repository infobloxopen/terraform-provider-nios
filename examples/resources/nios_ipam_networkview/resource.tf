// Create Network View with Basic Fields
resource "nios_ipam_networkview" "create_network_view" {
  name = "my_network_view"
}

// Create Network View with additional fields
resource "nios_ipam_networkview" "create_with_additional_fields" {
  name    = "test-network-view"
  comment = "Created Network View"

  cloud_info = {
    delegated_member = {
      ipv4addr = "172.172.172.172"
      ipv6addr = "2001::123"
      name     = "infoblox.cloudmem"
    }
  }

  ddns_zone_primaries = [
    {
      dns_grid_primary = "infoblox.cloudmem"
      zone_match       = "GRID"
      dns_grid_zone = {
        ref = "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5maXJzdA:first.com/default"
      }
    },
    {
      dns_grid_primary = "infoblox.cloudmem"
      zone_match       = "GRID"
      dns_grid_zone = {
        ref = "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5zZWNvbmQ:second.com/default"
      }
    }
  ]

  remote_reverse_zones = [
    {
      fqdn           = "0.168.192.in-addr.arpa"
      key_type       = "NONE"
      server_address = "192.168.12.12"
    },
    {
      fqdn           = "1.168.192.in-addr.arpa"
      key_type       = "TSIG"
      server_address = "192.168.12.12"
      tsig_key_name  = "aeiou"
      tsig_key_alg   = "HMAC-SHA256"
      tsig_key       = "dGhpc2lzdGVzdHRzaWdrZXk="
    }
  ]

  remote_forward_zones = [
    {
      fqdn           = "fwdzone1.com"
      key_type       = "NONE"
      server_address = "192.168.12.12"
    },
    {
      fqdn           = "fwdzone2.com"
      key_type       = "TSIG"
      server_address = "192.168.12.12"
      tsig_key_name  = "aeiou"
      tsig_key_alg   = "HMAC-SHA256"
      tsig_key       = "dGhpc2lzdGVzdHRzaWdrZXk="
    }
  ]

  federated_realms = [
    {
      id   = "123"
      name = "federated_realm1"
    }
  ]

  extattrs = {
    Site = "Mars"
  }

  mgm_private = false

  # ----------------------------------------------------------------------------
  # The following fields are only applicable during update operations
  # ----------------------------------------------------------------------------
  /*
  ddns_dns_view = "default.test-network-view"

  internal_forward_zones = [
    "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5maXJzdA:first.com/default.test-network-view",
    "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5zZWNvbmQ:second.com/default.test-network-view"
  ]
  */
}