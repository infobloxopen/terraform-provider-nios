package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccSharedrecordgroupDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_sharedrecordgroup.test"
	resourceName := "nios_dns_sharedrecordgroup.test"
	var v dns.Sharedrecordgroup
	name := acctest.RandomNameWithPrefix("sharedrecordgroup")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordgroupDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSharedrecordgroupExists(context.Background(), resourceName, &v),
					}, testAccCheckSharedrecordgroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccSharedrecordgroupDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_sharedrecordgroup.test"
	resourceName := "nios_dns_sharedrecordgroup.test"
	name := acctest.RandomNameWithPrefix("sharedrecordgroup")
	var v dns.Sharedrecordgroup
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSharedrecordgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSharedrecordgroupDataSourceConfigExtAttrFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSharedrecordgroupExists(context.Background(), resourceName, &v),
					}, testAccCheckSharedrecordgroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckSharedrecordgroupResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "record_name_policy", dataSourceName, "result.0.record_name_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "use_record_name_policy", dataSourceName, "result.0.use_record_name_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "zone_associations", dataSourceName, "result.0.zone_associations"),
	}
}

func testAccSharedrecordgroupDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecordgroup" "test" {
  name = %q
}

data "nios_dns_sharedrecordgroup" "test" {
  filters = {
	 name = nios_dns_sharedrecordgroup.test.name
  }
}
`, name)
}

func testAccSharedrecordgroupDataSourceConfigExtAttrFilters(name, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_sharedrecordgroup" "test" {
  name = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_dns_sharedrecordgroup" "test" {
  extattrfilters = {
	Site = nios_dns_sharedrecordgroup.test.extattrs.Site
  }
}
`, name, extAttrsValue)
}
