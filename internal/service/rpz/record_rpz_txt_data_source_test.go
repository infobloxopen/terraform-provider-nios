package rpz_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/rpz"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccRecordRpzTxtDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_txt.test"
	resourceName := "nios_rpz_record_txt.test"
	var v rpz.RecordRpzTxt
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	txt := "Record Text"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzTxtDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzTxtDataSourceConfigFilters(name, rpZone, txt),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzTxtResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordRpzTxtDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_txt.test"
	resourceName := "nios_rpz_record_txt.test"
	var v rpz.RecordRpzTxt
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	txt := "Record Text"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzTxtDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzTxtDataSourceConfigExtAttrFilters(name, rpZone, txt, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzTxtExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzTxtResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordRpzTxtResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "rp_zone", dataSourceName, "result.0.rp_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "text", dataSourceName, "result.0.text"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordRpzTxtDataSourceConfigFilters(name, rpZone, txt string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_txt" "test" {
	name = %q
	text = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}

data "nios_rpz_record_txt" "test" {
  filters = {
	name = nios_rpz_record_txt.test.name
  }
}
`, name, txt)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzTxtDataSourceConfigExtAttrFilters(name, rpZone, txt, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_txt" "test" {
	name = %q
	text = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	extattrs = {
		Site = %q
	} 
}

data "nios_rpz_record_txt" "test" {
  extattrfilters = {
	Site = nios_rpz_record_txt.test.extattrs.Site
  }
}
`, name, txt, extAttrsValue)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}
