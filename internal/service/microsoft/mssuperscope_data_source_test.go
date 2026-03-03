package microsoft_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoft"

	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccMssuperscopeDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_microsoft_mssuperscope.test"
	resourceName := "nios_microsoft_mssuperscope.test"
	var v microsoft.Mssuperscope
	name := acctest.RandomNameWithPrefix("mssuperscope")
	startAddrRange := "117.0.0.66"
	endAddrRange := "117.0.0.70"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMssuperscopeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMssuperscopeDataSourceConfigFilters(name, startAddrRange, endAddrRange),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					}, testAccCheckMssuperscopeResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccMssuperscopeDataSource_ExtAttrFilters(t *testing.T) {
	dataSourceName := "data.nios_microsoft_mssuperscope.test"
	resourceName := "nios_microsoft_mssuperscope.test"
	var v microsoft.Mssuperscope
	name := acctest.RandomNameWithPrefix("mssuperscope")
	startAddrRange := "117.0.0.71"
	endAddrRange := "117.0.0.75"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMssuperscopeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccMssuperscopeDataSourceConfigExtAttrFilters(name, startAddrRange, endAddrRange, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckMssuperscopeExists(context.Background(), resourceName, &v),
					}, testAccCheckMssuperscopeResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckMssuperscopeResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_utilization", dataSourceName, "result.0.dhcp_utilization"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_utilization_status", dataSourceName, "result.0.dhcp_utilization_status"),
		resource.TestCheckResourceAttrPair(resourceName, "disable", dataSourceName, "result.0.disable"),
		resource.TestCheckResourceAttrPair(resourceName, "dynamic_hosts", dataSourceName, "result.0.dynamic_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "extattrs", dataSourceName, "result.0.extattrs"),
		resource.TestCheckResourceAttrPair(resourceName, "high_water_mark", dataSourceName, "result.0.high_water_mark"),
		resource.TestCheckResourceAttrPair(resourceName, "high_water_mark_reset", dataSourceName, "result.0.high_water_mark_reset"),
		resource.TestCheckResourceAttrPair(resourceName, "low_water_mark", dataSourceName, "result.0.low_water_mark"),
		resource.TestCheckResourceAttrPair(resourceName, "low_water_mark_reset", dataSourceName, "result.0.low_water_mark_reset"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "network_view", dataSourceName, "result.0.network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "ranges", dataSourceName, "result.0.ranges"),
		resource.TestCheckResourceAttrPair(resourceName, "static_hosts", dataSourceName, "result.0.static_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "total_hosts", dataSourceName, "result.0.total_hosts"),
	}
}

func testAccMssuperscopeDataSourceConfigFilters(name, startAddr string, endAddr string) string {
	config := fmt.Sprintf(`
resource "nios_microsoft_mssuperscope" "test" {
    name = %q
    ranges = [nios_dhcp_range.test.ref]
    network_view = "ms_server"
}

data "nios_microsoft_mssuperscope" "test" {
  filters = {
	name = nios_microsoft_mssuperscope.test.name
  }
}
`, name)
	return strings.Join([]string{testAccBaseWithRanges(startAddr, endAddr), config}, "")
}

func testAccMssuperscopeDataSourceConfigExtAttrFilters(name, startAddr, endAddr, extAttrsValue string) string {
	config := fmt.Sprintf(`
resource "nios_microsoft_mssuperscope" "test" {
  name = %q
  ranges = [nios_dhcp_range.test.ref]
  network_view = "ms_server"
  extattrs = {
    Site = %q
  } 
}

data "nios_microsoft_mssuperscope" "test" {
  extattrfilters = {
    Site = nios_microsoft_mssuperscope.test.extattrs.Site
  }
}
`, name, extAttrsValue)
	return strings.Join([]string{testAccBaseWithRanges(startAddr, endAddr), config}, "")
}
