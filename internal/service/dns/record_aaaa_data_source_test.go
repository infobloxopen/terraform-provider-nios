package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/acctest"
)

func TestAccRecordAaaaDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_aaaa.test"
	resourceName := "nios_dns_record_aaaa.test"
	var v dns.RecordAaaa

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordAaaaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordAaaaDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordAaaaResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordAaaaDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_aaaa.test"
	resourceName := "nios_dns_record_aaaa.test"
	var v dns.RecordAaaa
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordAaaaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordAaaaDataSourceConfigExtAttrFilters("value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordAaaaExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordAaaaResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordAaaaResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
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
		resource.TestCheckResourceAttrPair(resourceName, "ipv6addr", dataSourceName, "result.0.ipv6addr"),
		resource.TestCheckResourceAttrPair(resourceName, "func_call", dataSourceName, "result.0.func_call"),
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

func testAccRecordAaaaDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test" {
}

data "nios_dns_record_aaaa" "test" {
  filters = {
	 = nios_dns_record_aaaa.test.
  }
}
`)
}

func testAccRecordAaaaDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_aaaa" "test" {
  extattrs = {
    Site = {
        value = %q
    }
  	}
}

data "nios_dns_record_aaaa" "test" {
  extattrfilters = {
	"Site" = nios_dns_record_aaaa.test.tags.tag1
  }
}
`, extAttrsValue)
}
