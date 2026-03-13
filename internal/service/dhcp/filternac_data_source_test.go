package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccFilternacDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_filternac.test"
	resourceName := "nios_dhcp_filternac.test"
	var v dhcp.Filternac
	name := acctest.RandomNameWithPrefix("tf-filternac-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFilternacDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFilternacDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFilternacExists(context.Background(), resourceName, &v),
					}, testAccCheckFilternacResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccFilternacDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_filternac.test"
	resourceName := "nios_dhcp_filternac.test"
	var v dhcp.Filternac
	name := acctest.RandomNameWithPrefix("tf-filternac-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFilternacDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFilternacDataSourceConfigExtAttrFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFilternacExists(context.Background(), resourceName, &v),
					}, testAccCheckFilternacResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckFilternacResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "expression", dataSourceName, "result.0.expression"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "lease_time", dataSourceName, "result.0.lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "options", dataSourceName, "result.0.options"),
	}
}

func testAccFilternacDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filternac" "test" {
  name = %q
}

data "nios_dhcp_filternac" "test" {
  filters = {
	name = nios_dhcp_filternac.test.name
  }
}
`, name)
}

func testAccFilternacDataSourceConfigExtAttrFilters(name, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filternac" "test" {
  name = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_dhcp_filternac" "test" {
  extattrfilters = {
	Site = nios_dhcp_filternac.test.extattrs.Site
  }
}
`, name, extAttrsValue)
}
