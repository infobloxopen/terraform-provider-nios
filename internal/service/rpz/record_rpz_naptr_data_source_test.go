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

func TestAccRecordRpzNaptrDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_naptr.test"
	resourceName := "nios_rpz_record_naptr.test"
	var v rpz.RecordRpzNaptr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzNaptrDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzNaptrDataSourceConfigFilters(name, rpZone, ".", 10, 10),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzNaptrResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordRpzNaptrDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_naptr.test"
	resourceName := "nios_rpz_record_naptr.test"
	var v rpz.RecordRpzNaptr
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzNaptrDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzNaptrDataSourceConfigExtAttrFilters(name, rpZone, ".", acctest.RandomName(), 10, 10),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzNaptrExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzNaptrResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordRpzNaptrResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "flags", dataSourceName, "result.0.flags"),
		resource.TestCheckResourceAttrPair(resourceName, "last_queried", dataSourceName, "result.0.last_queried"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "order", dataSourceName, "result.0.order"),
		resource.TestCheckResourceAttrPair(resourceName, "preference", dataSourceName, "result.0.preference"),
		resource.TestCheckResourceAttrPair(resourceName, "regexp", dataSourceName, "result.0.regexp"),
		resource.TestCheckResourceAttrPair(resourceName, "replacement", dataSourceName, "result.0.replacement"),
		resource.TestCheckResourceAttrPair(resourceName, "rp_zone", dataSourceName, "result.0.rp_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "services", dataSourceName, "result.0.services"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordRpzNaptrDataSourceConfigFilters(name, rpZone, replacement string, order, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_naptr" "test" {
	name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	order = %d
    preference = %d
    replacement = %q
}

data "nios_rpz_record_naptr" "test" {
  filters = {
	name = nios_rpz_record_naptr.test.name
  }
}
`, name, order, preference, replacement)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzNaptrDataSourceConfigExtAttrFilters(name, rpZone, replacement, extAttrsValue string, order, preference int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_naptr" "test" {
	name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	order = %d
    preference = %d
    replacement = %q
	extattrs = {
		Site = %q
	} 
}

data "nios_rpz_record_naptr" "test" {
  extattrfilters = {
	Site = nios_rpz_record_naptr.test.extattrs.Site
  }
}
`, name, order, preference, replacement, extAttrsValue)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}
