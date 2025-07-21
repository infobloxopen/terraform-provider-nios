
package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccViewDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_view.test"
	resourceName := "nios_dns_view.test"
	var v dns.View

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckViewDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccViewDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckViewExists(context.Background(), resourceName, &v),
						}, testAccCheckViewResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccViewDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_view.test"
	resourceName := "nios_dns_view.test"
	var v dns.View
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckViewDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccViewDataSourceConfigExtAttrFilters(, "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckViewExists(context.Background(), resourceName, &v),
						}, testAccCheckViewResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckViewResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "blacklist_action", dataSourceName, "result.0.blacklist_action"),
        resource.TestCheckResourceAttrPair(resourceName, "blacklist_log_query", dataSourceName, "result.0.blacklist_log_query"),
        resource.TestCheckResourceAttrPair(resourceName, "blacklist_redirect_addresses", dataSourceName, "result.0.blacklist_redirect_addresses"),
        resource.TestCheckResourceAttrPair(resourceName, "blacklist_redirect_ttl", dataSourceName, "result.0.blacklist_redirect_ttl"),
        resource.TestCheckResourceAttrPair(resourceName, "blacklist_rulesets", dataSourceName, "result.0.blacklist_rulesets"),
        resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
        resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
        resource.TestCheckResourceAttrPair(resourceName, "custom_root_name_servers", dataSourceName, "result.0.custom_root_name_servers"),
        resource.TestCheckResourceAttrPair(resourceName, "ddns_force_creation_timestamp_update", dataSourceName, "result.0.ddns_force_creation_timestamp_update"),
        resource.TestCheckResourceAttrPair(resourceName, "ddns_principal_group", dataSourceName, "result.0.ddns_principal_group"),
        resource.TestCheckResourceAttrPair(resourceName, "ddns_principal_tracking", dataSourceName, "result.0.ddns_principal_tracking"),
        resource.TestCheckResourceAttrPair(resourceName, "ddns_restrict_patterns", dataSourceName, "result.0.ddns_restrict_patterns"),
        resource.TestCheckResourceAttrPair(resourceName, "ddns_restrict_patterns_list", dataSourceName, "result.0.ddns_restrict_patterns_list"),
        resource.TestCheckResourceAttrPair(resourceName, "ddns_restrict_protected", dataSourceName, "result.0.ddns_restrict_protected"),
        resource.TestCheckResourceAttrPair(resourceName, "ddns_restrict_secure", dataSourceName, "result.0.ddns_restrict_secure"),
        resource.TestCheckResourceAttrPair(resourceName, "ddns_restrict_static", dataSourceName, "result.0.ddns_restrict_static"),
        resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
        resource.TestCheckResourceAttrPair(resourceName, "dns64_enabled", dataSourceName, "result.0.dns64_enabled"),
        resource.TestCheckResourceAttrPair(resourceName, "dns64_groups", dataSourceName, "result.0.dns64_groups"),
        resource.TestCheckResourceAttrPair(resourceName, "dnssec_enabled", dataSourceName, "result.0.dnssec_enabled"),
        resource.TestCheckResourceAttrPair(resourceName, "dnssec_expired_signatures_enabled", dataSourceName, "result.0.dnssec_expired_signatures_enabled"),
        resource.TestCheckResourceAttrPair(resourceName, "dnssec_negative_trust_anchors", dataSourceName, "result.0.dnssec_negative_trust_anchors"),
        resource.TestCheckResourceAttrPair(resourceName, "dnssec_trusted_keys", dataSourceName, "result.0.dnssec_trusted_keys"),
        resource.TestCheckResourceAttrPair(resourceName, "dnssec_validation_enabled", dataSourceName, "result.0.dnssec_validation_enabled"),
        resource.TestCheckResourceAttrPair(resourceName, "edns_udp_size", dataSourceName, "result.0.edns_udp_size"),
        resource.TestCheckResourceAttrPair(resourceName, "enable_blacklist", dataSourceName, "result.0.enable_blacklist"),
        resource.TestCheckResourceAttrPair(resourceName, "enable_fixed_rrset_order_fqdns", dataSourceName, "result.0.enable_fixed_rrset_order_fqdns"),
        resource.TestCheckResourceAttrPair(resourceName, "enable_match_recursive_only", dataSourceName, "result.0.enable_match_recursive_only"),
        resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
        resource.TestCheckResourceAttrPair(resourceName, "filter_aaaa", dataSourceName, "result.0.filter_aaaa"),
        resource.TestCheckResourceAttrPair(resourceName, "filter_aaaa_list", dataSourceName, "result.0.filter_aaaa_list"),
        resource.TestCheckResourceAttrPair(resourceName, "fixed_rrset_order_fqdns", dataSourceName, "result.0.fixed_rrset_order_fqdns"),
        resource.TestCheckResourceAttrPair(resourceName, "forward_only", dataSourceName, "result.0.forward_only"),
        resource.TestCheckResourceAttrPair(resourceName, "forwarders", dataSourceName, "result.0.forwarders"),
        resource.TestCheckResourceAttrPair(resourceName, "is_default", dataSourceName, "result.0.is_default"),
        resource.TestCheckResourceAttrPair(resourceName, "last_queried_acl", dataSourceName, "result.0.last_queried_acl"),
        resource.TestCheckResourceAttrPair(resourceName, "match_clients", dataSourceName, "result.0.match_clients"),
        resource.TestCheckResourceAttrPair(resourceName, "match_destinations", dataSourceName, "result.0.match_destinations"),
        resource.TestCheckResourceAttrPair(resourceName, "max_cache_ttl", dataSourceName, "result.0.max_cache_ttl"),
        resource.TestCheckResourceAttrPair(resourceName, "max_ncache_ttl", dataSourceName, "result.0.max_ncache_ttl"),
        resource.TestCheckResourceAttrPair(resourceName, "max_udp_size", dataSourceName, "result.0.max_udp_size"),
        resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
        resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
        resource.TestCheckResourceAttrPair(resourceName, "notify_delay", dataSourceName, "result.0.notify_delay"),
        resource.TestCheckResourceAttrPair(resourceName, "nxdomain_log_query", dataSourceName, "result.0.nxdomain_log_query"),
        resource.TestCheckResourceAttrPair(resourceName, "nxdomain_redirect", dataSourceName, "result.0.nxdomain_redirect"),
        resource.TestCheckResourceAttrPair(resourceName, "nxdomain_redirect_addresses", dataSourceName, "result.0.nxdomain_redirect_addresses"),
        resource.TestCheckResourceAttrPair(resourceName, "nxdomain_redirect_addresses_v6", dataSourceName, "result.0.nxdomain_redirect_addresses_v6"),
        resource.TestCheckResourceAttrPair(resourceName, "nxdomain_redirect_ttl", dataSourceName, "result.0.nxdomain_redirect_ttl"),
        resource.TestCheckResourceAttrPair(resourceName, "nxdomain_rulesets", dataSourceName, "result.0.nxdomain_rulesets"),
        resource.TestCheckResourceAttrPair(resourceName, "recursion", dataSourceName, "result.0.recursion"),
        resource.TestCheckResourceAttrPair(resourceName, "response_rate_limiting", dataSourceName, "result.0.response_rate_limiting"),
        resource.TestCheckResourceAttrPair(resourceName, "root_name_server_type", dataSourceName, "result.0.root_name_server_type"),
        resource.TestCheckResourceAttrPair(resourceName, "rpz_drop_ip_rule_enabled", dataSourceName, "result.0.rpz_drop_ip_rule_enabled"),
        resource.TestCheckResourceAttrPair(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv4", dataSourceName, "result.0.rpz_drop_ip_rule_min_prefix_length_ipv4"),
        resource.TestCheckResourceAttrPair(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv6", dataSourceName, "result.0.rpz_drop_ip_rule_min_prefix_length_ipv6"),
        resource.TestCheckResourceAttrPair(resourceName, "rpz_qname_wait_recurse", dataSourceName, "result.0.rpz_qname_wait_recurse"),
        resource.TestCheckResourceAttrPair(resourceName, "scavenging_settings", dataSourceName, "result.0.scavenging_settings"),
        resource.TestCheckResourceAttrPair(resourceName, "sortlist", dataSourceName, "result.0.sortlist"),
        resource.TestCheckResourceAttrPair(resourceName, "use_blacklist", dataSourceName, "result.0.use_blacklist"),
        resource.TestCheckResourceAttrPair(resourceName, "use_ddns_force_creation_timestamp_update", dataSourceName, "result.0.use_ddns_force_creation_timestamp_update"),
        resource.TestCheckResourceAttrPair(resourceName, "use_ddns_patterns_restriction", dataSourceName, "result.0.use_ddns_patterns_restriction"),
        resource.TestCheckResourceAttrPair(resourceName, "use_ddns_principal_security", dataSourceName, "result.0.use_ddns_principal_security"),
        resource.TestCheckResourceAttrPair(resourceName, "use_ddns_restrict_protected", dataSourceName, "result.0.use_ddns_restrict_protected"),
        resource.TestCheckResourceAttrPair(resourceName, "use_ddns_restrict_static", dataSourceName, "result.0.use_ddns_restrict_static"),
        resource.TestCheckResourceAttrPair(resourceName, "use_dns64", dataSourceName, "result.0.use_dns64"),
        resource.TestCheckResourceAttrPair(resourceName, "use_dnssec", dataSourceName, "result.0.use_dnssec"),
        resource.TestCheckResourceAttrPair(resourceName, "use_edns_udp_size", dataSourceName, "result.0.use_edns_udp_size"),
        resource.TestCheckResourceAttrPair(resourceName, "use_filter_aaaa", dataSourceName, "result.0.use_filter_aaaa"),
        resource.TestCheckResourceAttrPair(resourceName, "use_fixed_rrset_order_fqdns", dataSourceName, "result.0.use_fixed_rrset_order_fqdns"),
        resource.TestCheckResourceAttrPair(resourceName, "use_forwarders", dataSourceName, "result.0.use_forwarders"),
        resource.TestCheckResourceAttrPair(resourceName, "use_max_cache_ttl", dataSourceName, "result.0.use_max_cache_ttl"),
        resource.TestCheckResourceAttrPair(resourceName, "use_max_ncache_ttl", dataSourceName, "result.0.use_max_ncache_ttl"),
        resource.TestCheckResourceAttrPair(resourceName, "use_max_udp_size", dataSourceName, "result.0.use_max_udp_size"),
        resource.TestCheckResourceAttrPair(resourceName, "use_nxdomain_redirect", dataSourceName, "result.0.use_nxdomain_redirect"),
        resource.TestCheckResourceAttrPair(resourceName, "use_recursion", dataSourceName, "result.0.use_recursion"),
        resource.TestCheckResourceAttrPair(resourceName, "use_response_rate_limiting", dataSourceName, "result.0.use_response_rate_limiting"),
        resource.TestCheckResourceAttrPair(resourceName, "use_root_name_server", dataSourceName, "result.0.use_root_name_server"),
        resource.TestCheckResourceAttrPair(resourceName, "use_rpz_drop_ip_rule", dataSourceName, "result.0.use_rpz_drop_ip_rule"),
        resource.TestCheckResourceAttrPair(resourceName, "use_rpz_qname_wait_recurse", dataSourceName, "result.0.use_rpz_qname_wait_recurse"),
        resource.TestCheckResourceAttrPair(resourceName, "use_scavenging_settings", dataSourceName, "result.0.use_scavenging_settings"),
        resource.TestCheckResourceAttrPair(resourceName, "use_sortlist", dataSourceName, "result.0.use_sortlist"),
    }
}

func testAccViewDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test" {
}

data "nios_dns_view" "test" {
  filters = {
	 = nios_dns_view.test.
  }
}
`)
}

func testAccViewDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_view" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_dns_view" "test" {
  extattrfilters = {
	"Site" = nios_dns_view.test.extattrs.Site
  }
}
`,extAttrsValue)
}

