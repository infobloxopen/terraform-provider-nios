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

var readableAttributesForView = "blacklist_action,blacklist_log_query,blacklist_redirect_addresses,blacklist_redirect_ttl,blacklist_rulesets,cloud_info,comment,custom_root_name_servers,ddns_force_creation_timestamp_update,ddns_principal_group,ddns_principal_tracking,ddns_restrict_patterns,ddns_restrict_patterns_list,ddns_restrict_protected,ddns_restrict_secure,ddns_restrict_static,disable,dns64_enabled,dns64_groups,dnssec_enabled,dnssec_expired_signatures_enabled,dnssec_negative_trust_anchors,dnssec_trusted_keys,dnssec_validation_enabled,edns_udp_size,enable_blacklist,enable_fixed_rrset_order_fqdns,enable_match_recursive_only,extattrs,filter_aaaa,filter_aaaa_list,fixed_rrset_order_fqdns,forward_only,forwarders,is_default,last_queried_acl,match_clients,match_destinations,max_cache_ttl,max_ncache_ttl,max_udp_size,name,network_view,notify_delay,nxdomain_log_query,nxdomain_redirect,nxdomain_redirect_addresses,nxdomain_redirect_addresses_v6,nxdomain_redirect_ttl,nxdomain_rulesets,recursion,response_rate_limiting,root_name_server_type,rpz_drop_ip_rule_enabled,rpz_drop_ip_rule_min_prefix_length_ipv4,rpz_drop_ip_rule_min_prefix_length_ipv6,rpz_qname_wait_recurse,scavenging_settings,sortlist,use_blacklist,use_ddns_force_creation_timestamp_update,use_ddns_patterns_restriction,use_ddns_principal_security,use_ddns_restrict_protected,use_ddns_restrict_static,use_dns64,use_dnssec,use_edns_udp_size,use_filter_aaaa,use_fixed_rrset_order_fqdns,use_forwarders,use_max_cache_ttl,use_max_ncache_ttl,use_max_udp_size,use_nxdomain_redirect,use_recursion,use_response_rate_limiting,use_root_name_server,use_rpz_drop_ip_rule,use_rpz_qname_wait_recurse,use_scavenging_settings,use_sortlist"

