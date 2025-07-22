package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccZoneDelegatedDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_zone_delegated.test"
	resourceName := "nios_dns_zone_delegated.test"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneDelegatedDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneDelegatedDataSourceConfigFilters(fqdn, delegatedToName, "10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					}, testAccCheckZoneDelegatedResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccZoneDelegatedDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_zone_delegated.test"
	resourceName := "nios_dns_zone_delegated.test"
	var v dns.ZoneDelegated
	fqdn := acctest.RandomNameWithPrefix("zone-delegated") + ".example.com"
	delegatedToName := acctest.RandomNameWithPrefix("zone-delegated") + ".com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneDelegatedDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneDelegatedDataSourceConfigExtAttrFilters(fqdn, delegatedToName, "10.0.0.1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckZoneDelegatedExists(context.Background(), resourceName, &v),
					}, testAccCheckZoneDelegatedResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckZoneDelegatedResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "result.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "delegate_to", dataSourceName, "result.0.delegate_to"),
		resource.TestCheckResourceAttrPair(resourceName, "delegated_ttl", dataSourceName, "result.0.delegated_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "display_domain", dataSourceName, "result.0.display_domain"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_fqdn", dataSourceName, "result.0.dns_fqdn"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_rfc2317_exclusion", dataSourceName, "result.0.enable_rfc2317_exclusion"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "fqdn", dataSourceName, "result.0.fqdn"),
		resource.TestCheckResourceAttrPair(resourceName, "locked", dataSourceName, "result.0.locked"),
		resource.TestCheckResourceAttrPair(resourceName, "locked_by", dataSourceName, "result.0.locked_by"),
		resource.TestCheckResourceAttrPair(resourceName, "mask_prefix", dataSourceName, "result.0.mask_prefix"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_ad_integrated", dataSourceName, "result.0.ms_ad_integrated"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_ddns_mode", dataSourceName, "result.0.ms_ddns_mode"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_managed", dataSourceName, "result.0.ms_managed"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_read_only", dataSourceName, "result.0.ms_read_only"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_sync_master_name", dataSourceName, "result.0.ms_sync_master_name"),
		resource.TestCheckResourceAttrPair(resourceName, "ns_group", dataSourceName, "result.0.ns_group"),
		resource.TestCheckResourceAttrPair(resourceName, "parent", dataSourceName, "result.0.parent"),
		resource.TestCheckResourceAttrPair(resourceName, "prefix", dataSourceName, "result.0.prefix"),
		resource.TestCheckResourceAttrPair(resourceName, "use_delegated_ttl", dataSourceName, "result.0.use_delegated_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "using_srg_associations", dataSourceName, "result.0.using_srg_associations"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone_format", dataSourceName, "result.0.zone_format"),
	}
}

func testAccZoneDelegatedDataSourceConfigFilters(fqdn, delegateToName, delegateToAddress string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_delegated" "test" {
	fqdn = %q
    delegate_to = [
		{
			name = %q
			address = %q
		}
	]
}

data "nios_dns_zone_delegated" "test" {
  filters = {
	fqdn = nios_dns_zone_delegated.test.fqdn
  }
}
`, fqdn, delegateToName, delegateToAddress)
}

func testAccZoneDelegatedDataSourceConfigExtAttrFilters(fqdn, delegateToName, delegateToAddress string, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_delegated" "test" {
	fqdn = %q
    delegate_to = [
		{
			name = %q
			address = %q
		}
	]
	extattrs = {
		Site = %q
	} 
}

data "nios_dns_zone_delegated" "test" {
  extattrfilters = {
	Site = nios_dns_zone_delegated.test.extattrs.Site
  }
}
`, fqdn, delegateToName, delegateToAddress, extAttrsValue)
}
