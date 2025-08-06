
package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccNsgroupDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dns_nsgroup.test"
	resourceName := "nios_dns_nsgroup.test"
	var v dns.Nsgroup

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckNsgroupExists(context.Background(), resourceName, &v),
						}, testAccCheckNsgroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccNsgroupDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dns_nsgroup.test"
	resourceName := "nios_dns_nsgroup.test"
	var v dns.Nsgroup
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsgroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNsgroupDataSourceConfigExtAttrFilters(acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckNsgroupExists(context.Background(), resourceName, &v),
						}, testAccCheckNsgroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckNsgroupResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
        resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
        resource.TestCheckResourceAttrPair(resourceName, "external_primaries", dataSourceName, "result.0.external_primaries"),
        resource.TestCheckResourceAttrPair(resourceName, "external_secondaries", dataSourceName, "result.0.external_secondaries"),
        resource.TestCheckResourceAttrPair(resourceName, "grid_primary", dataSourceName, "result.0.grid_primary"),
        resource.TestCheckResourceAttrPair(resourceName, "grid_secondaries", dataSourceName, "result.0.grid_secondaries"),
        resource.TestCheckResourceAttrPair(resourceName, "is_grid_default", dataSourceName, "result.0.is_grid_default"),
        resource.TestCheckResourceAttrPair(resourceName, "is_multimaster", dataSourceName, "result.0.is_multimaster"),
        resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
        resource.TestCheckResourceAttrPair(resourceName, "use_external_primary", dataSourceName, "result.0.use_external_primary"),
    }
}

func testAccNsgroupDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test" {
}

data "nios_dns_nsgroup" "test" {
  filters = {
	 = nios_dns_nsgroup.test.
  }
}
`)
}

func testAccNsgroupDataSourceConfigExtAttrFilters(extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dns_nsgroup" "test" {
  extattrs = {
    Site = %q
  } 
}

data "nios_dns_nsgroup" "test" {
  extattrfilters = {
	Site = nios_dns_nsgroup.test.extattrs.Site
  }
}
`,extAttrsValue)
}

