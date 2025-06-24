
package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/acctest"
)

func TestAccZoneForwardDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_zone_forward.test"
	resourceName := "nios_dns_zone_forward.test"
	var v dns.ZoneForward

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneForwardDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneForwardDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
						}, testAccCheckZoneForwardResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccZoneForwardDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_zone_forward.test"
	resourceName := "nios_dns_zone_forward.test"
	var v dns.ZoneForward
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneForwardDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneForwardDataSourceConfigExtAttrFilters(, "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckZoneForwardExists(context.Background(), resourceName, &v),
						}, testAccCheckZoneForwardResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckZoneForwardResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "_ref", dataSourceName, "result.0._ref"),
        resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "result.0.address"),
        resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
        resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
        resource.TestCheckResourceAttrPair(resourceName, "disable_ns_generation", dataSourceName, "result.0.disable_ns_generation"),
        resource.TestCheckResourceAttrPair(resourceName, "display_domain", dataSourceName, "result.0.display_domain"),
        resource.TestCheckResourceAttrPair(resourceName, "dns_fqdn", dataSourceName, "result.0.dns_fqdn"),
        resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
        resource.TestCheckResourceAttrPair(resourceName, "external_ns_group", dataSourceName, "result.0.external_ns_group"),
        resource.TestCheckResourceAttrPair(resourceName, "forward_to", dataSourceName, "result.0.forward_to"),
        resource.TestCheckResourceAttrPair(resourceName, "forwarders_only", dataSourceName, "result.0.forwarders_only"),
        resource.TestCheckResourceAttrPair(resourceName, "forwarding_servers", dataSourceName, "result.0.forwarding_servers"),
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
        resource.TestCheckResourceAttrPair(resourceName, "using_srg_associations", dataSourceName, "result.0.using_srg_associations"),
        resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
        resource.TestCheckResourceAttrPair(resourceName, "zone_format", dataSourceName, "result.0.zone_format"),
    }
}

func testAccZoneForwardDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test" {
}

data "nios_dns_zone_forward" "test" {
  filters = {
	 = nios_dns_zone_forward.test.
  }
}
`)
}

func testAccZoneForwardDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_forward" "test" {
  extattrs = {
    Site = {
        value = %q
    }
  	}
}

data "nios_dns_zone_forward" "test" {
  extattrfilters = {
	"Site" = nios_dns_zone_forward.test.tags.tag1
  }
}
`,extAttrsValue)
}

