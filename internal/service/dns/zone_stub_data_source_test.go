package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccZoneStubDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_zone_stub.test"
	resourceName := "nios_dns_zone_stub.test"
	var v dns.ZoneStub

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneStubDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneStubDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					}, testAccCheckZoneStubResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccZoneStubDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_zone_stub.test"
	resourceName := "nios_dns_zone_stub.test"
	var v dns.ZoneStub
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneStubDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneStubDataSourceConfigExtAttrFilters(acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckZoneStubExists(context.Background(), resourceName, &v),
					}, testAccCheckZoneStubResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckZoneStubResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "result.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "disable_forwarding", dataSourceName, "result.0.disable_forwarding"),
		resource.TestCheckResourceAttrPair(resourceName, "display_domain", dataSourceName, "result.0.display_domain"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_fqdn", dataSourceName, "result.0.dns_fqdn"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "external_ns_group", dataSourceName, "result.0.external_ns_group"),
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
		resource.TestCheckResourceAttrPair(resourceName, "soa_email", dataSourceName, "result.0.soa_email"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_expire", dataSourceName, "result.0.soa_expire"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_mname", dataSourceName, "result.0.soa_mname"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_negative_ttl", dataSourceName, "result.0.soa_negative_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "soarefresh", dataSourceName, "result.0.soarefresh"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_retry", dataSourceName, "result.0.soa_retry"),
		resource.TestCheckResourceAttrPair(resourceName, "soa_serial_number", dataSourceName, "result.0.soa_serial_number"),
		resource.TestCheckResourceAttrPair(resourceName, "stub_from", dataSourceName, "result.0.stub_from"),
		resource.TestCheckResourceAttrPair(resourceName, "stub_members", dataSourceName, "result.0.stub_members"),
		resource.TestCheckResourceAttrPair(resourceName, "stub_msservers", dataSourceName, "result.0.stub_msservers"),
		resource.TestCheckResourceAttrPair(resourceName, "using_srg_associations", dataSourceName, "result.0.using_srg_associations"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone_format", dataSourceName, "result.0.zone_format"),
	}
}

func testAccZoneStubDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test" {
}

data "nios_dns_zone_stub" "test" {
  filters = {
	 = nios_dns_zone_stub.test.
  }
}
`)
}

func testAccZoneStubDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_stub" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_dns_zone_stub" "test" {
  extattrfilters = {
	Site = nios_dns_zone_stub.test.extattrs.Site
  }
}
`, extAttrsValue)
}