func TestAccViewResource_basic(t *testing.T) {
	var resourceName = "nios_dns_view.test"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_disappears(t *testing.T) {
	resourceName := "nios_dns_view.test"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckViewDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccViewBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					testAccCheckViewDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccViewResource_Ref(t *testing.T) {
	var resourceName = "nios_dns_view.test_ref"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_BlacklistAction(t *testing.T) {
	var resourceName = "nios_dns_view.test_blacklist_action"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewBlacklistAction("BLACKLIST_ACTION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_action", "BLACKLIST_ACTION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewBlacklistAction("BLACKLIST_ACTION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_action", "BLACKLIST_ACTION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_BlacklistLogQuery(t *testing.T) {
	var resourceName = "nios_dns_view.test_blacklist_log_query"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewBlacklistLogQuery("BLACKLIST_LOG_QUERY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_log_query", "BLACKLIST_LOG_QUERY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewBlacklistLogQuery("BLACKLIST_LOG_QUERY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_log_query", "BLACKLIST_LOG_QUERY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_BlacklistRedirectAddresses(t *testing.T) {
	var resourceName = "nios_dns_view.test_blacklist_redirect_addresses"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewBlacklistRedirectAddresses("BLACKLIST_REDIRECT_ADDRESSES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_redirect_addresses", "BLACKLIST_REDIRECT_ADDRESSES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewBlacklistRedirectAddresses("BLACKLIST_REDIRECT_ADDRESSES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_redirect_addresses", "BLACKLIST_REDIRECT_ADDRESSES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_BlacklistRedirectTtl(t *testing.T) {
	var resourceName = "nios_dns_view.test_blacklist_redirect_ttl"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewBlacklistRedirectTtl("BLACKLIST_REDIRECT_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_redirect_ttl", "BLACKLIST_REDIRECT_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewBlacklistRedirectTtl("BLACKLIST_REDIRECT_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_redirect_ttl", "BLACKLIST_REDIRECT_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_BlacklistRulesets(t *testing.T) {
	var resourceName = "nios_dns_view.test_blacklist_rulesets"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewBlacklistRulesets("BLACKLIST_RULESETS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_rulesets", "BLACKLIST_RULESETS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewBlacklistRulesets("BLACKLIST_RULESETS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "blacklist_rulesets", "BLACKLIST_RULESETS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_CloudInfo(t *testing.T) {
	var resourceName = "nios_dns_view.test_cloud_info"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewCloudInfo("CLOUD_INFO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewCloudInfo("CLOUD_INFO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_view.test_comment"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_CustomRootNameServers(t *testing.T) {
	var resourceName = "nios_dns_view.test_custom_root_name_servers"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewCustomRootNameServers("CUSTOM_ROOT_NAME_SERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "custom_root_name_servers", "CUSTOM_ROOT_NAME_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewCustomRootNameServers("CUSTOM_ROOT_NAME_SERVERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "custom_root_name_servers", "CUSTOM_ROOT_NAME_SERVERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DdnsForceCreationTimestampUpdate(t *testing.T) {
	var resourceName = "nios_dns_view.test_ddns_force_creation_timestamp_update"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDdnsForceCreationTimestampUpdate("DDNS_FORCE_CREATION_TIMESTAMP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_force_creation_timestamp_update", "DDNS_FORCE_CREATION_TIMESTAMP_UPDATE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDdnsForceCreationTimestampUpdate("DDNS_FORCE_CREATION_TIMESTAMP_UPDATE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_force_creation_timestamp_update", "DDNS_FORCE_CREATION_TIMESTAMP_UPDATE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DdnsPrincipalGroup(t *testing.T) {
	var resourceName = "nios_dns_view.test_ddns_principal_group"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDdnsPrincipalGroup("DDNS_PRINCIPAL_GROUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_group", "DDNS_PRINCIPAL_GROUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDdnsPrincipalGroup("DDNS_PRINCIPAL_GROUP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_group", "DDNS_PRINCIPAL_GROUP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DdnsPrincipalTracking(t *testing.T) {
	var resourceName = "nios_dns_view.test_ddns_principal_tracking"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDdnsPrincipalTracking("DDNS_PRINCIPAL_TRACKING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_tracking", "DDNS_PRINCIPAL_TRACKING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDdnsPrincipalTracking("DDNS_PRINCIPAL_TRACKING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_tracking", "DDNS_PRINCIPAL_TRACKING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DdnsRestrictPatterns(t *testing.T) {
	var resourceName = "nios_dns_view.test_ddns_restrict_patterns"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDdnsRestrictPatterns("DDNS_RESTRICT_PATTERNS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns", "DDNS_RESTRICT_PATTERNS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDdnsRestrictPatterns("DDNS_RESTRICT_PATTERNS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns", "DDNS_RESTRICT_PATTERNS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DdnsRestrictPatternsList(t *testing.T) {
	var resourceName = "nios_dns_view.test_ddns_restrict_patterns_list"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDdnsRestrictPatternsList("DDNS_RESTRICT_PATTERNS_LIST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns_list", "DDNS_RESTRICT_PATTERNS_LIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDdnsRestrictPatternsList("DDNS_RESTRICT_PATTERNS_LIST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns_list", "DDNS_RESTRICT_PATTERNS_LIST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DdnsRestrictProtected(t *testing.T) {
	var resourceName = "nios_dns_view.test_ddns_restrict_protected"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDdnsRestrictProtected("DDNS_RESTRICT_PROTECTED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_protected", "DDNS_RESTRICT_PROTECTED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDdnsRestrictProtected("DDNS_RESTRICT_PROTECTED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_protected", "DDNS_RESTRICT_PROTECTED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DdnsRestrictSecure(t *testing.T) {
	var resourceName = "nios_dns_view.test_ddns_restrict_secure"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDdnsRestrictSecure("DDNS_RESTRICT_SECURE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_secure", "DDNS_RESTRICT_SECURE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDdnsRestrictSecure("DDNS_RESTRICT_SECURE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_secure", "DDNS_RESTRICT_SECURE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DdnsRestrictStatic(t *testing.T) {
	var resourceName = "nios_dns_view.test_ddns_restrict_static"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDdnsRestrictStatic("DDNS_RESTRICT_STATIC_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_static", "DDNS_RESTRICT_STATIC_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDdnsRestrictStatic("DDNS_RESTRICT_STATIC_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_static", "DDNS_RESTRICT_STATIC_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_view.test_disable"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Dns64Enabled(t *testing.T) {
	var resourceName = "nios_dns_view.test_dns64_enabled"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDns64Enabled("DNS64_ENABLED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns64_enabled", "DNS64_ENABLED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDns64Enabled("DNS64_ENABLED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns64_enabled", "DNS64_ENABLED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Dns64Groups(t *testing.T) {
	var resourceName = "nios_dns_view.test_dns64_groups"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDns64Groups("DNS64_GROUPS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns64_groups", "DNS64_GROUPS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDns64Groups("DNS64_GROUPS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns64_groups", "DNS64_GROUPS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DnssecEnabled(t *testing.T) {
	var resourceName = "nios_dns_view.test_dnssec_enabled"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDnssecEnabled("DNSSEC_ENABLED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_enabled", "DNSSEC_ENABLED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDnssecEnabled("DNSSEC_ENABLED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_enabled", "DNSSEC_ENABLED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DnssecExpiredSignaturesEnabled(t *testing.T) {
	var resourceName = "nios_dns_view.test_dnssec_expired_signatures_enabled"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDnssecExpiredSignaturesEnabled("DNSSEC_EXPIRED_SIGNATURES_ENABLED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_expired_signatures_enabled", "DNSSEC_EXPIRED_SIGNATURES_ENABLED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDnssecExpiredSignaturesEnabled("DNSSEC_EXPIRED_SIGNATURES_ENABLED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_expired_signatures_enabled", "DNSSEC_EXPIRED_SIGNATURES_ENABLED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DnssecNegativeTrustAnchors(t *testing.T) {
	var resourceName = "nios_dns_view.test_dnssec_negative_trust_anchors"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDnssecNegativeTrustAnchors("DNSSEC_NEGATIVE_TRUST_ANCHORS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_negative_trust_anchors", "DNSSEC_NEGATIVE_TRUST_ANCHORS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDnssecNegativeTrustAnchors("DNSSEC_NEGATIVE_TRUST_ANCHORS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_negative_trust_anchors", "DNSSEC_NEGATIVE_TRUST_ANCHORS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DnssecTrustedKeys(t *testing.T) {
	var resourceName = "nios_dns_view.test_dnssec_trusted_keys"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDnssecTrustedKeys("DNSSEC_TRUSTED_KEYS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trusted_keys", "DNSSEC_TRUSTED_KEYS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDnssecTrustedKeys("DNSSEC_TRUSTED_KEYS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trusted_keys", "DNSSEC_TRUSTED_KEYS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DnssecValidationEnabled(t *testing.T) {
	var resourceName = "nios_dns_view.test_dnssec_validation_enabled"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDnssecValidationEnabled("DNSSEC_VALIDATION_ENABLED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_validation_enabled", "DNSSEC_VALIDATION_ENABLED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDnssecValidationEnabled("DNSSEC_VALIDATION_ENABLED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_validation_enabled", "DNSSEC_VALIDATION_ENABLED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_EdnsUdpSize(t *testing.T) {
	var resourceName = "nios_dns_view.test_edns_udp_size"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewEdnsUdpSize("EDNS_UDP_SIZE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "edns_udp_size", "EDNS_UDP_SIZE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewEdnsUdpSize("EDNS_UDP_SIZE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "edns_udp_size", "EDNS_UDP_SIZE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_EnableBlacklist(t *testing.T) {
	var resourceName = "nios_dns_view.test_enable_blacklist"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewEnableBlacklist("ENABLE_BLACKLIST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_blacklist", "ENABLE_BLACKLIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewEnableBlacklist("ENABLE_BLACKLIST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_blacklist", "ENABLE_BLACKLIST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_EnableFixedRrsetOrderFqdns(t *testing.T) {
	var resourceName = "nios_dns_view.test_enable_fixed_rrset_order_fqdns"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewEnableFixedRrsetOrderFqdns("ENABLE_FIXED_RRSET_ORDER_FQDNS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_fixed_rrset_order_fqdns", "ENABLE_FIXED_RRSET_ORDER_FQDNS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewEnableFixedRrsetOrderFqdns("ENABLE_FIXED_RRSET_ORDER_FQDNS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_fixed_rrset_order_fqdns", "ENABLE_FIXED_RRSET_ORDER_FQDNS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_EnableMatchRecursiveOnly(t *testing.T) {
	var resourceName = "nios_dns_view.test_enable_match_recursive_only"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewEnableMatchRecursiveOnly("ENABLE_MATCH_RECURSIVE_ONLY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_match_recursive_only", "ENABLE_MATCH_RECURSIVE_ONLY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewEnableMatchRecursiveOnly("ENABLE_MATCH_RECURSIVE_ONLY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_match_recursive_only", "ENABLE_MATCH_RECURSIVE_ONLY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_view.test_extattrs"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_FilterAaaa(t *testing.T) {
	var resourceName = "nios_dns_view.test_filter_aaaa"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewFilterAaaa("FILTER_AAAA_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa", "FILTER_AAAA_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewFilterAaaa("FILTER_AAAA_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa", "FILTER_AAAA_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_FilterAaaaList(t *testing.T) {
	var resourceName = "nios_dns_view.test_filter_aaaa_list"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewFilterAaaaList("FILTER_AAAA_LIST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_list", "FILTER_AAAA_LIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewFilterAaaaList("FILTER_AAAA_LIST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_list", "FILTER_AAAA_LIST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_FixedRrsetOrderFqdns(t *testing.T) {
	var resourceName = "nios_dns_view.test_fixed_rrset_order_fqdns"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewFixedRrsetOrderFqdns("FIXED_RRSET_ORDER_FQDNS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fixed_rrset_order_fqdns", "FIXED_RRSET_ORDER_FQDNS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewFixedRrsetOrderFqdns("FIXED_RRSET_ORDER_FQDNS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fixed_rrset_order_fqdns", "FIXED_RRSET_ORDER_FQDNS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_ForwardOnly(t *testing.T) {
	var resourceName = "nios_dns_view.test_forward_only"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewForwardOnly("FORWARD_ONLY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forward_only", "FORWARD_ONLY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewForwardOnly("FORWARD_ONLY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forward_only", "FORWARD_ONLY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Forwarders(t *testing.T) {
	var resourceName = "nios_dns_view.test_forwarders"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewForwarders("FORWARDERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders", "FORWARDERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewForwarders("FORWARDERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders", "FORWARDERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_LastQueriedAcl(t *testing.T) {
	var resourceName = "nios_dns_view.test_last_queried_acl"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewLastQueriedAcl("LAST_QUERIED_ACL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "last_queried_acl", "LAST_QUERIED_ACL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewLastQueriedAcl("LAST_QUERIED_ACL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "last_queried_acl", "LAST_QUERIED_ACL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_MatchClients(t *testing.T) {
	var resourceName = "nios_dns_view.test_match_clients"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewMatchClients("MATCH_CLIENTS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_clients", "MATCH_CLIENTS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewMatchClients("MATCH_CLIENTS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_clients", "MATCH_CLIENTS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_MatchDestinations(t *testing.T) {
	var resourceName = "nios_dns_view.test_match_destinations"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewMatchDestinations("MATCH_DESTINATIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_destinations", "MATCH_DESTINATIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewMatchDestinations("MATCH_DESTINATIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_destinations", "MATCH_DESTINATIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_MaxCacheTtl(t *testing.T) {
	var resourceName = "nios_dns_view.test_max_cache_ttl"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewMaxCacheTtl("MAX_CACHE_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_cache_ttl", "MAX_CACHE_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewMaxCacheTtl("MAX_CACHE_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_cache_ttl", "MAX_CACHE_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_MaxNcacheTtl(t *testing.T) {
	var resourceName = "nios_dns_view.test_max_ncache_ttl"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewMaxNcacheTtl("MAX_NCACHE_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_ncache_ttl", "MAX_NCACHE_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewMaxNcacheTtl("MAX_NCACHE_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_ncache_ttl", "MAX_NCACHE_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_MaxUdpSize(t *testing.T) {
	var resourceName = "nios_dns_view.test_max_udp_size"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewMaxUdpSize("MAX_UDP_SIZE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_udp_size", "MAX_UDP_SIZE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewMaxUdpSize("MAX_UDP_SIZE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_udp_size", "MAX_UDP_SIZE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Name(t *testing.T) {
	var resourceName = "nios_dns_view.test_name"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_NetworkView(t *testing.T) {
	var resourceName = "nios_dns_view.test_network_view"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewNetworkView("NETWORK_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewNetworkView("NETWORK_VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_view", "NETWORK_VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_NotifyDelay(t *testing.T) {
	var resourceName = "nios_dns_view.test_notify_delay"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewNotifyDelay("NOTIFY_DELAY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notify_delay", "NOTIFY_DELAY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewNotifyDelay("NOTIFY_DELAY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notify_delay", "NOTIFY_DELAY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_NxdomainLogQuery(t *testing.T) {
	var resourceName = "nios_dns_view.test_nxdomain_log_query"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewNxdomainLogQuery("NXDOMAIN_LOG_QUERY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_log_query", "NXDOMAIN_LOG_QUERY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewNxdomainLogQuery("NXDOMAIN_LOG_QUERY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_log_query", "NXDOMAIN_LOG_QUERY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_NxdomainRedirect(t *testing.T) {
	var resourceName = "nios_dns_view.test_nxdomain_redirect"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewNxdomainRedirect("NXDOMAIN_REDIRECT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect", "NXDOMAIN_REDIRECT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewNxdomainRedirect("NXDOMAIN_REDIRECT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect", "NXDOMAIN_REDIRECT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_NxdomainRedirectAddresses(t *testing.T) {
	var resourceName = "nios_dns_view.test_nxdomain_redirect_addresses"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewNxdomainRedirectAddresses("NXDOMAIN_REDIRECT_ADDRESSES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect_addresses", "NXDOMAIN_REDIRECT_ADDRESSES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewNxdomainRedirectAddresses("NXDOMAIN_REDIRECT_ADDRESSES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect_addresses", "NXDOMAIN_REDIRECT_ADDRESSES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_NxdomainRedirectAddressesV6(t *testing.T) {
	var resourceName = "nios_dns_view.test_nxdomain_redirect_addresses_v6"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewNxdomainRedirectAddressesV6("NXDOMAIN_REDIRECT_ADDRESSES_V6_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect_addresses_v6", "NXDOMAIN_REDIRECT_ADDRESSES_V6_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewNxdomainRedirectAddressesV6("NXDOMAIN_REDIRECT_ADDRESSES_V6_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect_addresses_v6", "NXDOMAIN_REDIRECT_ADDRESSES_V6_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_NxdomainRedirectTtl(t *testing.T) {
	var resourceName = "nios_dns_view.test_nxdomain_redirect_ttl"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewNxdomainRedirectTtl("NXDOMAIN_REDIRECT_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect_ttl", "NXDOMAIN_REDIRECT_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewNxdomainRedirectTtl("NXDOMAIN_REDIRECT_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_redirect_ttl", "NXDOMAIN_REDIRECT_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_NxdomainRulesets(t *testing.T) {
	var resourceName = "nios_dns_view.test_nxdomain_rulesets"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewNxdomainRulesets("NXDOMAIN_RULESETS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_rulesets", "NXDOMAIN_RULESETS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewNxdomainRulesets("NXDOMAIN_RULESETS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nxdomain_rulesets", "NXDOMAIN_RULESETS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Recursion(t *testing.T) {
	var resourceName = "nios_dns_view.test_recursion"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewRecursion("RECURSION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recursion", "RECURSION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewRecursion("RECURSION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recursion", "RECURSION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_ResponseRateLimiting(t *testing.T) {
	var resourceName = "nios_dns_view.test_response_rate_limiting"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewResponseRateLimiting("RESPONSE_RATE_LIMITING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "response_rate_limiting", "RESPONSE_RATE_LIMITING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewResponseRateLimiting("RESPONSE_RATE_LIMITING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "response_rate_limiting", "RESPONSE_RATE_LIMITING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_RootNameServerType(t *testing.T) {
	var resourceName = "nios_dns_view.test_root_name_server_type"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewRootNameServerType("ROOT_NAME_SERVER_TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "root_name_server_type", "ROOT_NAME_SERVER_TYPE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewRootNameServerType("ROOT_NAME_SERVER_TYPE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "root_name_server_type", "ROOT_NAME_SERVER_TYPE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_RpzDropIpRuleEnabled(t *testing.T) {
	var resourceName = "nios_dns_view.test_rpz_drop_ip_rule_enabled"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewRpzDropIpRuleEnabled("RPZ_DROP_IP_RULE_ENABLED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_enabled", "RPZ_DROP_IP_RULE_ENABLED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewRpzDropIpRuleEnabled("RPZ_DROP_IP_RULE_ENABLED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_enabled", "RPZ_DROP_IP_RULE_ENABLED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_RpzDropIpRuleMinPrefixLengthIpv4(t *testing.T) {
	var resourceName = "nios_dns_view.test_rpz_drop_ip_rule_min_prefix_length_ipv4"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewRpzDropIpRuleMinPrefixLengthIpv4("RPZ_DROP_IP_RULE_MIN_PREFIX_LENGTH_IPV4_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv4", "RPZ_DROP_IP_RULE_MIN_PREFIX_LENGTH_IPV4_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewRpzDropIpRuleMinPrefixLengthIpv4("RPZ_DROP_IP_RULE_MIN_PREFIX_LENGTH_IPV4_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv4", "RPZ_DROP_IP_RULE_MIN_PREFIX_LENGTH_IPV4_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_RpzDropIpRuleMinPrefixLengthIpv6(t *testing.T) {
	var resourceName = "nios_dns_view.test_rpz_drop_ip_rule_min_prefix_length_ipv6"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewRpzDropIpRuleMinPrefixLengthIpv6("RPZ_DROP_IP_RULE_MIN_PREFIX_LENGTH_IPV6_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv6", "RPZ_DROP_IP_RULE_MIN_PREFIX_LENGTH_IPV6_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewRpzDropIpRuleMinPrefixLengthIpv6("RPZ_DROP_IP_RULE_MIN_PREFIX_LENGTH_IPV6_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv6", "RPZ_DROP_IP_RULE_MIN_PREFIX_LENGTH_IPV6_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_RpzQnameWaitRecurse(t *testing.T) {
	var resourceName = "nios_dns_view.test_rpz_qname_wait_recurse"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewRpzQnameWaitRecurse("RPZ_QNAME_WAIT_RECURSE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_qname_wait_recurse", "RPZ_QNAME_WAIT_RECURSE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewRpzQnameWaitRecurse("RPZ_QNAME_WAIT_RECURSE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_qname_wait_recurse", "RPZ_QNAME_WAIT_RECURSE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_ScavengingSettings(t *testing.T) {
	var resourceName = "nios_dns_view.test_scavenging_settings"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewScavengingSettings("SCAVENGING_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings", "SCAVENGING_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewScavengingSettings("SCAVENGING_SETTINGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings", "SCAVENGING_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Sortlist(t *testing.T) {
	var resourceName = "nios_dns_view.test_sortlist"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewSortlist("SORTLIST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sortlist", "SORTLIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewSortlist("SORTLIST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sortlist", "SORTLIST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseBlacklist(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_blacklist"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseBlacklist("USE_BLACKLIST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blacklist", "USE_BLACKLIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseBlacklist("USE_BLACKLIST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_blacklist", "USE_BLACKLIST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseDdnsForceCreationTimestampUpdate(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_ddns_force_creation_timestamp_update"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseDdnsForceCreationTimestampUpdate("USE_DDNS_FORCE_CREATION_TIMESTAMP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_force_creation_timestamp_update", "USE_DDNS_FORCE_CREATION_TIMESTAMP_UPDATE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseDdnsForceCreationTimestampUpdate("USE_DDNS_FORCE_CREATION_TIMESTAMP_UPDATE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_force_creation_timestamp_update", "USE_DDNS_FORCE_CREATION_TIMESTAMP_UPDATE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseDdnsPatternsRestriction(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_ddns_patterns_restriction"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseDdnsPatternsRestriction("USE_DDNS_PATTERNS_RESTRICTION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_patterns_restriction", "USE_DDNS_PATTERNS_RESTRICTION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseDdnsPatternsRestriction("USE_DDNS_PATTERNS_RESTRICTION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_patterns_restriction", "USE_DDNS_PATTERNS_RESTRICTION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseDdnsPrincipalSecurity(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_ddns_principal_security"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseDdnsPrincipalSecurity("USE_DDNS_PRINCIPAL_SECURITY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_principal_security", "USE_DDNS_PRINCIPAL_SECURITY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseDdnsPrincipalSecurity("USE_DDNS_PRINCIPAL_SECURITY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_principal_security", "USE_DDNS_PRINCIPAL_SECURITY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseDdnsRestrictProtected(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_ddns_restrict_protected"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseDdnsRestrictProtected("USE_DDNS_RESTRICT_PROTECTED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_protected", "USE_DDNS_RESTRICT_PROTECTED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseDdnsRestrictProtected("USE_DDNS_RESTRICT_PROTECTED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_protected", "USE_DDNS_RESTRICT_PROTECTED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseDdnsRestrictStatic(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_ddns_restrict_static"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseDdnsRestrictStatic("USE_DDNS_RESTRICT_STATIC_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_static", "USE_DDNS_RESTRICT_STATIC_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseDdnsRestrictStatic("USE_DDNS_RESTRICT_STATIC_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_static", "USE_DDNS_RESTRICT_STATIC_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseDns64(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_dns64"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseDns64("USE_DNS64_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dns64", "USE_DNS64_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseDns64("USE_DNS64_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dns64", "USE_DNS64_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseDnssec(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_dnssec"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseDnssec("USE_DNSSEC_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dnssec", "USE_DNSSEC_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseDnssec("USE_DNSSEC_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dnssec", "USE_DNSSEC_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseEdnsUdpSize(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_edns_udp_size"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseEdnsUdpSize("USE_EDNS_UDP_SIZE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_edns_udp_size", "USE_EDNS_UDP_SIZE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseEdnsUdpSize("USE_EDNS_UDP_SIZE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_edns_udp_size", "USE_EDNS_UDP_SIZE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseFilterAaaa(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_filter_aaaa"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseFilterAaaa("USE_FILTER_AAAA_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_filter_aaaa", "USE_FILTER_AAAA_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseFilterAaaa("USE_FILTER_AAAA_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_filter_aaaa", "USE_FILTER_AAAA_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseFixedRrsetOrderFqdns(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_fixed_rrset_order_fqdns"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseFixedRrsetOrderFqdns("USE_FIXED_RRSET_ORDER_FQDNS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_fixed_rrset_order_fqdns", "USE_FIXED_RRSET_ORDER_FQDNS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseFixedRrsetOrderFqdns("USE_FIXED_RRSET_ORDER_FQDNS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_fixed_rrset_order_fqdns", "USE_FIXED_RRSET_ORDER_FQDNS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseForwarders(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_forwarders"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseForwarders("USE_FORWARDERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_forwarders", "USE_FORWARDERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseForwarders("USE_FORWARDERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_forwarders", "USE_FORWARDERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseMaxCacheTtl(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_max_cache_ttl"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseMaxCacheTtl("USE_MAX_CACHE_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_max_cache_ttl", "USE_MAX_CACHE_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseMaxCacheTtl("USE_MAX_CACHE_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_max_cache_ttl", "USE_MAX_CACHE_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseMaxNcacheTtl(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_max_ncache_ttl"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseMaxNcacheTtl("USE_MAX_NCACHE_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_max_ncache_ttl", "USE_MAX_NCACHE_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseMaxNcacheTtl("USE_MAX_NCACHE_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_max_ncache_ttl", "USE_MAX_NCACHE_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseMaxUdpSize(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_max_udp_size"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseMaxUdpSize("USE_MAX_UDP_SIZE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_max_udp_size", "USE_MAX_UDP_SIZE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseMaxUdpSize("USE_MAX_UDP_SIZE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_max_udp_size", "USE_MAX_UDP_SIZE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseNxdomainRedirect(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_nxdomain_redirect"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseNxdomainRedirect("USE_NXDOMAIN_REDIRECT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nxdomain_redirect", "USE_NXDOMAIN_REDIRECT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseNxdomainRedirect("USE_NXDOMAIN_REDIRECT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_nxdomain_redirect", "USE_NXDOMAIN_REDIRECT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseRecursion(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_recursion"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseRecursion("USE_RECURSION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recursion", "USE_RECURSION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseRecursion("USE_RECURSION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_recursion", "USE_RECURSION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseResponseRateLimiting(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_response_rate_limiting"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseResponseRateLimiting("USE_RESPONSE_RATE_LIMITING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_response_rate_limiting", "USE_RESPONSE_RATE_LIMITING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseResponseRateLimiting("USE_RESPONSE_RATE_LIMITING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_response_rate_limiting", "USE_RESPONSE_RATE_LIMITING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseRootNameServer(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_root_name_server"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseRootNameServer("USE_ROOT_NAME_SERVER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_root_name_server", "USE_ROOT_NAME_SERVER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseRootNameServer("USE_ROOT_NAME_SERVER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_root_name_server", "USE_ROOT_NAME_SERVER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseRpzDropIpRule(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_rpz_drop_ip_rule"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseRpzDropIpRule("USE_RPZ_DROP_IP_RULE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_rpz_drop_ip_rule", "USE_RPZ_DROP_IP_RULE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseRpzDropIpRule("USE_RPZ_DROP_IP_RULE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_rpz_drop_ip_rule", "USE_RPZ_DROP_IP_RULE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseRpzQnameWaitRecurse(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_rpz_qname_wait_recurse"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseRpzQnameWaitRecurse("USE_RPZ_QNAME_WAIT_RECURSE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_rpz_qname_wait_recurse", "USE_RPZ_QNAME_WAIT_RECURSE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseRpzQnameWaitRecurse("USE_RPZ_QNAME_WAIT_RECURSE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_rpz_qname_wait_recurse", "USE_RPZ_QNAME_WAIT_RECURSE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseScavengingSettings(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_scavenging_settings"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseScavengingSettings("USE_SCAVENGING_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_scavenging_settings", "USE_SCAVENGING_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseScavengingSettings("USE_SCAVENGING_SETTINGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_scavenging_settings", "USE_SCAVENGING_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseSortlist(t *testing.T) {
	var resourceName = "nios_dns_view.test_use_sortlist"
	var v dns.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseSortlist("USE_SORTLIST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_sortlist", "USE_SORTLIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseSortlist("USE_SORTLIST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_sortlist", "USE_SORTLIST_UPDATE_REPLACE_ME"),
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

func testAccViewBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_dns_view" "test" {
}
`)
}

func testAccViewRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccViewBlacklistAction(blacklistAction string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_blacklist_action" {
    blacklist_action = %q
}
`, blacklistAction)
}

func testAccViewBlacklistLogQuery(blacklistLogQuery string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_blacklist_log_query" {
    blacklist_log_query = %q
}
`, blacklistLogQuery)
}

func testAccViewBlacklistRedirectAddresses(blacklistRedirectAddresses string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_blacklist_redirect_addresses" {
    blacklist_redirect_addresses = %q
}
`, blacklistRedirectAddresses)
}

func testAccViewBlacklistRedirectTtl(blacklistRedirectTtl string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_blacklist_redirect_ttl" {
    blacklist_redirect_ttl = %q
}
`, blacklistRedirectTtl)
}

func testAccViewBlacklistRulesets(blacklistRulesets string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_blacklist_rulesets" {
    blacklist_rulesets = %q
}
`, blacklistRulesets)
}

func testAccViewCloudInfo(cloudInfo string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_cloud_info" {
    cloud_info = %q
}
`, cloudInfo)
}

func testAccViewComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccViewCustomRootNameServers(customRootNameServers string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_custom_root_name_servers" {
    custom_root_name_servers = %q
}
`, customRootNameServers)
}

func testAccViewDdnsForceCreationTimestampUpdate(ddnsForceCreationTimestampUpdate string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_ddns_force_creation_timestamp_update" {
    ddns_force_creation_timestamp_update = %q
}
`, ddnsForceCreationTimestampUpdate)
}

func testAccViewDdnsPrincipalGroup(ddnsPrincipalGroup string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_ddns_principal_group" {
    ddns_principal_group = %q
}
`, ddnsPrincipalGroup)
}

func testAccViewDdnsPrincipalTracking(ddnsPrincipalTracking string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_ddns_principal_tracking" {
    ddns_principal_tracking = %q
}
`, ddnsPrincipalTracking)
}

func testAccViewDdnsRestrictPatterns(ddnsRestrictPatterns string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_ddns_restrict_patterns" {
    ddns_restrict_patterns = %q
}
`, ddnsRestrictPatterns)
}

func testAccViewDdnsRestrictPatternsList(ddnsRestrictPatternsList string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_ddns_restrict_patterns_list" {
    ddns_restrict_patterns_list = %q
}
`, ddnsRestrictPatternsList)
}

func testAccViewDdnsRestrictProtected(ddnsRestrictProtected string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_ddns_restrict_protected" {
    ddns_restrict_protected = %q
}
`, ddnsRestrictProtected)
}

func testAccViewDdnsRestrictSecure(ddnsRestrictSecure string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_ddns_restrict_secure" {
    ddns_restrict_secure = %q
}
`, ddnsRestrictSecure)
}

func testAccViewDdnsRestrictStatic(ddnsRestrictStatic string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_ddns_restrict_static" {
    ddns_restrict_static = %q
}
`, ddnsRestrictStatic)
}

func testAccViewDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccViewDns64Enabled(dns64Enabled string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_dns64_enabled" {
    dns64_enabled = %q
}
`, dns64Enabled)
}

func testAccViewDns64Groups(dns64Groups string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_dns64_groups" {
    dns64_groups = %q
}
`, dns64Groups)
}

func testAccViewDnssecEnabled(dnssecEnabled string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_dnssec_enabled" {
    dnssec_enabled = %q
}
`, dnssecEnabled)
}

func testAccViewDnssecExpiredSignaturesEnabled(dnssecExpiredSignaturesEnabled string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_dnssec_expired_signatures_enabled" {
    dnssec_expired_signatures_enabled = %q
}
`, dnssecExpiredSignaturesEnabled)
}

func testAccViewDnssecNegativeTrustAnchors(dnssecNegativeTrustAnchors string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_dnssec_negative_trust_anchors" {
    dnssec_negative_trust_anchors = %q
}
`, dnssecNegativeTrustAnchors)
}

func testAccViewDnssecTrustedKeys(dnssecTrustedKeys string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_dnssec_trusted_keys" {
    dnssec_trusted_keys = %q
}
`, dnssecTrustedKeys)
}

func testAccViewDnssecValidationEnabled(dnssecValidationEnabled string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_dnssec_validation_enabled" {
    dnssec_validation_enabled = %q
}
`, dnssecValidationEnabled)
}

func testAccViewEdnsUdpSize(ednsUdpSize string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_edns_udp_size" {
    edns_udp_size = %q
}
`, ednsUdpSize)
}

func testAccViewEnableBlacklist(enableBlacklist string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_enable_blacklist" {
    enable_blacklist = %q
}
`, enableBlacklist)
}

func testAccViewEnableFixedRrsetOrderFqdns(enableFixedRrsetOrderFqdns string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_enable_fixed_rrset_order_fqdns" {
    enable_fixed_rrset_order_fqdns = %q
}
`, enableFixedRrsetOrderFqdns)
}

func testAccViewEnableMatchRecursiveOnly(enableMatchRecursiveOnly string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_enable_match_recursive_only" {
    enable_match_recursive_only = %q
}
`, enableMatchRecursiveOnly)
}

func testAccViewExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccViewFilterAaaa(filterAaaa string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_filter_aaaa" {
    filter_aaaa = %q
}
`, filterAaaa)
}

func testAccViewFilterAaaaList(filterAaaaList string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_filter_aaaa_list" {
    filter_aaaa_list = %q
}
`, filterAaaaList)
}

func testAccViewFixedRrsetOrderFqdns(fixedRrsetOrderFqdns string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_fixed_rrset_order_fqdns" {
    fixed_rrset_order_fqdns = %q
}
`, fixedRrsetOrderFqdns)
}

func testAccViewForwardOnly(forwardOnly string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_forward_only" {
    forward_only = %q
}
`, forwardOnly)
}

func testAccViewForwarders(forwarders string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_forwarders" {
    forwarders = %q
}
`, forwarders)
}

func testAccViewLastQueriedAcl(lastQueriedAcl string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_last_queried_acl" {
    last_queried_acl = %q
}
`, lastQueriedAcl)
}

func testAccViewMatchClients(matchClients string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_match_clients" {
    match_clients = %q
}
`, matchClients)
}

func testAccViewMatchDestinations(matchDestinations string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_match_destinations" {
    match_destinations = %q
}
`, matchDestinations)
}

func testAccViewMaxCacheTtl(maxCacheTtl string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_max_cache_ttl" {
    max_cache_ttl = %q
}
`, maxCacheTtl)
}

func testAccViewMaxNcacheTtl(maxNcacheTtl string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_max_ncache_ttl" {
    max_ncache_ttl = %q
}
`, maxNcacheTtl)
}

func testAccViewMaxUdpSize(maxUdpSize string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_max_udp_size" {
    max_udp_size = %q
}
`, maxUdpSize)
}

func testAccViewName(name string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_name" {
    name = %q
}
`, name)
}

func testAccViewNetworkView(networkView string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_network_view" {
    network_view = %q
}
`, networkView)
}

func testAccViewNotifyDelay(notifyDelay string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_notify_delay" {
    notify_delay = %q
}
`, notifyDelay)
}

func testAccViewNxdomainLogQuery(nxdomainLogQuery string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_nxdomain_log_query" {
    nxdomain_log_query = %q
}
`, nxdomainLogQuery)
}

func testAccViewNxdomainRedirect(nxdomainRedirect string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_nxdomain_redirect" {
    nxdomain_redirect = %q
}
`, nxdomainRedirect)
}

func testAccViewNxdomainRedirectAddresses(nxdomainRedirectAddresses string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_nxdomain_redirect_addresses" {
    nxdomain_redirect_addresses = %q
}
`, nxdomainRedirectAddresses)
}

func testAccViewNxdomainRedirectAddressesV6(nxdomainRedirectAddressesV6 string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_nxdomain_redirect_addresses_v6" {
    nxdomain_redirect_addresses_v6 = %q
}
`, nxdomainRedirectAddressesV6)
}

func testAccViewNxdomainRedirectTtl(nxdomainRedirectTtl string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_nxdomain_redirect_ttl" {
    nxdomain_redirect_ttl = %q
}
`, nxdomainRedirectTtl)
}

func testAccViewNxdomainRulesets(nxdomainRulesets string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_nxdomain_rulesets" {
    nxdomain_rulesets = %q
}
`, nxdomainRulesets)
}

func testAccViewRecursion(recursion string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_recursion" {
    recursion = %q
}
`, recursion)
}

func testAccViewResponseRateLimiting(responseRateLimiting string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_response_rate_limiting" {
    response_rate_limiting = %q
}
`, responseRateLimiting)
}

func testAccViewRootNameServerType(rootNameServerType string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_root_name_server_type" {
    root_name_server_type = %q
}
`, rootNameServerType)
}

func testAccViewRpzDropIpRuleEnabled(rpzDropIpRuleEnabled string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_rpz_drop_ip_rule_enabled" {
    rpz_drop_ip_rule_enabled = %q
}
`, rpzDropIpRuleEnabled)
}

func testAccViewRpzDropIpRuleMinPrefixLengthIpv4(rpzDropIpRuleMinPrefixLengthIpv4 string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_rpz_drop_ip_rule_min_prefix_length_ipv4" {
    rpz_drop_ip_rule_min_prefix_length_ipv4 = %q
}
`, rpzDropIpRuleMinPrefixLengthIpv4)
}

func testAccViewRpzDropIpRuleMinPrefixLengthIpv6(rpzDropIpRuleMinPrefixLengthIpv6 string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_rpz_drop_ip_rule_min_prefix_length_ipv6" {
    rpz_drop_ip_rule_min_prefix_length_ipv6 = %q
}
`, rpzDropIpRuleMinPrefixLengthIpv6)
}

func testAccViewRpzQnameWaitRecurse(rpzQnameWaitRecurse string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_rpz_qname_wait_recurse" {
    rpz_qname_wait_recurse = %q
}
`, rpzQnameWaitRecurse)
}

func testAccViewScavengingSettings(scavengingSettings string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_scavenging_settings" {
    scavenging_settings = %q
}
`, scavengingSettings)
}

func testAccViewSortlist(sortlist string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_sortlist" {
    sortlist = %q
}
`, sortlist)
}

func testAccViewUseBlacklist(useBlacklist string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_blacklist" {
    use_blacklist = %q
}
`, useBlacklist)
}

func testAccViewUseDdnsForceCreationTimestampUpdate(useDdnsForceCreationTimestampUpdate string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_ddns_force_creation_timestamp_update" {
    use_ddns_force_creation_timestamp_update = %q
}
`, useDdnsForceCreationTimestampUpdate)
}

func testAccViewUseDdnsPatternsRestriction(useDdnsPatternsRestriction string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_ddns_patterns_restriction" {
    use_ddns_patterns_restriction = %q
}
`, useDdnsPatternsRestriction)
}

func testAccViewUseDdnsPrincipalSecurity(useDdnsPrincipalSecurity string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_ddns_principal_security" {
    use_ddns_principal_security = %q
}
`, useDdnsPrincipalSecurity)
}

func testAccViewUseDdnsRestrictProtected(useDdnsRestrictProtected string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_ddns_restrict_protected" {
    use_ddns_restrict_protected = %q
}
`, useDdnsRestrictProtected)
}

func testAccViewUseDdnsRestrictStatic(useDdnsRestrictStatic string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_ddns_restrict_static" {
    use_ddns_restrict_static = %q
}
`, useDdnsRestrictStatic)
}

func testAccViewUseDns64(useDns64 string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_dns64" {
    use_dns64 = %q
}
`, useDns64)
}

func testAccViewUseDnssec(useDnssec string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_dnssec" {
    use_dnssec = %q
}
`, useDnssec)
}

func testAccViewUseEdnsUdpSize(useEdnsUdpSize string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_edns_udp_size" {
    use_edns_udp_size = %q
}
`, useEdnsUdpSize)
}

func testAccViewUseFilterAaaa(useFilterAaaa string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_filter_aaaa" {
    use_filter_aaaa = %q
}
`, useFilterAaaa)
}

func testAccViewUseFixedRrsetOrderFqdns(useFixedRrsetOrderFqdns string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_fixed_rrset_order_fqdns" {
    use_fixed_rrset_order_fqdns = %q
}
`, useFixedRrsetOrderFqdns)
}

func testAccViewUseForwarders(useForwarders string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_forwarders" {
    use_forwarders = %q
}
`, useForwarders)
}

func testAccViewUseMaxCacheTtl(useMaxCacheTtl string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_max_cache_ttl" {
    use_max_cache_ttl = %q
}
`, useMaxCacheTtl)
}

func testAccViewUseMaxNcacheTtl(useMaxNcacheTtl string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_max_ncache_ttl" {
    use_max_ncache_ttl = %q
}
`, useMaxNcacheTtl)
}

func testAccViewUseMaxUdpSize(useMaxUdpSize string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_max_udp_size" {
    use_max_udp_size = %q
}
`, useMaxUdpSize)
}

func testAccViewUseNxdomainRedirect(useNxdomainRedirect string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_nxdomain_redirect" {
    use_nxdomain_redirect = %q
}
`, useNxdomainRedirect)
}

func testAccViewUseRecursion(useRecursion string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_recursion" {
    use_recursion = %q
}
`, useRecursion)
}

func testAccViewUseResponseRateLimiting(useResponseRateLimiting string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_response_rate_limiting" {
    use_response_rate_limiting = %q
}
`, useResponseRateLimiting)
}

func testAccViewUseRootNameServer(useRootNameServer string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_root_name_server" {
    use_root_name_server = %q
}
`, useRootNameServer)
}

func testAccViewUseRpzDropIpRule(useRpzDropIpRule string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_rpz_drop_ip_rule" {
    use_rpz_drop_ip_rule = %q
}
`, useRpzDropIpRule)
}

func testAccViewUseRpzQnameWaitRecurse(useRpzQnameWaitRecurse string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_rpz_qname_wait_recurse" {
    use_rpz_qname_wait_recurse = %q
}
`, useRpzQnameWaitRecurse)
}

func testAccViewUseScavengingSettings(useScavengingSettings string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_scavenging_settings" {
    use_scavenging_settings = %q
}
`, useScavengingSettings)
}

func testAccViewUseSortlist(useSortlist string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test_use_sortlist" {
    use_sortlist = %q
}
`, useSortlist)
}
