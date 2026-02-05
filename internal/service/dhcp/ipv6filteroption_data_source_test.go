package dhcp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccIpv6filteroptionDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_ipv6filteroption.test"
	resourceName := "nios_dhcp_ipv6filteroption.test"
	var v dhcp.Ipv6filteroption
	name := acctest.RandomNameWithPrefix("ipv6filteroption")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6filteroptionDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6filteroptionDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					}, testAccCheckIpv6filteroptionResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccIpv6filteroptionDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_dhcp_ipv6filteroption.test"
	resourceName := "nios_dhcp_ipv6filteroption.test"
	var v dhcp.Ipv6filteroption
	name := acctest.RandomNameWithPrefix("ipv6filteroption")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6filteroptionDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6filteroptionDataSourceConfigExtAttrFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpv6filteroptionExists(context.Background(), resourceName, &v),
					}, testAccCheckIpv6filteroptionResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckIpv6filteroptionResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
        resource.TestCheckResourceAttrPair(resourceName, "uuid", dataSourceName, "result.0.uuid"),
		resource.TestCheckResourceAttrPair(resourceName, "apply_as_class", dataSourceName, "result.0.apply_as_class"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "expression", dataSourceName, "result.0.expression"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "lease_time", dataSourceName, "result.0.lease_time"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "option_list", dataSourceName, "result.0.option_list"),
		resource.TestCheckResourceAttrPair(resourceName, "option_space", dataSourceName, "result.0.option_space"),
	}
}

func testAccIpv6filteroptionDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6filteroption" "test" {
  name = %q
}

data "nios_dhcp_ipv6filteroption" "test" {
  filters = {
	name = nios_dhcp_ipv6filteroption.test.name
  }
}
`, name)
}

func testAccIpv6filteroptionDataSourceConfigExtAttrFilters(name, extAttrsValue string) string {
	return fmt.Sprintf(`
resource "nios_dhcp_ipv6filteroption" "test" {
  name = %q
  extattrs = {
    Site = %q
  } 
}

data "nios_dhcp_ipv6filteroption" "test" {
  extattrfilters = {
    Site = nios_dhcp_ipv6filteroption.test.extattrs.Site
  }
}
`, name, extAttrsValue)
}
