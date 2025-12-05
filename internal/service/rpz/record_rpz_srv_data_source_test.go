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

func TestAccRecordRpzSrvDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_srv.test"
	resourceName := "nios_rpz_record_srv.test"
	var v rpz.RecordRpzSrv
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	target := acctest.RandomName() + ".target.com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzSrvDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzSrvDataSourceConfigFilters(name, rpZone, target, 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzSrvResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordRpzSrvDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_rpz_record_srv.test"
	resourceName := "nios_rpz_record_srv.test"
	var v rpz.RecordRpzSrv
	rpZone := acctest.RandomNameWithPrefix("test-zone") + ".com"
	name := acctest.RandomName() + "." + rpZone
	target := acctest.RandomName() + ".target.com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordRpzSrvDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordRpzSrvDataSourceConfigExtAttrFilters(name, rpZone, target, acctest.RandomName(), 80, 10, 360),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordRpzSrvExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordRpzSrvResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRecordRpzSrvResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "port", dataSourceName, "result.0.port"),
		resource.TestCheckResourceAttrPair(resourceName, "priority", dataSourceName, "result.0.priority"),
		resource.TestCheckResourceAttrPair(resourceName, "rp_zone", dataSourceName, "result.0.rp_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "target", dataSourceName, "result.0.target"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "result.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "weight", dataSourceName, "result.0.weight"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "result.0.zone"),
	}
}

func testAccRecordRpzSrvDataSourceConfigFilters(name, rpZone, target string, port, priority, weight int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_srv" "test" {
	name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	port = %d
	priority = %d
	weight = %d
	target = %q
}

data "nios_rpz_record_srv" "test" {
  filters = {
	name = nios_rpz_record_srv.test.name
  }
}
`, name, port, priority, weight, target)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}

func testAccRecordRpzSrvDataSourceConfigExtAttrFilters(name, rpZone, target, extAttrsValue string, port, priority, weight int) string {
	config := fmt.Sprintf(`
resource "nios_rpz_record_srv" "test" {
	name = %q
	rp_zone = nios_dns_zone_rp.test.fqdn
	port = %d
	priority = %d
	weight = %d
	target = %q
	extattrs = {
		Site = %q
	} 
}

data "nios_rpz_record_srv" "test" {
  extattrfilters = {
	Site = nios_rpz_record_srv.test.extattrs.Site
  }
}
`, name, port, priority, weight, target, extAttrsValue)

	return strings.Join([]string{testAccBaseWithZone(rpZone, ""), config}, "")
}
