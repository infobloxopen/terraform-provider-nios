// Create DNS View with Basic Fields
resource "nios_dns_view" "create_view" {
  name = "example_view"
}

// Create DNS View with additional fields
resource "nios_dns_view" "create_with_additional_fields" {
  name         = "example_custom_view"
  comment      = "DNS View"
  network_view = "default"

  //Blacklist settings
  use_blacklist                = true
  blacklist_action             = "REDIRECT"
  blacklist_log_query          = true
  blacklist_redirect_addresses = ["10.2.3.2", "10.23.2.2"]
  blacklist_redirect_ttl       = 34
  blacklist_rulesets           = ["ruleset1", "ruleset2"]
  enable_blacklist             = true

  //root name server settings
  root_name_server_type = "CUSTOM"
  custom_root_name_servers = [
    {
      address = "10.0.0.2"
      name    = "external-server-1"
    }
  ]
  use_root_name_server = true

  //DDNS settings
  ddns_force_creation_timestamp_update     = true
  use_ddns_force_creation_timestamp_update = true
  ddns_principal_group                     = "dynamic_update_grp_1"
  use_ddns_principal_security              = true
  ddns_restrict_patterns                   = true
  use_ddns_patterns_restriction            = true
  ddns_restrict_patterns_list              = ["pattern2.example.com", "pattern3.example.com"]
  ddns_restrict_protected                  = true
  use_ddns_restrict_protected              = true
  ddns_restrict_secure                     = true
  ddns_restrict_static                     = true
  use_ddns_restrict_static                 = true

  //DNS64 settings
  dns64_enabled = true
  use_dns64     = true
  dns64_groups  = ["dns64_group"]

  //DNSSEC settings
  dnssec_enabled                    = true
  use_dnssec                        = true
  dnssec_expired_signatures_enabled = false
  dnssec_negative_trust_anchors     = ["examplezone2.com", "examplezone3.com"]
  dnssec_trusted_keys = [
    {
      algorithm             = "14"
      dnssec_must_be_secure = true
      fqdn                  = "test2.com"
      key                   = "dsfdfdfdfdfdfdfdfdfdfdfdfdfdfdfdfdfdfd"
      secure_entry_point    = true
    }
  ]
  dnssec_validation_enabled = true

  //udp size configuration
  edns_udp_size     = 1234
  use_edns_udp_size = true
  max_udp_size      = 1233
  use_max_udp_size  = true

  //filter AAAA 
  filter_aaaa     = "BREAK_DNSSEC"
  use_filter_aaaa = true
  filter_aaaa_list = [
    {
      address    = "10.0.0.12"
      permission = "ALLOW"
    }
  ]

  //forwarders settings 
  use_forwarders = true
  forwarders     = ["10.192.81.23"]
  forward_only   = true

  //last queried ACL
  last_queried_acl = [
    {
      address    = "10.0.0.23"
      permission = "ALLOW"
    }
  ]

  //match client and destinations
  match_destinations = [
    {
      struct     = "addressac"
      address    = "192.168.0.45"
      permission = "ALLOW"
    },
  ]
  match_clients = [
    {
      struct     = "addressac"
      address    = "92.168.0.23"
      permission = "ALLOW"
    },
  ]

  //cache and ncache ttl settings 
  max_cache_ttl      = 3454
  use_max_cache_ttl  = true
  max_ncache_ttl     = 3600
  use_max_ncache_ttl = true

  //notify delay settings 
  notify_delay = 6

  //NXDOMAIN redirect settings
  nxdomain_log_query             = true
  use_nxdomain_redirect          = true
  nxdomain_redirect              = true
  nxdomain_redirect_addresses    = ["12.0.0.0"]
  nxdomain_redirect_addresses_v6 = ["2001:db8::1"]
  nxdomain_redirect_ttl          = 23
  nxdomain_rulesets              = ["nxdomain_ruleset"]

  //recursion settings 
  use_recursion = true
  recursion     = true

  //response rate limiting settings
  use_response_rate_limiting = true
  response_rate_limiting = {
    enable_rrl           = true
    log_only             = true
    responses_per_second = 200
    slip                 = 3
    window               = 30
  }

  //Rpz settings
  rpz_drop_ip_rule_enabled                = true
  use_rpz_drop_ip_rule                    = true
  rpz_drop_ip_rule_min_prefix_length_ipv4 = 24
  rpz_drop_ip_rule_min_prefix_length_ipv6 = 48
  rpz_qname_wait_recurse                  = true
  use_rpz_qname_wait_recurse              = true

  //scavenging settings
  use_scavenging_settings = true
  scavenging_settings = {
    enable_scavenging = true
    expression_list = [
      {
        op       = "AND",
        op1_type = "LIST",
      },
      {
        op       = "EQ",
        op1      = "rtype",
        op1_type = "FIELD",
        op2      = "AAAA",
        op2_type = "STRING",
      },
      {
        op = "ENDLIST",
      }
    ]
  }

  //sortlist settings
  sortlist = [
    {
      address = "10.0.0.0"
    }
  ]

  extattrs = {
    Site = "location-1"
  }

}
