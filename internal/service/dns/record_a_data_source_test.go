package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/acctest"
)

func TestAccRecordaDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_a.test"
	resourceName := "nios_dns_record_a.test"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordaDataSourceConfigFilters(name, "10.0.0.20", "default"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordaExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordaResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordaDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.nios_datasource_nios_RecordA.test"
	resourceName := "nios_dns_record_a.test"
	var v dns.RecordA
	name := acctest.RandomName() + ".example.com"
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordaDataSourceConfigExtAttrFilters(name, "10.0.0.20", "default", extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordaExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordaResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordaResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "_ref", dataSourceName, "result.0._ref"),
		resource.TestCheckResourceAttrPair(resourceName, "aws_rte53_record_info", dataSourceName, "result.0.aws_rte53_record_info"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "creation_time", dataSourceName, "result.0.creation_time"),
		resource.TestCheckResourceAttrPair(resourceName, "creator", dataSourceName, "result.0.creator"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_principal", dataSourceName, "result.0.ddns_principal"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_protected", dataSourceName, "result.0.ddns_protected"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "discovered_data", dataSourceName, "result.0.discovered_data"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name", dataSourceName, "result.0.dns_name"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "forbid_reclamation", dataSourceName, "result.0.forbid_reclamation"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv4addr", dataSourceName, "result.0.ipv4addr"),
		resource.TestCheckResourceAttrPair(resourceName, "last_queried", dataSourceName, "result.0.last_queried"),
		resource.TestCheckResourceAttrPair(resourceName, "ms_ad_user_data", dataSourceName, "result.0.ms_ad_user_data"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "reclaimable", dataSourceName, "result.0.reclaimable"),
		resource.TestCheckResourceAttrPair(resourceName, "remove_associated_ptr", dataSourceName, "result.0.remove_associated_ptr"),
		resource.TestCheckResourceAttrPair(resourceName, "shared_record_group", dataSourceName, "result.0.shared_record_group"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordaDataSourceConfigFilters(name, ipV4Addr, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_a" "test" {
	name = %q
	ipv4addr = %q
	view = %q
}

data "nios_dns_record_a" "test" {
  filters = {
	name = nios_dns_record_a.test.name
  }
}
`, name, ipV4Addr, view)
}

func testAccRecordaDataSourceConfigExtAttrFilters(name, ipV4Addr, view, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_a" "test" {
	name = %q
	ipv4addr = %q
	view = %q
	extattrs = {
		Site = {
			value = %q
		}
	}
}

data "nios_dns_record_a" "test" {
	extattrfilters = {
		"Site" = nios_dns_record_a.test.extattrs.Site.value
	}
}
`, name, ipV4Addr, view, extAttrsValue)
}
