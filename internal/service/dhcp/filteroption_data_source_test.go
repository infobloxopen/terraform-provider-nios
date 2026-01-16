package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccFilteroptionDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_filteroption.test"
	resourceName := "nios_dhcp_filteroption.test"
	var v dhcp.Filteroption

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFilteroptionDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFilteroptionDataSourceConfigFilters(acctest.RandomNameWithPrefix("filteroption")),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					}, testAccCheckFilteroptionResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccFilteroptionDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_filteroption.test"
	resourceName := "nios_dhcp_filteroption.test"
	var v dhcp.Filteroption
	name := acctest.RandomNameWithPrefix("filteroption")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFilteroptionDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFilteroptionDataSourceConfigExtAttrFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFilteroptionExists(context.Background(), resourceName, &v),
					}, testAccCheckFilteroptionResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckFilteroptionResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "apply_as_class", dataSourceName, "result.0.apply_as_class"),
		resource.TestCheckResourceAttrPair(resourceName, "bootfile", dataSourceName, "result.0.bootfile"),
		resource.TestCheckResourceAttrPair(resourceName, "bootserver", dataSourceName, "result.0.bootserver"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "expression", dataSourceName, "result.0.expression"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "lease_time", dataSourceName, "result.0.lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "next_server", dataSourceName, "result.0.next_server"),
		resource.TestCheckResourceAttrPair(resourceName, "option_list", dataSourceName, "result.0.option_list"),
		resource.TestCheckResourceAttrPair(resourceName, "option_space", dataSourceName, "result.0.option_space"),
		resource.TestCheckResourceAttrPair(resourceName, "pxe_lease_time", dataSourceName, "result.0.pxe_lease_time"),
	}
}

func testAccFilteroptionDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filteroption" "test" {
  name = %q
}

data "nios_dhcp_filteroption" "test" {
  filters = {
	 name = nios_dhcp_filteroption.test.name
  }
}
`, name)
}

func testAccFilteroptionDataSourceConfigExtAttrFilters(name string, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filteroption" "test" {
  name = %q
  extattrs = {
    Site = %q
  }
}

data "nios_dhcp_filteroption" "test" {
  extattrfilters = {
	Site = nios_dhcp_filteroption.test.extattrs.Site
  }
}
`, name, extAttrsValue)
}
