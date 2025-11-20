package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccFiltermacDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_filtermac.test"
	resourceName := "nios_dhcp_filtermac.test"
	var v dhcp.Filtermac

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFiltermacDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFiltermacDataSourceConfigFilters("mac_filter"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					}, testAccCheckFiltermacResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccFiltermacDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_filtermac.test"
	resourceName := "nios_dhcp_filtermac.test"
	var v dhcp.Filtermac
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFiltermacDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFiltermacDataSourceConfigExtAttrFilters("mac_filter", acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFiltermacExists(context.Background(), resourceName, &v),
					}, testAccCheckFiltermacResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckFiltermacResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "default_mac_address_expiration", dataSourceName, "result.0.default_mac_address_expiration"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "enforce_expiration_times", dataSourceName, "result.0.enforce_expiration_times"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "lease_time", dataSourceName, "result.0.lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "never_expires", dataSourceName, "result.0.never_expires"),
		resource.TestCheckResourceAttrPair(resourceName, "options", dataSourceName, "result.0.options"),
		resource.TestCheckResourceAttrPair(resourceName, "reserved_for_infoblox", dataSourceName, "result.0.reserved_for_infoblox"),
	}
}

func testAccFiltermacDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filtermac" "test" {
  name = %q
}

data "nios_dhcp_filtermac" "test" {
  filters = {
	 name = nios_dhcp_filtermac.test.name
  }
}
`, name)
}

func testAccFiltermacDataSourceConfigExtAttrFilters(name string, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_filtermac" "test" {
  name = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_dhcp_filtermac" "test" {
  extattrfilters = {
	Site = nios_dhcp_filtermac.test.extattrs.Site
  }
}
`, name, extAttrsValue)
}
