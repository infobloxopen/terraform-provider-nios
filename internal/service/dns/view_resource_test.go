package dns_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// TODO : Objects to be present in the grid for testing
//DDNS Principal Group = "dynamic_update_grp_1"
//NX Domain Ruleset = "nxdomain_ruleset_1"
//Blacklist Ruleset = "ruleset1, ruleset2, ruleset3, ruleset4"
//network view = "custom "

var readableAttributesForView = "blacklist_action,blacklist_log_query,blacklist_redirect_addresses,blacklist_redirect_ttl,blacklist_rulesets,cloud_info,comment,custom_root_name_servers,ddns_force_creation_timestamp_update,ddns_principal_group,ddns_principal_tracking,ddns_restrict_patterns,ddns_restrict_patterns_list,ddns_restrict_protected,ddns_restrict_secure,ddns_restrict_static,disable,dns64_enabled,dns64_groups,dnssec_enabled,dnssec_expired_signatures_enabled,dnssec_negative_trust_anchors,dnssec_trusted_keys,dnssec_validation_enabled,edns_udp_size,enable_blacklist,enable_fixed_rrset_order_fqdns,enable_match_recursive_only,extattrs,filter_aaaa,filter_aaaa_list,fixed_rrset_order_fqdns,forward_only,forwarders,is_default,last_queried_acl,match_clients,match_destinations,max_cache_ttl,max_ncache_ttl,max_udp_size,name,network_view,notify_delay,nxdomain_log_query,nxdomain_redirect,nxdomain_redirect_addresses,nxdomain_redirect_addresses_v6,nxdomain_redirect_ttl,nxdomain_rulesets,recursion,response_rate_limiting,root_name_server_type,rpz_drop_ip_rule_enabled,rpz_drop_ip_rule_min_prefix_length_ipv4,rpz_drop_ip_rule_min_prefix_length_ipv6,rpz_qname_wait_recurse,scavenging_settings,sortlist,use_blacklist,use_ddns_force_creation_timestamp_update,use_ddns_patterns_restriction,use_ddns_principal_security,use_ddns_restrict_protected,use_ddns_restrict_static,use_dns64,use_dnssec,use_edns_udp_size,use_filter_aaaa,use_fixed_rrset_order_fqdns,use_forwarders,use_max_cache_ttl,use_max_ncache_ttl,use_max_udp_size,use_nxdomain_redirect,use_recursion,use_response_rate_limiting,use_root_name_server,use_rpz_drop_ip_rule,use_rpz_qname_wait_recurse,use_scavenging_settings,use_sortlist"

