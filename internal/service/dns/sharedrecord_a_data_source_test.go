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

func TestAccSharedrecordADataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_sharedrecord_a.test"
	resourceName := "nios_dns_sharedrecord_a.test"
	var v dns.SharedrecordA
	name := acctest.RandomNameWithPrefix("sharedrecord-a")
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordADestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordADataSourceConfigFilters(name, "10.0.0.0", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					}, testAccCheckSharedrecordAResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccSharedrecordADataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_sharedrecord_a.test"
	resourceName := "nios_dns_sharedrecord_a.test"
	name := acctest.RandomNameWithPrefix("sharedrecord-a") + ".example.com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	extAttrValue := acctest.RandomName()

	var v dns.SharedrecordA
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordADestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordADataSourceConfigExtAttrFilters(name, "10.0.0.0", sharedRecordGroup, extAttrValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSharedrecordAExists(context.Background(), resourceName, &v),
					}, testAccCheckSharedrecordAResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckSharedrecordAResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name", dataSourceName, "result.0.dns_name"),
		resource.TestCheckResourceAttrPair(resourceName, "ipv4addr", dataSourceName, "result.0.ipv4addr"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "shared_record_group", dataSourceName, "result.0.shared_record_group"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
	}
}

func testAccSharedrecordADataSourceConfigFilters(name, ipv4addr, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_a" "test" {
	name                = %q
	ipv4addr            = %q
	shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}

data "nios_dns_sharedrecord_a" "test" {
	filters = {
		name = nios_dns_sharedrecord_a.test.name
	}
}
`, name, ipv4addr)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordADataSourceConfigExtAttrFilters(name, ipv4addr, sharedRecordGroup, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_a" "test" {
    name                = %q
    ipv4addr            = %q
    shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
    extattrs = {
        Site = %q
    }
}

data "nios_dns_sharedrecord_a" "test" {
    extattrfilters = {
        Site = nios_dns_sharedrecord_a.test.extattrs.Site
    }
}
`, name, ipv4addr, extAttrsValue)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}
