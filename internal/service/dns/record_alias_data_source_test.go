package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccRecordAliasDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_alias.test"
	resourceName := "nios_dns_record_alias.test"
	var v dns.RecordAlias
	name := acctest.RandomName() + ".example.com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordAliasDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordAliasDataSourceConfigFilters(name, "server.example.com", "A", "default"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordAliasResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordAliasDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_alias.test"
	resourceName := "nios_dns_record_alias.test"
	var v dns.RecordAlias
	name := acctest.RandomName() + ".example.com"
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordAliasDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordAliasDataSourceConfigExtAttrFilters(name, "server.example.com", "A", "default", extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordAliasExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordAliasResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordAliasResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "aws_rte53_record_info", dataSourceName, "result.0.aws_rte53_record_info"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "creator", dataSourceName, "result.0.creator"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name", dataSourceName, "result.0.dns_name"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_target_name", dataSourceName, "result.0.dns_target_name"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "last_queried", dataSourceName, "result.0.last_queried"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "target_name", dataSourceName, "result.0.target_name"),
		resource.TestCheckResourceAttrPair(resourceName, "target_type", dataSourceName, "result.0.target_type"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordAliasDataSourceConfigFilters(name, target_name, target_type, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_alias" "test" {
    name 		= %q
	target_name = %q
	target_type = %q
	view 		= %q
}

data "nios_dns_record_alias" "test" {
  filters = {
	name = nios_dns_record_alias.test.name
  }
}
`, name, target_name, target_type, view)
}

func testAccRecordAliasDataSourceConfigExtAttrFilters(name, target_name, target_type, view, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_alias" "test" {
	name 		= %q
	target_name = %q
	target_type = %q
	view 		= %q
	extattrs = {
		Site 	= %q
	} 
}

data "nios_dns_record_alias" "test" {
  extattrfilters = {
	Site = nios_dns_record_alias.test.extattrs.Site
  }
}
`, name, target_name, target_type, view, extAttrsValue)
}
