package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccRecordCnameDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_cname.test"
	resourceName := "nios_dns_record_cname.test"
	var v dns.RecordCname
	canonical := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	name := acctest.RandomNameWithPrefix("test-cname") + ".example.com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordCnameDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordCnameDataSourceConfigFilters(canonical, name, "default"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordCnameResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordCnameDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_record_cname.test"
	resourceName := "nios_dns_record_cname.test"
	var v dns.RecordCname
	canonical := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	name := acctest.RandomNameWithPrefix("test-cname") + ".example.com"
	extAttrs := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordCnameDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordCnameDataSourceConfigExtAttrFilters(canonical, name, "default", extAttrs),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordCnameExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordCnameResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordCnameResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "aws_rte53_record_info", dataSourceName, "result.0.aws_rte53_record_info"),
		resource.TestCheckResourceAttrPair(resourceName, "canonical", dataSourceName, "result.0.canonical"),
		resource.TestCheckResourceAttrPair(resourceName, "cloud_info", dataSourceName, "result.0.cloud_info"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "creation_time", dataSourceName, "result.0.creation_time"),
		resource.TestCheckResourceAttrPair(resourceName, "creator", dataSourceName, "result.0.creator"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_principal", dataSourceName, "result.0.ddns_principal"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_protected", dataSourceName, "result.0.ddns_protected"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_canonical", dataSourceName, "result.0.dns_canonical"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name", dataSourceName, "result.0.dns_name"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "forbid_reclamation", dataSourceName, "result.0.forbid_reclamation"),
		resource.TestCheckResourceAttrPair(resourceName, "last_queried", dataSourceName, "result.0.last_queried"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "reclaimable", dataSourceName, "result.0.reclaimable"),
		resource.TestCheckResourceAttrPair(resourceName, "shared_record_group", dataSourceName, "result.0.shared_record_group"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordCnameDataSourceConfigFilters(canonical, name, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test" {
	canonical = %q
	name = %q
	view = %q
}

data "nios_dns_record_cname" "test" {
  filters = {
	name = nios_dns_record_cname.test.name
  }
}
`, canonical, name, view)
}

func testAccRecordCnameDataSourceConfigExtAttrFilters(canonical, name, view, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_record_cname" "test" {
	canonical = %q
	name      = %q
	view      = %q
	extattrs = {
		Site = %q
	} 
}

data "nios_dns_record_cname" "test" {
  extattrfilters = {
	Site = nios_dns_record_cname.test.extattrs.Site
  }
}
`, canonical, name, view, extAttrsValue)
}
