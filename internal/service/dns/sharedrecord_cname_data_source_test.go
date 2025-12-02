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

func TestAccSharedrecordCnameDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_sharedrecord_cname.test"
	resourceName := "nios_dns_sharedrecord_cname.test"
	var v dns.SharedrecordCname
	name := acctest.RandomNameWithPrefix("sharedrecord-cname-")
	canonical := acctest.RandomNameWithPrefix("canonical-name") + ".com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordCnameDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordCnameDataSourceConfigFilters(name, canonical, sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					}, testAccCheckSharedrecordCnameResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccSharedrecordCnameDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_sharedrecord_cname.test"
	resourceName := "nios_dns_sharedrecord_cname.test"
	var v dns.SharedrecordCname
	name := acctest.RandomNameWithPrefix("sharedrecord-cname-")
	canonical := acctest.RandomNameWithPrefix("canonical-name") + ".com"
	sharedRecordGroup := acctest.RandomNameWithPrefix("sharedrecordgroup-")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordCnameDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordCnameDataSourceConfigExtAttrFilters(name, canonical, sharedRecordGroup, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSharedrecordCnameExists(context.Background(), resourceName, &v),
					}, testAccCheckSharedrecordCnameResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckSharedrecordCnameResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "canonical", dataSourceName, "result.0.canonical"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_canonical", dataSourceName, "result.0.dns_canonical"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name", dataSourceName, "result.0.dns_name"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "shared_record_group", dataSourceName, "result.0.shared_record_group"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "result.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "use_ttl", dataSourceName, "result.0.use_ttl"),
	}
}

func testAccSharedrecordCnameDataSourceConfigFilters(name, canonical, sharedRecordGroup string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test" {
  name = %q
  canonical = %q
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
}

data "nios_dns_sharedrecord_cname" "test" {
  filters = {
	 name = nios_dns_sharedrecord_cname.test.name
  }
}
`, name, canonical)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}

func testAccSharedrecordCnameDataSourceConfigExtAttrFilters(name, canonical, sharedRecordGroup, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_dns_sharedrecord_cname" "test" {
  name = %q
  canonical = %q
  shared_record_group = nios_dns_sharedrecordgroup.parent_sharedrecord_group.name
  extattrs = {
    Site = %q
  } 
}

data "nios_dns_sharedrecord_cname" "test" {
  extattrfilters = {
	Site = nios_dns_sharedrecord_cname.test.extattrs.Site
  }
}
`, name, canonical, extAttrsValue)
	return strings.Join([]string{testAccBaseSharedRecordGroup(sharedRecordGroup), config}, "")
}
