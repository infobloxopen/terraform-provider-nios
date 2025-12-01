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

func TestAccRecordRpzMxDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_mx.test"
	resourceName := "nios_rpz_record_mx.test"
	var v rpz.RecordRpzMx
	rpZone := acctest.RandomNameWithPrefix("rpz") + ".example.com"
	name := acctest.RandomName() + "." + rpZone
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzMxDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzMxDataSourceConfigFilters(name, mailExchanger, rpZone, 10),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzMxResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordRpzMxDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_mx.test"
	resourceName := "nios_rpz_record_mx.test"
	var v rpz.RecordRpzMx
	rpZone := acctest.RandomNameWithPrefix("rpz") + ".example.com"
	name := acctest.RandomName() + "." + rpZone
	mailExchanger := acctest.RandomNameWithPrefix("mail-exchanger") + ".example.com"
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzMxDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzMxDataSourceConfigExtAttrFilters(name, mailExchanger, rpZone, 10, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzMxExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzMxResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordRpzMxResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "mail_exchanger", dataSourceName, "result.0.mail_exchanger"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "preference", dataSourceName, "result.0.preference"),
		resource.TestCheckResourceAttrPair(resourceName, "rp_zone", dataSourceName, "result.0.rp_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordRpzMxDataSourceConfigFilters(name, mailExchanger, rpZone string, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_mx" "test" {
	name            = %q
	mail_exchanger  = %q
	rp_zone         = nios_dns_zone_rp.test.fqdn
	preference      = %d
}

data "nios_rpz_record_mx" "test" {
	filters = {
		name = nios_rpz_record_mx.test.name
	}
}
`, name, mailExchanger, preference)
	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzMxDataSourceConfigExtAttrFilters(name, mailExchanger, rpZone string, preference int, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_mx" "test" {
	name            = %q
	mail_exchanger  = %q
	rp_zone         = nios_dns_zone_rp.test.fqdn
	preference      = %d
	extattrs = {
		Site = %q
	} 
}

data "nios_rpz_record_mx" "test" {
	extattrfilters = {
		Site = nios_rpz_record_mx.test.extattrs.Site
	}
}
`, name, mailExchanger, preference, extAttrsValue)
	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")

}
