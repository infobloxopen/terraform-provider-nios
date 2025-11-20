package dns_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccSharedrecordSrvDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_sharedrecord_srv.test"
	resourceName := "nios_dns_sharedrecord_srv.test"
	var v dns.SharedrecordSrv
	name := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	target := acctest.RandomName() + ".target.com"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordSrvDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordSrvDataSourceConfigFilters(name, 80, 10, sharedRecordGroup, target, 10),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					}, testAccCheckSharedrecordSrvResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccSharedrecordSrvDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_sharedrecord_srv.test"
	resourceName := "nios_dns_sharedrecord_srv.test"
	name := acctest.RandomNameWithPrefix("sharedrecord-srv") + ".example.com"
	target := acctest.RandomName() + ".target.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	extAttrValue := acctest.RandomName()

	var v dns.SharedrecordSrv
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordSrvDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordSrvDataSourceConfigExtAttrFilters(name, 80, 10, sharedRecordGroup, target, 10, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSharedrecordSrvExists(context.Background(), resourceName, &v),
					}, testAccCheckSharedrecordSrvResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckSharedrecordSrvResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name", dataSourceName, "result.0.dns_name"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_target", dataSourceName, "result.0.dns_target"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "port", dataSourceName, "result.0.port"),
		resource.TestCheckResourceAttrPair(resourceName, "priority", dataSourceName, "result.0.priority"),
		resource.TestCheckResourceAttrPair(resourceName, "shared_record_group", dataSourceName, "result.0.shared_record_group"),
		resource.TestCheckResourceAttrPair(resourceName, "target", dataSourceName, "result.0.target"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "weight", dataSourceName, "result.0.weight"),
	}
}

func testAccSharedrecordSrvDataSourceConfigFilters(name string, port, priority int, sharedRecordGroup, target string, weight int) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_srv" "test" {
  name               = %q
  port               = %d
  priority           = %d
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  target             = %q
  weight             = %d
}

data "nios_dns_sharedrecord_srv" "test" {
  filters = {
    name = nios_dns_sharedrecord_srv.test.name
  }
}
`, name, port, priority, target, weight)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordSrvDataSourceConfigExtAttrFilters(name string, port, priority int, sharedRecordGroup, target string, weight int, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_srv" "test" {
  name               = %q
  port               = %d
  priority           = %d
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  target             = %q
  weight             = %d
  extattrs = {
    Site = %q
  }
}

data "nios_dns_sharedrecord_srv" "test" {
  extattrfilters = {
    Site = nios_dns_sharedrecord_srv.test.extattrs.Site
  }
}
`, name, port, priority, target, weight, extAttrsValue)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}