func TestAccViewResource_basic(t *testing.T) {
	var resourceName = "nios_dns_view.test"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "blacklist_action", "REDIRECT"),
					resource.TestCheckResourceAttr(resourceName, "blacklist_log_query", "false"),
					resource.TestCheckResourceAttr(resourceName, "blacklist_redirect_ttl", "60"),
					resource.TestCheckResourceAttr(resourceName, "ddns_force_creation_timestamp_update", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_tracking", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_protected", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_secure", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_static", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "dns64_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_expired_signatures_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_validation_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "edns_udp_size", "1220"),
					resource.TestCheckResourceAttr(resourceName, "enable_blacklist", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_fixed_rrset_order_fqdns", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_match_recursive_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa", "NO"),
					resource.TestCheckResourceAttr(resourceName, "forward_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "max_cache_ttl", "604800"),
					resource.TestCheckResourceAttr(resourceName, "max_ncache_ttl", "10800"),
					resource.TestCheckResourceAttr(resourceName, "max_udp_size", "1220"),
					resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
					resource.TestCheckResourceAttr(resourceName, "notify_delay", "5"),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_log_query", "false"),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect", "false"),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect_ttl", "60"),
					resource.TestCheckResourceAttr(resourceName, "recursion", "false"),
					resource.TestCheckResourceAttr(resourceName, "response_rate_limiting.enable_rrl", "false"),
					resource.TestCheckResourceAttr(resourceName, "response_rate_limiting.log_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "response_rate_limiting.responses_per_second", "100"),
					resource.TestCheckResourceAttr(resourceName, "response_rate_limiting.window", "15"),
					resource.TestCheckResourceAttr(resourceName, "response_rate_limiting.slip", "2"),
					resource.TestCheckResourceAttr(resourceName, "root_name_server_type", "INTERNET"),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv4", "29"),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv6", "112"),
					resource.TestCheckResourceAttr(resourceName, "rpz_qname_wait_recurse", "false"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.enable_auto_reclamation", "false"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.enable_recurrent_scavenging", "false"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.enable_rr_last_queried", "false"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.enable_scavenging", "false"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.enable_zone_last_queried", "false"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.reclaim_associated_records", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_blacklist", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_force_creation_timestamp_update", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_patterns_restriction", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_principal_security", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_protected", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_static", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_dns64", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_dnssec", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_edns_udp_size", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_filter_aaaa", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_fixed_rrset_order_fqdns", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_forwarders", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_max_cache_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_max_ncache_ttl", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_max_udp_size", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_nxdomain_redirect", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_recursion", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_response_rate_limiting", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_root_name_server", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_rpz_drop_ip_rule", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_rpz_qname_wait_recurse", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_scavenging_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_sortlist", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_disappears(t *testing.T) {
	resourceName := "nios_dns_view.test"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckViewDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccViewBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					testAccCheckViewDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccViewResource_BlacklistAction(t *testing.T) {
	var resourceName = "nios_dns_view.test_blacklist_action"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	blacklistAction := "REFUSE"
	blacklistActionUpdate := "REDIRECT"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewBlacklistAction(name, blacklistAction),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_action", blacklistAction),
				),
			},
			// Update and Read
			{
				Config: testAccViewBlacklistAction(name, blacklistActionUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_action", blacklistActionUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_BlacklistLogQuery(t *testing.T) {
	var resourceName = "nios_dns_view.test_blacklist_log_query"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	blacklistLogQuery := true
	blacklistLogQueryUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewBlacklistLogQuery(name, blacklistLogQuery),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_log_query", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewBlacklistLogQuery(name, blacklistLogQueryUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_log_query", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_BlacklistRedirectAddresses(t *testing.T) {
	var resourceName = "nios_dns_view.test_blacklist_redirect_addresses"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	blacklistRedirectAddresses := []string{"10.0.0.1", "10.0.0.29"}
	blacklistRedirectAddressesUpdate := []string{"10.0.0.23", "10.0.0.54"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewBlacklistRedirectAddresses(name, blacklistRedirectAddresses),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_redirect_addresses.0", "10.0.0.1"),
					resource.TestCheckResourceAttr(resourceName, "blacklist_redirect_addresses.1", "10.0.0.29"),
				),
			},
			// Update and Read
			{
				Config: testAccViewBlacklistRedirectAddresses(name, blacklistRedirectAddressesUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_redirect_addresses.0", "10.0.0.23"),
					resource.TestCheckResourceAttr(resourceName, "blacklist_redirect_addresses.1", "10.0.0.54"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_BlacklistRedirectTtl(t *testing.T) {
	var resourceName = "nios_dns_view.test_blacklist_redirect_ttl"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	blacklistRedirctTtl := 75
	blacklistRedirctTtlUpdate := 90

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewBlacklistRedirectTtl(name, blacklistRedirctTtl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_redirect_ttl", "75"),
				),
			},
			// Update and Read
			{
				Config: testAccViewBlacklistRedirectTtl(name, blacklistRedirctTtlUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_redirect_ttl", "90"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_BlacklistRulesets(t *testing.T) {
	var resourceName = "nios_dns_view.test_blacklist_rulesets"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	blacklistRulesets := []string{"ruleset1", "ruleset2"}
	blacklistRulesetsUpdate := []string{"ruleset3", "ruleset4"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewBlacklistRulesets(name, blacklistRulesets),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_rulesets.0", "ruleset1"),
					resource.TestCheckResourceAttr(resourceName, "blacklist_rulesets.1", "ruleset2"),
				),
			},
			// Update and Read
			{
				Config: testAccViewBlacklistRulesets(name, blacklistRulesetsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_rulesets.0", "ruleset3"),
					resource.TestCheckResourceAttr(resourceName, "blacklist_rulesets.1", "ruleset4"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_view.test_comment"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	comment := "dns view comment"
	commentUpdate := "updated dns view comment"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewComment(name, comment),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", comment),
				),
			},
			// Update and Read
			{
				Config: testAccViewComment(name, commentUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", commentUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_CustomRootNameServers(t *testing.T) {
	var resourceName = "nios_dns_view.test_custom_root_name_servers"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	customRootNameServers := []map[string]any{
		{
			"address": "10.0.0.2",
			"name":    "external-server-1",
		},
	}
	customRootNameServersUpdate := []map[string]any{
		{
			"address": "10.0.0.23",
			"name":    "external-server-2",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewCustomRootNameServers(name, customRootNameServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "custom_root_name_servers.0.address", "10.0.0.2"),
					resource.TestCheckResourceAttr(resourceName, "custom_root_name_servers.0.name", "external-server-1"),
				),
			},
			// Update and Read
			{
				Config: testAccViewCustomRootNameServers(name, customRootNameServersUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "custom_root_name_servers.0.address", "10.0.0.23"),
					resource.TestCheckResourceAttr(resourceName, "custom_root_name_servers.0.name", "external-server-2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DdnsForceCreationTimestampUpdate(t *testing.T) {
	var resourceName = "nios_dns_view.test_ddns_force_creation_timestamp_update"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	ddnsForceCreationTimestampUpdate := true
	ddnsForceCreationTimestampUpdateUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDdnsForceCreationTimestampUpdate(name, ddnsForceCreationTimestampUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_force_creation_timestamp_update", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDdnsForceCreationTimestampUpdate(name, ddnsForceCreationTimestampUpdateUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_force_creation_timestamp_update", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DdnsPrincipalGroup(t *testing.T) {
	var resourceName = "nios_dns_view.test_ddns_principal_group"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	ddnsPrincipalGroup := "dynamic_update_grp_1"
	ddnsPrincipalGroupUpdate := "dynamic_update_grp_2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDdnsPrincipalGroup(name, ddnsPrincipalGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_group", "dynamic_update_grp_1"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDdnsPrincipalGroup(name, ddnsPrincipalGroupUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_group", "dynamic_update_grp_2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DdnsPrincipalTracking(t *testing.T) {
	var resourceName = "nios_dns_view.test_ddns_principal_tracking"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	ddnsPrincipalTracking := true
	ddnsPrincipalTrackingUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDdnsPrincipalTracking(name, ddnsPrincipalTracking),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_tracking", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDdnsPrincipalTracking(name, ddnsPrincipalTrackingUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_tracking", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DdnsRestrictPatterns(t *testing.T) {
	var resourceName = "nios_dns_view.test_ddns_restrict_patterns"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	ddnsRestrictPatterns := true
	ddnsRestrictPatternsUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDdnsRestrictPatterns(name, ddnsRestrictPatterns),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDdnsRestrictPatterns(name, ddnsRestrictPatternsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DdnsRestrictPatternsList(t *testing.T) {
	var resourceName = "nios_dns_view.test_ddns_restrict_patterns_list"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	ddnsRestrictPatternList := []string{"pattern1.example.com"}
	ddnsRestrictPatternListUpdate := []string{"pattern2.example.com", "pattern3.example.com"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDdnsRestrictPatternsList(name, ddnsRestrictPatternList),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns_list.0", "pattern1.example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDdnsRestrictPatternsList(name, ddnsRestrictPatternListUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns_list.0", "pattern2.example.com"),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns_list.1", "pattern3.example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DdnsRestrictProtected(t *testing.T) {
	var resourceName = "nios_dns_view.test_ddns_restrict_protected"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	ddnsRestrictProtected := true
	ddnsRestrictProtectedUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDdnsRestrictProtected(name, ddnsRestrictProtected),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_protected", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDdnsRestrictProtected(name, ddnsRestrictProtectedUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_protected", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DdnsRestrictSecure(t *testing.T) {
	var resourceName = "nios_dns_view.test_ddns_restrict_secure"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	ddnsRestrictSecure := true
	ddnsRestrictSecureUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDdnsRestrictSecure(name, ddnsRestrictSecure),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_secure", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDdnsRestrictSecure(name, ddnsRestrictSecureUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_secure", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DdnsRestrictStatic(t *testing.T) {
	var resourceName = "nios_dns_view.test_ddns_restrict_static"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	ddnsRestrictStatic := true
	ddnsRestrictStaticUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDdnsRestrictStatic(name, ddnsRestrictStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_static", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDdnsRestrictStatic(name, ddnsRestrictStaticUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_static", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_view.test_disable"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	disable := true
	disableUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDisable(name, disable),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDisable(name, disableUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Dns64Enabled(t *testing.T) {
	var resourceName = "nios_dns_view.test_dns64_enabled"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	dns64Enabled := true
	dns64EnabledUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDns64Enabled(name, dns64Enabled),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns64_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDns64Enabled(name, dns64EnabledUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns64_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Dns64Groups(t *testing.T) {
	var resourceName = "nios_dns_view.test_dns64_groups"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	dns64Groups := []string{"dns64_group"}
	dns64GroupsUpdate := []string{"dns64_group_2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDns64Groups(name, dns64Groups),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns64_groups.0", "dns64_group"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDns64Groups(name, dns64GroupsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns64_groups.0", "dns64_group_2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DnssecEnabled(t *testing.T) {
	var resourceName = "nios_dns_view.test_dnssec_enabled"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	dnssecEnabled := true
	dnssecEnabledUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDnssecEnabled(name, dnssecEnabled),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDnssecEnabled(name, dnssecEnabledUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DnssecExpiredSignaturesEnabled(t *testing.T) {
	var resourceName = "nios_dns_view.test_dnssec_expired_signatures_enabled"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	dnssecExpiredSignaturesEnabled := true
	dnssecExpiredSignaturesEnabledUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDnssecExpiredSignaturesEnabled(name, dnssecExpiredSignaturesEnabled),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_expired_signatures_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDnssecExpiredSignaturesEnabled(name, dnssecExpiredSignaturesEnabledUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_expired_signatures_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DnssecNegativeTrustAnchors(t *testing.T) {
	var resourceName = "nios_dns_view.test_dnssec_negative_trust_anchors"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	dnssecNegativeTrustAnchors := []string{"examplezone.com"}
	dnssecNegativeTrustAnchorsUpdate := []string{"examplezone2.com", "examplezone3.com"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDnssecNegativeTrustAnchors(name, dnssecNegativeTrustAnchors),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_negative_trust_anchors.0", "examplezone.com"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDnssecNegativeTrustAnchors(name, dnssecNegativeTrustAnchorsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_negative_trust_anchors.0", "examplezone2.com"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_negative_trust_anchors.1", "examplezone3.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DnssecTrustedKeys(t *testing.T) {
	var resourceName = "nios_dns_view.test_dnssec_trusted_keys"
	var v dns.View
	key := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAweW4MAnsKGjk4dt6a42CrIA/BV9YEKThzXZVlBSdUfn0D2YDOMkWvlMxUPVd5iEc2DXulrpBSNbxL1y7Ude11fs1+cOgvcgmQX1Yvu9e14OzeMfk3ZJB8Ldnmb5xrNR9y4ASqh771PZA6xK3qVS+k7YLGp3xnRrd1+zMLcUMI5J+8ZBOIn/6K37DkirKhBv5hKfttTNQbPiwDXwS/vduUv0vUN/xLUKg6099abOn05nefWg+BoxuMySVtqhB6pgW+1BrGrSISOTZDTKojguftya3vqFhb5m/G3F39BdIAlNWP/P2lP8ksuER/pczE6muS8CS2ArCbaN+Z7iddg5P6wIDAQAB"
	name := acctest.RandomNameWithPrefix("view")
	dnssecTrustedKeys := []map[string]any{
		{
			"algorithm":             "14",
			"dnssec_must_be_secure": false,
			"fqdn":                  "test.com",
			"key":                   key,
			"secure_entry_point":    true,
		},
	}
	dnssecTrustedKeysUpdate := []map[string]any{
		{
			"algorithm":             "14",
			"dnssec_must_be_secure": false,
			"fqdn":                  "test2.com",
			"key":                   key,
			"secure_entry_point":    true,
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDnssecTrustedKeys(name, dnssecTrustedKeys),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trusted_keys.0.algorithm", "14"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trusted_keys.0.dnssec_must_be_secure", "false"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trusted_keys.0.fqdn", "test.com"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trusted_keys.0.key", key),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trusted_keys.0.secure_entry_point", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDnssecTrustedKeys(name, dnssecTrustedKeysUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trusted_keys.0.algorithm", "14"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trusted_keys.0.dnssec_must_be_secure", "false"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trusted_keys.0.fqdn", "test2.com"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trusted_keys.0.key", key),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trusted_keys.0.secure_entry_point", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DnssecValidationEnabled(t *testing.T) {
	var resourceName = "nios_dns_view.test_dnssec_validation_enabled"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	dnssecValidationEnabled := false
	dnssecValidationEnabledUpdate := true

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDnssecValidationEnabled(name, dnssecValidationEnabled),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_validation_enabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDnssecValidationEnabled(name, dnssecValidationEnabledUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_validation_enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_EdnsUdpSize(t *testing.T) {
	var resourceName = "nios_dns_view.test_edns_udp_size"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	ednsUdpSize := 1232
	ednsUdpSizeUpdate := 4096

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewEdnsUdpSize(name, ednsUdpSize),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "edns_udp_size", "1232"),
				),
			},
			// Update and Read
			{
				Config: testAccViewEdnsUdpSize(name, ednsUdpSizeUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "edns_udp_size", "4096"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_EnableBlacklist(t *testing.T) {
	var resourceName = "nios_dns_view.test_enable_blacklist"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	enableBlacklist := true
	enableBlacklistUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewEnableBlacklist(name, enableBlacklist),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_blacklist", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewEnableBlacklist(name, enableBlacklistUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_blacklist", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_EnableFixedRrsetOrderFqdns(t *testing.T) {
	var resourceName = "nios_dns_view.test_enable_fixed_rrset_order_fqdns"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	enableFixedRrsetOrderFqdns := true
	enableFixedRrsetOrderFqdnsUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewEnableFixedRrsetOrderFqdns(name, enableFixedRrsetOrderFqdns),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_fixed_rrset_order_fqdns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewEnableFixedRrsetOrderFqdns(name, enableFixedRrsetOrderFqdnsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_fixed_rrset_order_fqdns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_EnableMatchRecursiveOnly(t *testing.T) {
	var resourceName = "nios_dns_view.test_enable_match_recursive_only"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	enableMatchRecursiveOnly := true
	enableMatchRecursiveOnlyUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewEnableMatchRecursiveOnly(name, enableMatchRecursiveOnly),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_match_recursive_only", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewEnableMatchRecursiveOnly(name, enableMatchRecursiveOnlyUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_match_recursive_only", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_view.test_extattrs"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewExtAttrs(name, map[string]any{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccViewExtAttrs(name, map[string]any{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_FilterAaaa(t *testing.T) {
	var resourceName = "nios_dns_view.test_filter_aaaa"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	filterAaaa := "BREAK_DNSSEC"
	filterAaaaUpdate := "YES"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewFilterAaaa(name, filterAaaa),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa", "BREAK_DNSSEC"),
				),
			},
			// Update and Read
			{
				Config: testAccViewFilterAaaa(name, filterAaaaUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa", "YES"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_FilterAaaaList(t *testing.T) {
	var resourceName = "nios_dns_view.test_filter_aaaa_list"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	filterAaaaList := []map[string]any{
		{
			"address":    "10.0.0.23",
			"permission": "DENY",
		},
	}
	filterAaaaListUpdate := []map[string]any{
		{
			"address":    "10.0.0.12",
			"permission": "ALLOW",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewFilterAaaaList(name, filterAaaaList),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_list.0.address", "10.0.0.23"),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_list.0.permission", "DENY"),
				),
			},
			// Update and Read
			{
				Config: testAccViewFilterAaaaList(name, filterAaaaListUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_list.0.address", "10.0.0.12"),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_list.0.permission", "ALLOW"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_FixedRrsetOrderFqdns(t *testing.T) {
	var resourceName = "nios_dns_view.test_fixed_rrset_order_fqdns"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	fixedRrsetOrderFqdns := []map[string]any{
		{
			"fqdn":        "example.com",
			"record_type": "AAAA",
		},
	}
	fixedRrsetOrderFqdnsUpdate := []map[string]any{
		{
			"fqdn":        "example.org",
			"record_type": "BOTH",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewFixedRrsetOrderFqdns(name, fixedRrsetOrderFqdns),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fixed_rrset_order_fqdns.0.fqdn", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "fixed_rrset_order_fqdns.0.record_type", "AAAA"),
				),
			},
			// Update and Read
			{
				Config: testAccViewFixedRrsetOrderFqdns(name, fixedRrsetOrderFqdnsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fixed_rrset_order_fqdns.0.fqdn", "example.org"),
					resource.TestCheckResourceAttr(resourceName, "fixed_rrset_order_fqdns.0.record_type", "BOTH"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_ForwardOnly(t *testing.T) {
	var resourceName = "nios_dns_view.test_forward_only"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	forwardOnly := true
	forwardOnlyUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewForwardOnly(name, forwardOnly),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forward_only", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewForwardOnly(name, forwardOnlyUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forward_only", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Forwarders(t *testing.T) {
	var resourceName = "nios_dns_view.test_forwarders"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	forwarders := []string{"10.123.86.42"}
	forwardersUpdate := []string{"10.252.23.44"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewForwarders(name, forwarders),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders.0", "10.123.86.42"),
				),
			},
			// Update and Read
			{
				Config: testAccViewForwarders(name, forwardersUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders.0", "10.252.23.44"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_LastQueriedAcl(t *testing.T) {
	var resourceName = "nios_dns_view.test_last_queried_acl"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	lastQueriedAcl := []map[string]any{
		{
			"address":    "10.0.0.23",
			"permission": "DENY",
		},
	}
	lastQueriedAclUpdate := []map[string]any{
		{
			"address":    "10.0.0.12",
			"permission": "ALLOW",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewLastQueriedAcl(name, lastQueriedAcl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "last_queried_acl.0.address", "10.0.0.23"),
					resource.TestCheckResourceAttr(resourceName, "last_queried_acl.0.permission", "DENY"),
				),
			},
			// Update and Read
			{
				Config: testAccViewLastQueriedAcl(name, lastQueriedAclUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "last_queried_acl.0.address", "10.0.0.12"),
					resource.TestCheckResourceAttr(resourceName, "last_queried_acl.0.permission", "ALLOW"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_MatchClients(t *testing.T) {
	var resourceName = "nios_dns_view.test_match_clients"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	matchClients := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "10.0.0.0",
			"permission": "ALLOW",
		},
	}
	matchClientsUpdate := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "192.168.0.0",
			"permission": "ALLOW",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewMatchClients(name, matchClients),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_clients.0.address", "10.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "match_clients.0.permission", "ALLOW"),
				),
			},
			// Update and Read
			{
				Config: testAccViewMatchClients(name, matchClientsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_clients.0.address", "192.168.0.0"),
					resource.TestCheckResourceAttr(resourceName, "match_clients.0.permission", "ALLOW"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_MatchDestinations(t *testing.T) {
	var resourceName = "nios_dns_view.test_match_destinations"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	matchDestinations := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "10.0.0.0",
			"permission": "ALLOW",
		},
	}
	matchDestinationsUpdate := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "192.168.0.0",
			"permission": "ALLOW",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewMatchDestinations(name, matchDestinations),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_destinations.0.address", "10.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "match_destinations.0.permission", "ALLOW"),
				),
			},
			// Update and Read
			{
				Config: testAccViewMatchDestinations(name, matchDestinationsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_destinations.0.address", "192.168.0.0"),
					resource.TestCheckResourceAttr(resourceName, "match_destinations.0.permission", "ALLOW"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_MaxCacheTtl(t *testing.T) {
	var resourceName = "nios_dns_view.test_max_cache_ttl"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	maxCacheTtl := 3600
	maxCacheTtlUpdate := 7200

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewMaxCacheTtl(name, maxCacheTtl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_cache_ttl", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccViewMaxCacheTtl(name, maxCacheTtlUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_cache_ttl", "7200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_MaxNcacheTtl(t *testing.T) {
	var resourceName = "nios_dns_view.test_max_ncache_ttl"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	maxNcacheTtl := 300
	maxNcacheTtlUpdate := 600

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewMaxNcacheTtl(name, maxNcacheTtl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_ncache_ttl", "300"),
				),
			},
			// Update and Read
			{
				Config: testAccViewMaxNcacheTtl(name, maxNcacheTtlUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_ncache_ttl", "600"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_MaxUdpSize(t *testing.T) {
	var resourceName = "nios_dns_view.test_max_udp_size"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	maxUdpSize := 512
	maxUdpSizeUpdate := 1024

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewMaxUdpSize(name, maxUdpSize),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_udp_size", "512"),
				),
			},
			// Update and Read
			{
				Config: testAccViewMaxUdpSize(name, maxUdpSizeUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_udp_size", "1024"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Name(t *testing.T) {
	var resourceName = "nios_dns_view.test_name"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	nameUpdate := acctest.RandomNameWithPrefix("view_update")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccViewName(nameUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_NetworkView(t *testing.T) {
	var resourceName = "nios_dns_view.test_network_view"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	networkView := "default"
	networkViewUpdate := "custom "

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewNetworkView(name, networkView),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccViewNetworkView(name, networkViewUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "custom "),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_NotifyDelay(t *testing.T) {
	var resourceName = "nios_dns_view.test_notify_delay"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	notifyDelay := 78
	notifyDelayUpdate := 10

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewNotifyDelay(name, notifyDelay),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notify_delay", "78"),
				),
			},
			// Update and Read
			{
				Config: testAccViewNotifyDelay(name, notifyDelayUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notify_delay", "10"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_NxdomainLogQuery(t *testing.T) {
	var resourceName = "nios_dns_view.test_nxdomain_log_query"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	nxdomainLogQuery := true
	nxdomainLogQueryUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewNxdomainLogQuery(name, nxdomainLogQuery),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_log_query", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewNxdomainLogQuery(name, nxdomainLogQueryUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_log_query", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_NxdomainRedirect(t *testing.T) {
	var resourceName = "nios_dns_view.test_nxdomain_redirect"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	nxdomainRedirect := true
	nxdomainRedirectUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewNxdomainRedirect(name, nxdomainRedirect),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewNxdomainRedirect(name, nxdomainRedirectUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_NxdomainRedirectAddresses(t *testing.T) {
	var resourceName = "nios_dns_view.test_nxdomain_redirect_addresses"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	nxDomainRedirectAddress := []string{"10.87.9.7"}
	nxDomainRedirectAddressUpdate := []string{"10.3.23.56", "5.4.3.5"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewNxdomainRedirectAddresses(name, nxDomainRedirectAddress),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect_addresses.0", "10.87.9.7"),
				),
			},
			// Update and Read
			{
				Config: testAccViewNxdomainRedirectAddresses(name, nxDomainRedirectAddressUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect_addresses.0", "10.3.23.56"),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect_addresses.1", "5.4.3.5"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_NxdomainRedirectAddressesV6(t *testing.T) {
	var resourceName = "nios_dns_view.test_nxdomain_redirect_addresses_v6"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	nxdomainRedirectAddressesV6 := []string{"2001:db8::1", "2001:db8::2"}
	nxdomainRedirectAddressesV6Update := []string{"2001:db8::3", "2001:db8::4"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewNxdomainRedirectAddressesV6(name, nxdomainRedirectAddressesV6),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect_addresses_v6.0", "2001:db8::1"),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect_addresses_v6.1", "2001:db8::2"),
				),
			},
			// Update and Read
			{
				Config: testAccViewNxdomainRedirectAddressesV6(name, nxdomainRedirectAddressesV6Update),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect_addresses_v6.0", "2001:db8::3"),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect_addresses_v6.1", "2001:db8::4"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_NxdomainRedirectTtl(t *testing.T) {
	var resourceName = "nios_dns_view.test_nxdomain_redirect_ttl"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	nxdomainRedirectTtl := 3600
	nxdomainRedirectTtlUpdate := 7200

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewNxdomainRedirectTtl(name, nxdomainRedirectTtl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect_ttl", "3600"),
				),
			},
			// Update and Read
			{
				Config: testAccViewNxdomainRedirectTtl(name, nxdomainRedirectTtlUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect_ttl", "7200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_NxdomainRulesets(t *testing.T) {
	var resourceName = "nios_dns_view.test_nxdomain_rulesets"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	nxdomainRuleset := []string{"nxdomain_ruleset"}
	nxdomainRulesetUpdate := []string{"nxdomain_ruleset2"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewNxdomainRulesets(name, nxdomainRuleset),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_rulesets.0", "nxdomain_ruleset"),
				),
			},
			// Update and Read
			{
				Config: testAccViewNxdomainRulesets(name, nxdomainRulesetUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_rulesets.0", "nxdomain_ruleset2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Recursion(t *testing.T) {
	var resourceName = "nios_dns_view.test_recursion"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	recursion := true
	recursionUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewRecursion(name, recursion),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recursion", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewRecursion(name, recursionUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recursion", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_ResponseRateLimiting(t *testing.T) {
	var resourceName = "nios_dns_view.test_response_rate_limiting"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	responseRateLimitingMap := map[string]interface{}{
		"enable_rrl":           false,
		"log_only":             false,
		"responses_per_second": 100,
		"slip":                 2,
		"window":               15,
	}
	responseRateLimitingMapUpdate := map[string]interface{}{
		"enable_rrl":           true,
		"log_only":             true,
		"responses_per_second": 200,
		"slip":                 3,
		"window":               30,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewResponseRateLimiting(name, responseRateLimitingMap),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "response_rate_limiting.enable_rrl", "false"),
					resource.TestCheckResourceAttr(resourceName, "response_rate_limiting.log_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "response_rate_limiting.responses_per_second", "100"),
					resource.TestCheckResourceAttr(resourceName, "response_rate_limiting.slip", "2"),
					resource.TestCheckResourceAttr(resourceName, "response_rate_limiting.window", "15"),
				),
			},
			// Update and Read
			{
				Config: testAccViewResponseRateLimiting(name, responseRateLimitingMapUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "response_rate_limiting.enable_rrl", "true"),
					resource.TestCheckResourceAttr(resourceName, "response_rate_limiting.log_only", "true"),
					resource.TestCheckResourceAttr(resourceName, "response_rate_limiting.responses_per_second", "200"),
					resource.TestCheckResourceAttr(resourceName, "response_rate_limiting.slip", "3"),
					resource.TestCheckResourceAttr(resourceName, "response_rate_limiting.window", "30"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_RootNameServerType(t *testing.T) {
	var resourceName = "nios_dns_view.test_root_name_server_type"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	rootNameServerType := "CUSTOM"
	rootNameServerTypeUpdate := "INTERNET"
	customRootNameServers := []map[string]any{
		{
			"address": "10.0.0.2",
			"name":    "external-server-1",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewRootNameServerType(name, rootNameServerType, customRootNameServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "root_name_server_type", rootNameServerType),
				),
			},
			// Update and Read
			{
				Config: testAccViewRootNameServerType(name, rootNameServerTypeUpdate, customRootNameServers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "root_name_server_type", rootNameServerTypeUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_RpzDropIpRuleEnabled(t *testing.T) {
	var resourceName = "nios_dns_view.test_rpz_drop_ip_rule_enabled"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	rpzDropIpRuleEnabled := true
	rpzDropIpRuleEnabledUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewRpzDropIpRuleEnabled(name, rpzDropIpRuleEnabled),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewRpzDropIpRuleEnabled(name, rpzDropIpRuleEnabledUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_RpzDropIpRuleMinPrefixLengthIpv4(t *testing.T) {
	var resourceName = "nios_dns_view.test_rpz_drop_ip_rule_min_prefix_length_ipv4"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	rpzDropIpRuleMinPrefixLengthIpv4 := 30
	rpzDropIpRuleMinPrefixLengthIpv4Update := 25

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewRpzDropIpRuleMinPrefixLengthIpv4(name, rpzDropIpRuleMinPrefixLengthIpv4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv4", "30"),
				),
			},
			// Update and Read
			{
				Config: testAccViewRpzDropIpRuleMinPrefixLengthIpv4(name, rpzDropIpRuleMinPrefixLengthIpv4Update),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv4", "25"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_RpzDropIpRuleMinPrefixLengthIpv6(t *testing.T) {
	var resourceName = "nios_dns_view.test_rpz_drop_ip_rule_min_prefix_length_ipv6"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	rpzDropIpRuleMinPrefixLengthIpv6 := 64
	rpzDropIpRuleMinPrefixLengthIpv6Update := 48

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewRpzDropIpRuleMinPrefixLengthIpv6(name, rpzDropIpRuleMinPrefixLengthIpv6),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv6", "64"),
				),
			},
			// Update and Read
			{
				Config: testAccViewRpzDropIpRuleMinPrefixLengthIpv6(name, rpzDropIpRuleMinPrefixLengthIpv6Update),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv6", "48"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_RpzQnameWaitRecurse(t *testing.T) {
	var resourceName = "nios_dns_view.test_rpz_qname_wait_recurse"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	rpzQnameWaitRecurse := true
	rpzQnameWaitRecurseUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewRpzQnameWaitRecurse(name, rpzQnameWaitRecurse),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_qname_wait_recurse", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewRpzQnameWaitRecurse(name, rpzQnameWaitRecurseUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_qname_wait_recurse", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// TODO
func TestAccViewResource_ScavengingSettings(t *testing.T) {
	var resourceName = "nios_dns_view.test_scavenging_settings"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	scavengingSettings := map[string]any{
		"enable_scavenging": true,
		"expression_list": []map[string]any{
			{
				"op":       "AND",
				"op1_type": "LIST",
			},
			{
				"op":       "EQ",
				"op1":      "rtype",
				"op1_type": "FIELD",
				"op2":      "A",
				"op2_type": "STRING",
			},
			{
				"op": "ENDLIST",
			},
		},
	}
	updatedScavengingSettings := map[string]any{
		"enable_scavenging": true,
		"expression_list": []map[string]any{
			{
				"op":       "AND",
				"op1_type": "LIST",
			},
			{
				"op":       "EQ",
				"op1":      "rtype",
				"op1_type": "FIELD",
				"op2":      "AAAA",
				"op2_type": "STRING",
			},
			{
				"op": "ENDLIST",
			},
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewScavengingSettings(name, scavengingSettings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.enable_auto_reclamation", "false"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.enable_recurrent_scavenging", "false"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.enable_rr_last_queried", "false"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.enable_scavenging", "true"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.enable_zone_last_queried", "false"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.reclaim_associated_records", "false"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.0.op", "AND"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.0.op1_type", "LIST"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op1", "rtype"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op1_type", "FIELD"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op2", "A"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op2_type", "STRING"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.2.op", "ENDLIST"),
				),
			},
			// Update and Read
			{
				Config: testAccViewScavengingSettings(name, updatedScavengingSettings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.enable_auto_reclamation", "false"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.enable_recurrent_scavenging", "false"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.enable_rr_last_queried", "false"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.enable_scavenging", "true"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.enable_zone_last_queried", "false"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.reclaim_associated_records", "false"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.0.op", "AND"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.0.op1_type", "LIST"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op1", "rtype"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op1_type", "FIELD"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op2", "AAAA"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op2_type", "STRING"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.2.op", "ENDLIST"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Sortlist(t *testing.T) {
	var resourceName = "nios_dns_view.test_sortlist"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	address := "13.0.0.0/24"
	addressUpdate := "10.0.0.0/24"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewSortlist(name, address),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sortlist.0.address", address),
				),
			},
			// Update and Read
			{
				Config: testAccViewSortlist(name, addressUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sortlist.0.address", addressUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseBlacklist(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_blacklist"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useBlacklist := true
	useBlacklistUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseBlacklist(name, useBlacklist),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blacklist", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseBlacklist(name, useBlacklistUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blacklist", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseDdnsForceCreationTimestampUpdate(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_ddns_force_creation_timestamp_update"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useDdnsForceCreationTimestampUpdate := true
	useDdnsForceCreationTimestampUpdateUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseDdnsForceCreationTimestampUpdate(name, useDdnsForceCreationTimestampUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_force_creation_timestamp_update", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseDdnsForceCreationTimestampUpdate(name, useDdnsForceCreationTimestampUpdateUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_force_creation_timestamp_update", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseDdnsPatternsRestriction(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_ddns_patterns_restriction"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useDdnsPatternsRestriction := true
	useDdnsPatternsRestrictionUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseDdnsPatternsRestriction(name, useDdnsPatternsRestriction),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_patterns_restriction", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseDdnsPatternsRestriction(name, useDdnsPatternsRestrictionUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_patterns_restriction", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseDdnsPrincipalSecurity(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_ddns_principal_security"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useDdnsPrincipalSecurity := true
	useDdnsPrincipalSecurityUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseDdnsPrincipalSecurity(name, useDdnsPrincipalSecurity),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_principal_security", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseDdnsPrincipalSecurity(name, useDdnsPrincipalSecurityUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_principal_security", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseDdnsRestrictProtected(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_ddns_restrict_protected"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useDdnsRestrictProtected := true
	useDdnsRestrictProtectedUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseDdnsRestrictProtected(name, useDdnsRestrictProtected),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_protected", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseDdnsRestrictProtected(name, useDdnsRestrictProtectedUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_protected", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseDdnsRestrictStatic(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_ddns_restrict_static"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useDdnsRestrictStatic := true
	useDdnsRestrictStaticUpdate := false
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseDdnsRestrictStatic(name, useDdnsRestrictStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_static", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseDdnsRestrictStatic(name, useDdnsRestrictStaticUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_static", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseDns64(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_dns64"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useDns64 := true
	useDns64Update := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseDns64(name, useDns64),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dns64", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseDns64(name, useDns64Update),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dns64", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseDnssec(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_dnssec"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useDnssec := true
	useDnssecUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseDnssec(name, useDnssec),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dnssec", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseDnssec(name, useDnssecUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dnssec", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseEdnsUdpSize(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_edns_udp_size"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useEdnsUpSize := true
	useEdnsUpSizeUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseEdnsUdpSize(name, useEdnsUpSize),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_edns_udp_size", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseEdnsUdpSize(name, useEdnsUpSizeUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_edns_udp_size", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseFilterAaaa(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_filter_aaaa"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useFilterAaaa := true
	useFilterAaaaUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseFilterAaaa(name, useFilterAaaa),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_filter_aaaa", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseFilterAaaa(name, useFilterAaaaUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_filter_aaaa", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseFixedRrsetOrderFqdns(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_fixed_rrset_order_fqdns"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useFixedRrsetOrderFqdns := true
	useFixedRrsetOrderFqdnsUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseFixedRrsetOrderFqdns(name, useFixedRrsetOrderFqdns),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_fixed_rrset_order_fqdns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseFixedRrsetOrderFqdns(name, useFixedRrsetOrderFqdnsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_fixed_rrset_order_fqdns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseForwarders(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_forwarders"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useForwarders := true
	useForwardersUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseForwarders(name, useForwarders),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_forwarders", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseForwarders(name, useForwardersUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_forwarders", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseMaxCacheTtl(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_max_cache_ttl"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useMaxCacheTtl := true
	useMaxCacheTtlUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseMaxCacheTtl(name, useMaxCacheTtl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_max_cache_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseMaxCacheTtl(name, useMaxCacheTtlUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_max_cache_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseMaxNcacheTtl(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_max_ncache_ttl"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useMaxNcacheTtl := true
	useMaxNcacheTtlUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseMaxNcacheTtl(name, useMaxNcacheTtl),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_max_ncache_ttl", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseMaxNcacheTtl(name, useMaxNcacheTtlUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_max_ncache_ttl", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseMaxUdpSize(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_max_udp_size"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useMaxUdpSize := true
	useMaxUdpSizeUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseMaxUdpSize(name, useMaxUdpSize),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_max_udp_size", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseMaxUdpSize(name, useMaxUdpSizeUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_max_udp_size", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseNxdomainRedirect(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_nxdomain_redirect"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useNxdomainRedirect := true
	useNxdomainRedirectUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseNxdomainRedirect(name, useNxdomainRedirect),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nxdomain_redirect", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseNxdomainRedirect(name, useNxdomainRedirectUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nxdomain_redirect", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseRecursion(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_recursion"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useRecursion := true
	useRecursionUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseRecursion(name, useRecursion),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recursion", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseRecursion(name, useRecursionUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recursion", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseResponseRateLimiting(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_response_rate_limiting"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useResponseRateLimiting := true
	useResponseRateLimitingUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseResponseRateLimiting(name, useResponseRateLimiting),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_response_rate_limiting", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseResponseRateLimiting(name, useResponseRateLimitingUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_response_rate_limiting", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseRootNameServer(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_root_name_server"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useRootNameServer := true
	useRootNameServerUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseRootNameServer(name, useRootNameServer),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_root_name_server", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseRootNameServer(name, useRootNameServerUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_root_name_server", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseRpzDropIpRule(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_rpz_drop_ip_rule"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useRpzDropIpRule := true
	useRpzDropIpRuleUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseRpzDropIpRule(name, useRpzDropIpRule),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_rpz_drop_ip_rule", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseRpzDropIpRule(name, useRpzDropIpRuleUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_rpz_drop_ip_rule", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseRpzQnameWaitRecurse(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_rpz_qname_wait_recurse"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useRpzQnameWaitRecurse := true
	useRpzQnameWaitRecurseUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseRpzQnameWaitRecurse(name, useRpzQnameWaitRecurse),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_rpz_qname_wait_recurse", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseRpzQnameWaitRecurse(name, useRpzQnameWaitRecurseUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_rpz_qname_wait_recurse", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseScavengingSettings(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_scavenging_settings"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useScavengingSettings := true
	useScavengingSettingsUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseScavengingSettings(name, useScavengingSettings),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_scavenging_settings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseScavengingSettings(name, useScavengingSettingsUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_scavenging_settings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseSortlist(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_sortlist"
	var v dns.View
	name := acctest.RandomNameWithPrefix("view")
	useSortlist := true
	useSortlistUpdate := false

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseSortlist(name, useSortlist),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_sortlist", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseSortlist(name, useSortlistUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_sortlist", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckViewExists(ctx context.Context, resourceName string, v *dns.View) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			ViewAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForView).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetViewResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetViewResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckViewDestroy(ctx context.Context, v *dns.View) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			ViewAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForView).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				// resource was deleted
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckViewDisappears(ctx context.Context, v *dns.View) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			ViewAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccViewBasicConfig(name string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dns_view" "test" {
	name = %q
}
`, name)
}

func testAccViewBlacklistAction(name, blacklistAction string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_blacklist_action" {
	name = %q
    blacklist_action = %q
	use_blacklist = true
}
`, name, blacklistAction)
}

func testAccViewBlacklistLogQuery(name string, blacklistLogQuery bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_blacklist_log_query" {
	name = %q
    blacklist_log_query = %t
	use_blacklist = true
}
`, name, blacklistLogQuery)
}

func testAccViewBlacklistRedirectAddresses(name string, blacklistRedirectAddresses []string) string {
	blacklistRedirectAddressesStr := utils.ConvertStringSliceToHCL(blacklistRedirectAddresses)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_blacklist_redirect_addresses" {
	name = %q
    blacklist_redirect_addresses = %s
	use_blacklist = true
}
`, name, blacklistRedirectAddressesStr)
}

func testAccViewBlacklistRedirectTtl(name string, blacklistRedirectTtl int) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_blacklist_redirect_ttl" {
	name = %q
    blacklist_redirect_ttl = %d
	use_blacklist = true
}
`, name, blacklistRedirectTtl)
}

func testAccViewBlacklistRulesets(name string, blacklistRulesets []string) string {
	blacklistRulesetsStr := utils.ConvertStringSliceToHCL(blacklistRulesets)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_blacklist_rulesets" {
	name = %q
    blacklist_rulesets = %s
}
`, name, blacklistRulesetsStr)
}

func testAccViewComment(name, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_comment" {
	name = %q
    comment = %q
}
`, name, comment)
}

func testAccViewCustomRootNameServers(name string, customRootNameServers []map[string]any) string {

	customRootNameServersStr := utils.ConvertSliceOfMapsToHCL(customRootNameServers)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_custom_root_name_servers" {
	name = %q
    custom_root_name_servers = %s
}
`, name, customRootNameServersStr)
}

func testAccViewDdnsForceCreationTimestampUpdate(name string, ddnsForceCreationTimestampUpdate bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_ddns_force_creation_timestamp_update" {
	name = %q
    ddns_force_creation_timestamp_update = %t
	use_ddns_force_creation_timestamp_update = true
}
`, name, ddnsForceCreationTimestampUpdate)
}

func testAccViewDdnsPrincipalGroup(name, ddnsPrincipalGroup string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_ddns_principal_group" {
	name = %q
    ddns_principal_group = %q
	use_ddns_principal_security  = true
}
`, name, ddnsPrincipalGroup)
}

func testAccViewDdnsPrincipalTracking(name string, ddnsPrincipalTracking bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_ddns_principal_tracking" {
	name = %q
    ddns_principal_tracking = %t
	use_ddns_principal_security = true
}
`, name, ddnsPrincipalTracking)
}

func testAccViewDdnsRestrictPatterns(name string, ddnsRestrictPatterns bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_ddns_restrict_patterns" {
	name = %q
    ddns_restrict_patterns = %t
	use_ddns_patterns_restriction = true
}
`, name, ddnsRestrictPatterns)
}

func testAccViewDdnsRestrictPatternsList(name string, ddnsRestrictPatternsList []string) string {
	ddnsRestrictPatternsListStr := utils.ConvertStringSliceToHCL(ddnsRestrictPatternsList)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_ddns_restrict_patterns_list" {
	name = %q
    ddns_restrict_patterns_list = %s
}
`, name, ddnsRestrictPatternsListStr)
}

func testAccViewDdnsRestrictProtected(name string, ddnsRestrictProtected bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_ddns_restrict_protected" {
	name = %q
    ddns_restrict_protected = %t
	use_ddns_restrict_protected= true
}
`, name, ddnsRestrictProtected)
}

func testAccViewDdnsRestrictSecure(name string, ddnsRestrictSecure bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_ddns_restrict_secure" {
	name = %q
    ddns_restrict_secure = %t
	use_ddns_principal_security = true
}
`, name, ddnsRestrictSecure)
}

func testAccViewDdnsRestrictStatic(name string, ddnsRestrictStatic bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_ddns_restrict_static" {
	name = %q
	ddns_restrict_static = %t
	use_ddns_restrict_static = true
}
`, name, ddnsRestrictStatic)
}

func testAccViewDisable(name string, disable bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_disable" {
	name = %q
    disable = %t
}
`, name, disable)
}

func testAccViewDns64Enabled(name string, dns64Enabled bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_dns64_enabled" {
	name = %q
	dns64_enabled = %t
	use_dns64 = true
	dns64_groups = [
		"dns64_group"
	]
}
`, name, dns64Enabled)
}

func testAccViewDns64Groups(name string, dns64Groups []string) string {
	dns64GroupString := utils.ConvertStringSliceToHCL(dns64Groups)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_dns64_groups" {
	name = %q
	dns64_groups = %s
	use_dns64 = true
}
`, name, dns64GroupString)
}

func testAccViewDnssecEnabled(name string, dnssecEnabled bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_dnssec_enabled" {
	name = %q
	dnssec_enabled = %t
	use_dnssec = true
}
`, name, dnssecEnabled)
}

func testAccViewDnssecExpiredSignaturesEnabled(name string, dnssecExpiredSignaturesEnabled bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_dnssec_expired_signatures_enabled" {
	name = %q
	dnssec_expired_signatures_enabled = %t
	use_dnssec = true
}
`, name, dnssecExpiredSignaturesEnabled)
}

func testAccViewDnssecNegativeTrustAnchors(name string, dnssecNegativeTrustAnchors []string) string {
	dnssecNegativeTrustAnchorsStr := utils.ConvertStringSliceToHCL(dnssecNegativeTrustAnchors)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_dnssec_negative_trust_anchors" {
	name = %q
    dnssec_negative_trust_anchors = %s
}
`, name, dnssecNegativeTrustAnchorsStr)
}

func testAccViewDnssecTrustedKeys(name string, dnssecTrustedKeys []map[string]any) string {
	dnssecTrustedKeysStr := utils.ConvertSliceOfMapsToHCL(dnssecTrustedKeys)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_dnssec_trusted_keys" {
	name = %q
    dnssec_trusted_keys = %s
	use_dnssec = true
}
`, name, dnssecTrustedKeysStr)
}

func testAccViewDnssecValidationEnabled(name string, dnssecValidationEnabled bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_dnssec_validation_enabled" {
	name = %q
    dnssec_validation_enabled = %t
	use_dnssec = true
}
`, name, dnssecValidationEnabled)
}

func testAccViewEdnsUdpSize(name string, ednsUdpSize int) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_edns_udp_size" {
	name = %q
    edns_udp_size = %d
}
`, name, ednsUdpSize)
}

func testAccViewEnableBlacklist(name string, enableBlacklist bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_enable_blacklist" {
	name = %q
	enable_blacklist = %t
	use_blacklist = true
	blacklist_redirect_addresses = ["10.0.0.2"]
	blacklist_rulesets = ["ruleset1"]
}
`, name, enableBlacklist)
}

func testAccViewEnableFixedRrsetOrderFqdns(name string, enableFixedRrsetOrderFqdns bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_enable_fixed_rrset_order_fqdns" {
	name = %q
    enable_fixed_rrset_order_fqdns = %t
	use_fixed_rrset_order_fqdns = true
}
`, name, enableFixedRrsetOrderFqdns)
}

func testAccViewEnableMatchRecursiveOnly(name string, enableMatchRecursiveOnly bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_enable_match_recursive_only" {
	name = %q
    enable_match_recursive_only = %t
}
`, name, enableMatchRecursiveOnly)
}

func testAccViewExtAttrs(name string, extAttrs map[string]any) string {
	extAttrsStr := utils.ConvertMapToHCL(extAttrs)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_extattrs" {
	name = %q
    extattrs = %s
}
`, name, extAttrsStr)
}

func testAccViewFilterAaaa(name, filterAaaa string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_filter_aaaa" {
	name = %q
    filter_aaaa = %q
	use_filter_aaaa = true
}
`, name, filterAaaa)
}

func testAccViewFilterAaaaList(name string, filterAaaaList []map[string]any) string {
	filterAaaaListStr := utils.ConvertSliceOfMapsToHCL(filterAaaaList)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_filter_aaaa_list" {
	name = %q
	filter_aaaa_list = %s
	use_filter_aaaa = true
}
`, name, filterAaaaListStr)
}

func testAccViewFixedRrsetOrderFqdns(name string, fixedRrsetOrderFqdns []map[string]any) string {
	fixedRrsetOrderFqdnsStr := utils.ConvertSliceOfMapsToHCL(fixedRrsetOrderFqdns)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_fixed_rrset_order_fqdns" {
	name = %q
    fixed_rrset_order_fqdns = %s
}
`, name, fixedRrsetOrderFqdnsStr)
}

func testAccViewForwardOnly(name string, forwardOnly bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_forward_only" {
	name = %q
	forward_only = %t
	use_forwarders = true
	forwarders = ["10.192.81.23"]
}
`, name, forwardOnly)
}

func testAccViewForwarders(name string, forwarders []string) string {
	forwardersStr := utils.ConvertStringSliceToHCL(forwarders)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_forwarders" {
	name = %q
    forwarders = %s
	use_forwarders = true
}
`, name, forwardersStr)
}

func testAccViewLastQueriedAcl(name string, lastQueriedAcl []map[string]any) string {
	lastQueriedAclStr := utils.ConvertSliceOfMapsToHCL(lastQueriedAcl)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_last_queried_acl" {
	name = %q
	last_queried_acl = %s
}
`, name, lastQueriedAclStr)
}

func testAccViewMatchClients(name string, matchClients []map[string]any) string {
	matchClientsHCL := utils.ConvertSliceOfMapsToHCL(matchClients)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_match_clients" {
	name = %q
    match_clients = %s
}
`, name, matchClientsHCL)
}

func testAccViewMatchDestinations(name string, matchDestinations []map[string]any) string {
	matchDestinationsHCL := utils.ConvertSliceOfMapsToHCL(matchDestinations)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_match_destinations" {
	name = %q
    match_destinations = %s
}
`, name, matchDestinationsHCL)
}

func testAccViewMaxCacheTtl(name string, maxCacheTtl int) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_max_cache_ttl" {
	name = %q
	max_cache_ttl = %d
	use_max_cache_ttl = true
}
`, name, maxCacheTtl)
}

func testAccViewMaxNcacheTtl(name string, maxNcacheTtl int) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_max_ncache_ttl" {
	name = %q
	max_ncache_ttl = %d
	use_max_ncache_ttl = true
}
`, name, maxNcacheTtl)
}

func testAccViewMaxUdpSize(name string, maxUdpSize int) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_max_udp_size" {
	name = %q
	max_udp_size = %d
	use_max_udp_size = true
}
`, name, maxUdpSize)
}

func testAccViewName(name string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_name" {
    name = %q
}
`, name)
}

func testAccViewNetworkView(name, networkView string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_network_view" {
	name = %q
    network_view = %q
}
`, name, networkView)
}

func testAccViewNotifyDelay(name string, notifyDelay int) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_notify_delay" {
	name = %q
    notify_delay = %d
}
`, name, notifyDelay)
}

func testAccViewNxdomainLogQuery(name string, nxdomainLogQuery bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_nxdomain_log_query" {
	name = %q
	nxdomain_log_query = %t
	use_nxdomain_redirect =true
}
`, name, nxdomainLogQuery)
}

func testAccViewNxdomainRedirect(name string, nxdomainRedirect bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_nxdomain_redirect" {
	name = %q
	nxdomain_redirect = %t
	use_nxdomain_redirect = true
	nxdomain_redirect_addresses = ["10.45.3.2"]
}
`, name, nxdomainRedirect)
}

func testAccViewNxdomainRedirectAddresses(name string, nxdomainRedirectAddresses []string) string {
	nxdomainRedirectAddressesStr := utils.ConvertStringSliceToHCL(nxdomainRedirectAddresses)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_nxdomain_redirect_addresses" {
	name = %q
	nxdomain_redirect_addresses = %s
	use_nxdomain_redirect = true
}
`, name, nxdomainRedirectAddressesStr)
}

func testAccViewNxdomainRedirectAddressesV6(name string, nxdomainRedirectAddressesV6 []string) string {
	nxdomainRedirectAddressesV6Str := utils.ConvertStringSliceToHCL(nxdomainRedirectAddressesV6)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_nxdomain_redirect_addresses_v6" {
    name = %q
    nxdomain_redirect_addresses_v6 = %s
    use_nxdomain_redirect = true
}
`, name, nxdomainRedirectAddressesV6Str)
}

func testAccViewNxdomainRedirectTtl(name string, nxdomainRedirectTtl int) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_nxdomain_redirect_ttl" {
	name = %q
	nxdomain_redirect_ttl = %d
	use_nxdomain_redirect =true
}
`, name, nxdomainRedirectTtl)
}

func testAccViewNxdomainRulesets(name string, nxdomainRulesets []string) string {
	nxdomainRulesetsStr := utils.ConvertStringSliceToHCL(nxdomainRulesets)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_nxdomain_rulesets" {
	name = %q
    nxdomain_rulesets = %s
	use_nxdomain_redirect = true
}
`, name, nxdomainRulesetsStr)
}

func testAccViewRecursion(name string, recursion bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_recursion" {
	name = %q
	recursion = %t
	use_recursion = true
}
`, name, recursion)
}

func testAccViewResponseRateLimiting(name string, responseRateLimiting map[string]any) string {
	responseRateLimitingStr := utils.ConvertMapToHCL(responseRateLimiting)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_response_rate_limiting" {
	name = %q
    response_rate_limiting = %s
	use_response_rate_limiting = true
}
`, name, responseRateLimitingStr)
}

func testAccViewRootNameServerType(name string, rootNameServerType string, customRootNameServers []map[string]any) string {
	customRootNameServersStr := utils.ConvertSliceOfMapsToHCL(customRootNameServers)
	return fmt.Sprintf(`
resource "nios_dns_view" "test_root_name_server_type" {
	name = %q
    root_name_server_type = %q
    custom_root_name_servers = %s
	use_root_name_server = true

}
`, name, rootNameServerType, customRootNameServersStr)
}

func testAccViewRpzDropIpRuleEnabled(name string, rpzDropIpRuleEnabled bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_rpz_drop_ip_rule_enabled" {
	name = %q
	rpz_drop_ip_rule_enabled = %t
	use_rpz_drop_ip_rule = true
}
`, name, rpzDropIpRuleEnabled)
}

func testAccViewRpzDropIpRuleMinPrefixLengthIpv4(name string, rpzDropIpRuleMinPrefixLengthIpv4 int) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_rpz_drop_ip_rule_min_prefix_length_ipv4" {
	name = %q
    rpz_drop_ip_rule_min_prefix_length_ipv4 = %d
	use_rpz_drop_ip_rule = true
}
`, name, rpzDropIpRuleMinPrefixLengthIpv4)
}

func testAccViewRpzDropIpRuleMinPrefixLengthIpv6(name string, rpzDropIpRuleMinPrefixLengthIpv6 int) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_rpz_drop_ip_rule_min_prefix_length_ipv6" {
	name = %q
	rpz_drop_ip_rule_min_prefix_length_ipv6 = %d
	use_rpz_drop_ip_rule = true
}
`, name, rpzDropIpRuleMinPrefixLengthIpv6)
}

func testAccViewRpzQnameWaitRecurse(name string, rpzQnameWaitRecurse bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_rpz_qname_wait_recurse" {
	name = %q
	rpz_qname_wait_recurse = %t
	use_rpz_qname_wait_recurse = true
}
`, name, rpzQnameWaitRecurse)
}

func testAccViewScavengingSettings(name string, scavengingSetting map[string]any) string {
	scavengingSettingsHCL := utils.ConvertMapToHCL(scavengingSetting)

	return fmt.Sprintf(`
resource "nios_dns_view" "test_scavenging_settings" {
	name = %q
	use_scavenging_settings = true
    scavenging_settings = %s
}
`, name, scavengingSettingsHCL)
}

func testAccViewSortlist(name, address string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_sortlist" {
	name = %q
    sortlist = [
	{
			"address": %q
		}]
}
`, name, address)
}

func testAccViewUseBlacklist(name string, useBlacklist bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_blacklist" {
	name = %q
	use_blacklist = %t
}
`, name, useBlacklist)
}

func testAccViewUseDdnsForceCreationTimestampUpdate(name string, useDdnsForceCreationTimestampUpdate bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_ddns_force_creation_timestamp_update" {
	name = %q
    use_ddns_force_creation_timestamp_update = %t
}
`, name, useDdnsForceCreationTimestampUpdate)
}

func testAccViewUseDdnsPatternsRestriction(name string, useDdnsPatternsRestriction bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_ddns_patterns_restriction" {
	name = %q
    use_ddns_patterns_restriction = %t
}
`, name, useDdnsPatternsRestriction)
}

func testAccViewUseDdnsPrincipalSecurity(name string, useDdnsPrincipalSecurity bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_ddns_principal_security" {
	name = %q
	use_ddns_principal_security = %t
}
`, name, useDdnsPrincipalSecurity)
}

func testAccViewUseDdnsRestrictProtected(name string, useDdnsRestrictProtected bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_ddns_restrict_protected" {
	name = %q
	use_ddns_restrict_protected = %t
}
`, name, useDdnsRestrictProtected)
}

func testAccViewUseDdnsRestrictStatic(name string, useDdnsRestrictStatic bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_ddns_restrict_static" {
	name = %q
	use_ddns_restrict_static = %t
}
`, name, useDdnsRestrictStatic)
}

func testAccViewUseDns64(name string, useDns64 bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_dns64" {
	name = %q
    use_dns64 = %t
}
`, name, useDns64)
}

func testAccViewUseDnssec(name string, useDnssec bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_dnssec" {
	name = %q
    use_dnssec = %t
}
`, name, useDnssec)
}

func testAccViewUseEdnsUdpSize(name string, useEdnsUdpSize bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_edns_udp_size" {
	name = %q
	use_edns_udp_size = %t
}
`, name, useEdnsUdpSize)
}

func testAccViewUseFilterAaaa(name string, useFilterAaaa bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_filter_aaaa" {
	name = %q
	use_filter_aaaa = %t
}
`, name, useFilterAaaa)
}

func testAccViewUseFixedRrsetOrderFqdns(name string, useFixedRrsetOrderFqdns bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_fixed_rrset_order_fqdns" {
	name = %q
	use_fixed_rrset_order_fqdns = %t
}
`, name, useFixedRrsetOrderFqdns)
}

func testAccViewUseForwarders(name string, useForwarders bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_forwarders" {
	name = %q
    use_forwarders = %t
}
`, name, useForwarders)
}

func testAccViewUseMaxCacheTtl(name string, useMaxCacheTtl bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_max_cache_ttl" {
	name = %q
	use_max_cache_ttl = %t
}
`, name, useMaxCacheTtl)
}

func testAccViewUseMaxNcacheTtl(name string, useMaxNcacheTtl bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_max_ncache_ttl" {
    name = %q
    use_max_ncache_ttl = %t
}
`, name, useMaxNcacheTtl)
}

func testAccViewUseMaxUdpSize(name string, useMaxUdpSize bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_max_udp_size" {
	name = %q
    use_max_udp_size = %t
}
`, name, useMaxUdpSize)
}

func testAccViewUseNxdomainRedirect(name string, useNxdomainRedirect bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_nxdomain_redirect" {
	name = %q
	use_nxdomain_redirect = %t
}
`, name, useNxdomainRedirect)
}

func testAccViewUseRecursion(name string, useRecursion bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_recursion" {
	name = %q
    use_recursion = %t
}
`, name, useRecursion)
}

func testAccViewUseResponseRateLimiting(name string, useResponseRateLimiting bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_response_rate_limiting" {
	name = %q
	use_response_rate_limiting = %t
}
`, name, useResponseRateLimiting)
}

func testAccViewUseRootNameServer(name string, useRootNameServer bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_root_name_server" {
	name = %q
	use_root_name_server = %t
}
`, name, useRootNameServer)
}

func testAccViewUseRpzDropIpRule(name string, useRpzDropIpRule bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_rpz_drop_ip_rule" {
	name = %q
    use_rpz_drop_ip_rule = %t
}
`, name, useRpzDropIpRule)
}

func testAccViewUseRpzQnameWaitRecurse(name string, useRpzQnameWaitRecurse bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_rpz_qname_wait_recurse" {
	name = %q
	use_rpz_qname_wait_recurse = %t
}
`, name, useRpzQnameWaitRecurse)
}

func testAccViewUseScavengingSettings(name string, useScavengingSettings bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_scavenging_settings" {
	name = %q
	use_scavenging_settings = %t
}
`, name, useScavengingSettings)
}

func testAccViewUseSortlist(name string, useSortlist bool) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_sortlist" {
	name = %q
    use_sortlist = %t
}
`, name, useSortlist)
}


