package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccZoneRpDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_zone_rp.test"
	resourceName := "nios_dns_zone_rp.test"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneRpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneRpDataSourceConfigFilters(zoneFqdn, "default"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					}, testAccCheckZoneRpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccZoneRpDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_zone_rp.test"
	resourceName := "nios_dns_zone_rp.test"
	var v dns.ZoneRp
	zoneFqdn := acctest.RandomNameWithPrefix("zone-rp") + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneRpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneRpDataSourceConfigExtAttrFilters(zoneFqdn, "default", acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckZoneRpExists(context.Background(), resourceName, &v),
					}, testAccCheckZoneRpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckZoneRpResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "result.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "display_domain", dataSourceName, "result.0.display_domain"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_soa_email", dataSourceName, "result.0.dns_soa_email"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "external_primaries", dataSourceName, "result.0.external_primaries"),
		resource.TestCheckResourceAttrPair(resourceName, "external_secondaries", dataSourceName, "result.0.external_secondaries"),
		resource.TestCheckResourceAttrPair(resourceName, "fireeye_rule_mapping", dataSourceName, "result.0.fireeye_rule_mapping"),
		resource.TestCheckResourceAttrPair(resourceName, "fqdn", dataSourceName, "result.0.fqdn"),
		resource.TestCheckResourceAttrPair(resourceName, "grid_primary", dataSourceName, "result.0.grid_primary"),
		resource.TestCheckResourceAttrPair(resourceName, "grid_secondaries", dataSourceName, "result.0.grid_secondaries"),
		resource.TestCheckResourceAttrPair(resourceName, "locked", dataSourceName, "result.0.locked"),
		resource.TestCheckResourceAttrPair(resourceName, "locked_by", dataSourceName, "result.0.locked_by"),
		resource.TestCheckResourceAttrPair(resourceName, "log_rpz", dataSourceName, "result.0.log_rpz"),
		resource.TestCheckResourceAttrPair(resourceName, "mask_prefix", dataSourceName, "result.0.mask_prefix"),
		resource.TestCheckResourceAttrPair(resourceName, "member_soa_mnames", dataSourceName, "result.0.member_soa_mnames"),
		resource.TestCheckResourceAttrPair(resourceName, "member_soa_serials", dataSourceName, "result.0.member_soa_serials"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "ns_group", dataSourceName, "result.0.ns_group"),
		resource.TestCheckResourceAttrPair(resourceName, "parent", dataSourceName, "result.0.parent"),
		resource.TestCheckResourceAttrPair(resourceName, "prefix", dataSourceName, "result.0.prefix"),
		resource.TestCheckResourceAttrPair(resourceName, "primary_type", dataSourceName, "result.0.primary_type"),
		resource.TestCheckResourceAttrPair(resourceName, "record_name_policy", dataSourceName, "result.0.record_name_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "rpz_drop_ip_rule_enabled", dataSourceName, "result.0.rpz_drop_ip_rule_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv4", dataSourceName, "result.0.rpz_drop_ip_rule_min_prefix_length_ipv4"),
		resource.TestCheckResourceAttrPair(resourceName, "rpz_drop_ip_rule_min_prefix_length_ipv6", dataSourceName, "result.0.rpz_drop_ip_rule_min_prefix_length_ipv6"),
		resource.TestCheckResourceAttrPair(resourceName, "rpz_last_updated_time", dataSourceName, "result.0.rpz_last_updated_time"),
		resource.TestCheckResourceAttrPair(resourceName, "rpz_policy", dataSourceName, "result.0.rpz_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "rpz_priority", dataSourceName, "result.0.rpz_priority"),
		resource.TestCheckResourceAttrPair(resourceName, "rpz_priority_end", dataSourceName, "result.0.rpz_priority_end"),
		resource.TestCheckResourceAttrPair(resourceName, "rpz_severity", dataSourceName, "result.0.rpz_severity"),
		resource.TestCheckResourceAttrPair(resourceName, "rpz_type", dataSourceName, "result.0.rpz_type"),
		resource.TestCheckResourceAttrPair(resourceName, "set_soa_serial_number", dataSourceName, "result.0.set_soa_serial_number"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_default_ttl", dataSourceName, "result.0.soa_default_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_email", dataSourceName, "result.0.soa_email"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_expire", dataSourceName, "result.0.soa_expire"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_negative_ttl", dataSourceName, "result.0.soa_negative_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_refresh", dataSourceName, "result.0.soa_refresh"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_retry", dataSourceName, "result.0.soa_retry"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_serial_number", dataSourceName, "result.0.soa_serial_number"),
		resource.TestCheckResourceAttrPair(resourceName, "substitute_name", dataSourceName, "result.0.substitute_name"),
		resource.TestCheckResourceAttrPair(resourceName, "use_external_primary", dataSourceName, "result.0.use_external_primary"),
		resource.TestCheckResourceAttrPair(resourceName, "use_grid_zone_timer", dataSourceName, "result.0.use_grid_zone_timer"),
		resource.TestCheckResourceAttrPair(resourceName, "use_log_rpz", dataSourceName, "result.0.use_log_rpz"),
		resource.TestCheckResourceAttrPair(resourceName, "use_record_name_policy", dataSourceName, "result.0.use_record_name_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "use_rpz_drop_ip_rule", dataSourceName, "result.0.use_rpz_drop_ip_rule"),
		resource.TestCheckResourceAttrPair(resourceName, "use_soa_email", dataSourceName, "result.0.use_soa_email"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
	}
}

func testAccZoneRpDataSourceConfigFilters(zoneFqdn, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test" {
	fqdn = %q
	view = %q
}

data "nios_dns_zone_rp" "test" {
  filters = {
	fqdn = nios_dns_zone_rp.test.fqdn
	view = nios_dns_zone_rp.test.view
  }
}
`, zoneFqdn, view)
}

func testAccZoneRpDataSourceConfigExtAttrFilters(zoneFqdn, view, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_rp" "test" {
	fqdn = %q
	view = %q
	extattrs = {
		Site = %q
	}
}

data "nios_dns_zone_rp" "test" {
	extattrfilters = {
		Site = nios_dns_zone_rp.test.extattrs.Site
	}
}
`, zoneFqdn, view, extAttrsValue)
}
