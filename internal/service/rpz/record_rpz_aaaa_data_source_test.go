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

func TestAccRecordRpzAaaaDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_aaaa.test"
	resourceName := "nios_rpz_record_aaaa.test"
	var v rpz.RecordRpzAaaa
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzAaaaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzAaaaDataSourceConfigFilters(name, "2002:1f93::12:1", rpZone),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzAaaaResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordRpzAaaaDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_aaaa.test"
	resourceName := "nios_rpz_record_aaaa.test"
	var v rpz.RecordRpzAaaa
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	extAttrValue := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzAaaaDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzAaaaDataSourceConfigExtAttrFilters(name, "2002:1f93::12:1", rpZone, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzAaaaExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzAaaaResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordRpzAaaaResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv6addr", dataSourceName, "result.0.ipv6addr"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "rp_zone", dataSourceName, "result.0.rp_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordRpzAaaaDataSourceConfigFilters(name, ipV6Addr, rpZone string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa" "test" {
	name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
}

data "nios_rpz_record_aaaa" "test" {
  filters = {
	name = nios_rpz_record_aaaa.test.name
  }
}
`, name, ipV6Addr)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzAaaaDataSourceConfigExtAttrFilters(name, ipV6Addr, rpZone, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_aaaa" "test" {
	name = %q
	ipv6addr = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	extattrs = {
		Site = %q
	} 
}

data "nios_rpz_record_aaaa" "test" {
  extattrfilters = {
	Site = nios_rpz_record_aaaa.test.extattrs.Site
  }
}
`, name, ipV6Addr, extAttrsValue)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}
