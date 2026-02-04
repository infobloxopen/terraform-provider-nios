package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccZoneAuthDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_zone_auth.test"
	resourceName := "nios_dns_zone_auth.test"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("auth-zone") + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneAuthDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneAuthDataSourceConfigFilters(zoneFqdn, "default"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					}, testAccCheckZoneAuthResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccZoneAuthDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_zone_auth.test"
	resourceName := "nios_dns_zone_auth.test"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("auth-zone") + ".com"
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneAuthDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneAuthDataSourceConfigExtAttrFilters(zoneFqdn, "default", extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					}, testAccCheckZoneAuthResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckZoneAuthResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "result.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "allow_active_dir", dataSourceName, "result.0.allow_active_dir"),
		resource.TestCheckResourceAttrPair(resourceName, "allow_fixed_rrset_order", dataSourceName, "result.0.allow_fixed_rrset_order"),
		resource.TestCheckResourceAttrPair(resourceName, "allow_gss_tsig_for_underscore_zone", dataSourceName, "result.0.allow_gss_tsig_for_underscore_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "allow_gss_tsig_zone_updates", dataSourceName, "result.0.allow_gss_tsig_zone_updates"),
		resource.TestCheckResourceAttrPair(resourceName, "allow_query", dataSourceName, "result.0.allow_query"),
		resource.TestCheckResourceAttrPair(resourceName, "allow_transfer", dataSourceName, "result.0.allow_transfer"),
		resource.TestCheckResourceAttrPair(resourceName, "allow_update", dataSourceName, "result.0.allow_update"),
		resource.TestCheckResourceAttrPair(resourceName, "allow_update_forwarding", dataSourceName, "result.0.allow_update_forwarding"),
		resource.TestCheckResourceAttrPair(resourceName, "aws_rte53_zone_info", dataSourceName, "result.0.aws_rte53_zone_info"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "copy_xfer_to_notify", dataSourceName, "result.0.copy_xfer_to_notify"),
		resource.TestCheckResourceAttrPair(resourceName, "create_underscore_zones", dataSourceName, "result.0.create_underscore_zones"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_force_creation_timestamp_update", dataSourceName, "result.0.ddns_force_creation_timestamp_update"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_principal_group", dataSourceName, "result.0.ddns_principal_group"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_principal_tracking", dataSourceName, "result.0.ddns_principal_tracking"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_restrict_patterns", dataSourceName, "result.0.ddns_restrict_patterns"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_restrict_patterns_list", dataSourceName, "result.0.ddns_restrict_patterns_list"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_restrict_protected", dataSourceName, "result.0.ddns_restrict_protected"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_restrict_secure", dataSourceName, "result.0.ddns_restrict_secure"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_restrict_static", dataSourceName, "result.0.ddns_restrict_static"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "disable_forwarding", dataSourceName, "result.0.disable_forwarding"),
		resource.TestCheckResourceAttrPair(resourceName, "display_domain", dataSourceName, "result.0.display_domain"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_fqdn", dataSourceName, "result.0.dns_fqdn"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_integrity_enable", dataSourceName, "result.0.dns_integrity_enable"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_integrity_frequency", dataSourceName, "result.0.dns_integrity_frequency"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_integrity_member", dataSourceName, "result.0.dns_integrity_member"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_integrity_verbose_logging", dataSourceName, "result.0.dns_integrity_verbose_logging"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_soa_email", dataSourceName, "result.0.dns_soa_email"),
		resource.TestCheckResourceAttrPair(resourceName, "dnssec_key_params", dataSourceName, "result.0.dnssec_key_params"),
		resource.TestCheckResourceAttrPair(resourceName, "dnssec_keys", dataSourceName, "result.0.dnssec_keys"),
		resource.TestCheckResourceAttrPair(resourceName, "dnssec_ksk_rollover_date", dataSourceName, "result.0.dnssec_ksk_rollover_date"),
		resource.TestCheckResourceAttrPair(resourceName, "dnssec_zsk_rollover_date", dataSourceName, "result.0.dnssec_zsk_rollover_date"),
		resource.TestCheckResourceAttrPair(resourceName, "effective_check_names_policy", dataSourceName, "result.0.effective_check_names_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "effective_record_name_policy", dataSourceName, "result.0.effective_record_name_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "external_primaries", dataSourceName, "result.0.external_primaries"),
		resource.TestCheckResourceAttrPair(resourceName, "external_secondaries", dataSourceName, "result.0.external_secondaries"),
		resource.TestCheckResourceAttrPair(resourceName, "fqdn", dataSourceName, "result.0.fqdn"),
		resource.TestCheckResourceAttrPair(resourceName, "grid_primary", dataSourceName, "result.0.grid_primary"),
		resource.TestCheckResourceAttrPair(resourceName, "grid_primary_shared_with_ms_parent_delegation", dataSourceName, "result.0.grid_primary_shared_with_ms_parent_delegation"),
		resource.TestCheckResourceAttrPair(resourceName, "grid_secondaries", dataSourceName, "result.0.grid_secondaries"),
		resource.TestCheckResourceAttrPair(resourceName, "is_dnssec_enabled", dataSourceName, "result.0.is_dnssec_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "is_dnssec_signed", dataSourceName, "result.0.is_dnssec_signed"),
		resource.TestCheckResourceAttrPair(resourceName, "is_multimaster", dataSourceName, "result.0.is_multimaster"),
		resource.TestCheckResourceAttrPair(resourceName, "last_queried", dataSourceName, "result.0.last_queried"),
		resource.TestCheckResourceAttrPair(resourceName, "last_queried_acl", dataSourceName, "result.0.last_queried_acl"),
		resource.TestCheckResourceAttrPair(resourceName, "locked", dataSourceName, "result.0.locked"),
		resource.TestCheckResourceAttrPair(resourceName, "locked_by", dataSourceName, "result.0.locked_by"),
		resource.TestCheckResourceAttrPair(resourceName, "mask_prefix", dataSourceName, "result.0.mask_prefix"),
		resource.TestCheckResourceAttrPair(resourceName, "member_soa_mnames", dataSourceName, "result.0.member_soa_mnames"),
		resource.TestCheckResourceAttrPair(resourceName, "member_soa_serials", dataSourceName, "result.0.member_soa_serials"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_ad_integrated", dataSourceName, "result.0.ms_ad_integrated"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_allow_transfer", dataSourceName, "result.0.ms_allow_transfer"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_allow_transfer_mode", dataSourceName, "result.0.ms_allow_transfer_mode"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_dc_ns_record_creation", dataSourceName, "result.0.ms_dc_ns_record_creation"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_ddns_mode", dataSourceName, "result.0.ms_ddns_mode"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_managed", dataSourceName, "result.0.ms_managed"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_primaries", dataSourceName, "result.0.ms_primaries"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_read_only", dataSourceName, "result.0.ms_read_only"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_secondaries", dataSourceName, "result.0.ms_secondaries"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_sync_disabled", dataSourceName, "result.0.ms_sync_disabled"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_sync_master_name", dataSourceName, "result.0.ms_sync_master_name"),
		resource.TestCheckResourceAttrPair(resourceName, "network_associations", dataSourceName, "result.0.network_associations"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "notify_delay", dataSourceName, "result.0.notify_delay"),
		resource.TestCheckResourceAttrPair(resourceName, "ns_group", dataSourceName, "result.0.ns_group"),
		resource.TestCheckResourceAttrPair(resourceName, "parent", dataSourceName, "result.0.parent"),
		resource.TestCheckResourceAttrPair(resourceName, "prefix", dataSourceName, "result.0.prefix"),
		resource.TestCheckResourceAttrPair(resourceName, "primary_type", dataSourceName, "result.0.primary_type"),
		resource.TestCheckResourceAttrPair(resourceName, "record_name_policy", dataSourceName, "result.0.record_name_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "records_monitored", dataSourceName, "result.0.records_monitored"),
		resource.TestCheckResourceAttrPair(resourceName, "remove_subzones", dataSourceName, "result.0.remove_subzones"),
		resource.TestCheckResourceAttrPair(resourceName, "rr_not_queried_enabled_time", dataSourceName, "result.0.rr_not_queried_enabled_time"),
		resource.TestCheckResourceAttrPair(resourceName, "scavenging_settings", dataSourceName, "result.0.scavenging_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_default_ttl", dataSourceName, "result.0.soa_default_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_email", dataSourceName, "result.0.soa_email"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_expire", dataSourceName, "result.0.soa_expire"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_negative_ttl", dataSourceName, "result.0.soa_negative_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_refresh", dataSourceName, "result.0.soa_refresh"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_retry", dataSourceName, "result.0.soa_retry"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_serial_number", dataSourceName, "result.0.soa_serial_number"),
		resource.TestCheckResourceAttrPair(resourceName, "srgs", dataSourceName, "result.0.srgs"),
		resource.TestCheckResourceAttrPair(resourceName, "update_forwarding", dataSourceName, "result.0.update_forwarding"),
		resource.TestCheckResourceAttrPair(resourceName, "use_allow_active_dir", dataSourceName, "result.0.use_allow_active_dir"),
		resource.TestCheckResourceAttrPair(resourceName, "use_allow_query", dataSourceName, "result.0.use_allow_query"),
		resource.TestCheckResourceAttrPair(resourceName, "use_allow_transfer", dataSourceName, "result.0.use_allow_transfer"),
		resource.TestCheckResourceAttrPair(resourceName, "use_allow_update", dataSourceName, "result.0.use_allow_update"),
		resource.TestCheckResourceAttrPair(resourceName, "use_allow_update_forwarding", dataSourceName, "result.0.use_allow_update_forwarding"),
		resource.TestCheckResourceAttrPair(resourceName, "use_check_names_policy", dataSourceName, "result.0.use_check_names_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "use_copy_xfer_to_notify", dataSourceName, "result.0.use_copy_xfer_to_notify"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_force_creation_timestamp_update", dataSourceName, "result.0.use_ddns_force_creation_timestamp_update"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_patterns_restriction", dataSourceName, "result.0.use_ddns_patterns_restriction"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_principal_security", dataSourceName, "result.0.use_ddns_principal_security"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_restrict_protected", dataSourceName, "result.0.use_ddns_restrict_protected"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ddns_restrict_static", dataSourceName, "result.0.use_ddns_restrict_static"),
		resource.TestCheckResourceAttrPair(resourceName, "use_dnssec_key_params", dataSourceName, "result.0.use_dnssec_key_params"),
		resource.TestCheckResourceAttrPair(resourceName, "use_external_primary", dataSourceName, "result.0.use_external_primary"),
		resource.TestCheckResourceAttrPair(resourceName, "use_grid_zone_timer", dataSourceName, "result.0.use_grid_zone_timer"),
		resource.TestCheckResourceAttrPair(resourceName, "use_import_from", dataSourceName, "result.0.use_import_from"),
		resource.TestCheckResourceAttrPair(resourceName, "use_notify_delay", dataSourceName, "result.0.use_notify_delay"),
		resource.TestCheckResourceAttrPair(resourceName, "use_record_name_policy", dataSourceName, "result.0.use_record_name_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "use_scavenging_settings", dataSourceName, "result.0.use_scavenging_settings"),
		resource.TestCheckResourceAttrPair(resourceName, "use_soa_email", dataSourceName, "result.0.use_soa_email"),
		resource.TestCheckResourceAttrPair(resourceName, "using_srg_associations", dataSourceName, "result.0.using_srg_associations"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone_format", dataSourceName, "result.0.zone_format"),
		resource.TestCheckResourceAttrPair(resourceName, "zone_not_queried_enabled_time", dataSourceName, "result.0.zone_not_queried_enabled_time"),
	}
}

func testAccZoneAuthDataSourceConfigFilters(zoneFqdn, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test" {
  fqdn = %q
  view = %q
}

data "nios_dns_zone_auth" "test" {
  filters = {
	fqdn = nios_dns_zone_auth.test.fqdn
	view = nios_dns_zone_auth.test.view
  }
}
`, zoneFqdn, view)
}

func testAccZoneAuthDataSourceConfigExtAttrFilters(zoneFqdn, view, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test" {
	fqdn = %q
	view = %q
   extattrs = {
		Site = %q
	}
}

data "nios_dns_zone_auth" "test" {
  extattrfilters = {
		Site = nios_dns_zone_auth.test.extattrs.Site
}
}
`, zoneFqdn, view, extAttrsValue)
}
